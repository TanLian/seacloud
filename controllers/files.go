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
)

type FileController struct {
	beego.Controller
}

func (this *FileController)Get() {
	ret := make(map[string][]utils.File)
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

	err = os.Remove(fullPath)
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