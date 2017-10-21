package controllers

import (
	//"io"
	"os"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"seacloud/utils"
	"path/filepath"
	"seacloud/models"
	"encoding/json"
)

type FileController struct {
	beego.Controller
}

func (this *FileController)Get() {
	//ret := make(map[string][]utils.File)
	ret := make(map[string]interface{})
	username := this.GetSession("username")
	dirPath := this.GetString("path")
	fmt.Println("username:", username)
	fmt.Println("path:", dirPath)

	files := make([]utils.File, 0)
	ret["files"] = files
	if username == nil {
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	files, err := utils.GetFilelistByPath(username.(string), dirPath)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	ret["files"] = files
	ret["success"] = true

	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *FileController)UploadFile() {
	ret := make(map[string]string)
	//从url中拿到token
	token := this.GetString("token")

	//校验token是否合法
	obj, err := models.GetTmpUploadObjByToken(token)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测token值是否超出使用限制
	if obj.UsedTimes >= obj.UsedLimits {
		ret["error"] = "token exceed used limits."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测token是否过期
	if obj.TokenIsExpired() {
		ret["error"] = "token is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测通过
	//token使用次数加1
	obj.TokenUsedTimesAddOne()

	//开始上传
	p := obj.Path
	
	f, h, _ := this.GetFile("file")                  //获取上传的文件
	fmt.Println(h.Filename)
	defer f.Close()

	username := this.GetSession("username")
	fmt.Println(username)

	dataDir := utils.GetDataBaseDir()
	fmt.Println("baseDir:", dataDir)

	this.SaveToFile("file", filepath.Join(dataDir, "root", p, h.Filename))
	
	ret2 := make(map[string][]utils.File)
	files := make([]utils.File, 0)
	mtime := time.Now().Unix()
	obj2 := utils.File{
		Name: h.Filename,
		Size: h.Size,
		Type: "file",
		Mtime: mtime,
		MtimeRelative: utils.Translate_seacloud_time(mtime)}
	files = append(files, obj2)
	ret2["files"] = files
	this.Data["json"] = &ret2
	this.ServeJSON()
}

func (this *FileController)DownloadFile() {
	ret := make(map[string]string)
	//从url中拿到token
	token := this.GetString("token")

	//校验token是否合法
	obj, err := models.GetTmpDownloadObjByToken(token)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测token值是否超出使用限制
	if obj.UsedTimes >= obj.UsedLimits {
		ret["error"] = "token exceed used limits."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测token是否过期
	if obj.TokenIsExpired() {
		ret["error"] = "token is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//检测通过
	//token使用次数加1
	obj.TokenUsedTimesAddOne()

	//开始下载
	username := this.GetSession("username")
	p := obj.Path
	dataDir := utils.GetDataBaseDir()
	fullPath := filepath.Join(dataDir, username.(string), p)
	file, err := os.Open(fullPath)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	defer file.Close()

	filename := filepath.Base(p)
	fmt.Println(filename)
	fmt.Println("+++++")
	/*this.Ctx.Output.Header("Content-Type", "application/octet-stream")
	this.Ctx.Output.Header("content-disposition", "attachment; filename=\""+filename+"\"")
	io.Copy(this.Ctx.ResponseWriter, file)*/
	this.Ctx.Output.Download(fullPath, filename)
}

func (this *FileController)DeleteFile() {
	ret := make(map[string]string)
	username := this.GetSession("username")
	p := this.GetString("p")
	dataDir := utils.GetDataBaseDir()
	fullPath := filepath.Join(dataDir, username.(string), p)

	isExist, err := utils.PathExists(fullPath)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	if isExist == false {
		ret["error"] = "File does not exist."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	err = os.RemoveAll(fullPath)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
}

type RenameFileInfo struct {
	ParentDir string `json:"parent_dir"`
	OldFileName string `json:"old_name"`
	NewFileName string `json:"new_name"`
}
func (this *FileController)RenameFile(){
	errRet := make(map[string]string)
	var params RenameFileInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	username := this.GetSession("username")

	dataDir := utils.GetDataBaseDir()
	fullPath := filepath.Join(dataDir, username.(string), params.ParentDir, params.OldFileName)

	isExist, err := utils.PathExists(fullPath)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	if isExist == false {
		errRet["error"] = "File does not exist."
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	//do rename file
	err = os.Rename(fullPath, filepath.Join(dataDir, username.(string), params.ParentDir, params.NewFileName))
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	ret := make(map[string]bool)
	ret["success"] = true
	this.Data["json"] = &ret
	this.ServeJSON()
}

type FileInfo struct {
	ParentDir string `json:"parent_dir"`
	FileName string `json:"name"`
}
func (this *FileController)NewDir() {
	errRet := make(map[string]string)
	var params FileInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	username := this.GetSession("username")

	dataDir := utils.GetDataBaseDir()
	fullPath := filepath.Join(dataDir, username.(string), params.ParentDir, params.FileName)

	isExist, err := utils.PathExists(fullPath)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	if isExist == true {
		errRet["error"] = "Dir already exist."
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	err = os.Mkdir(fullPath, 0777)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	ret := make(map[string]bool)
	ret["success"] = true
	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *FileController)NewFile() {
	errRet := make(map[string]string)
	var params FileInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	username := this.GetSession("username")

	dataDir := utils.GetDataBaseDir()
	fullPath := filepath.Join(dataDir, username.(string), params.ParentDir, params.FileName)

	isExist, err := utils.PathExists(fullPath)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	if isExist == true {
		errRet["error"] = "File already exist."
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	_, err = os.Create(fullPath)
	if err != nil {
		errRet["error"] = err.Error()
		this.Data["json"] = &errRet
		this.ServeJSON()
		return
	}

	ret := make(map[string]bool)
	ret["success"] = true
	this.Data["json"] = &ret
	this.ServeJSON()
}