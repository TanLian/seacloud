package controllers

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"seacloud/models"
	"seacloud/utils"

	"github.com/astaxie/beego"
)

type ImageFileController struct {
	beego.Controller
}

type ImageFileInfo struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

func (this *ImageFileController) Get() {
	ret := make(map[string]interface{})
	username := this.GetSession("username")
	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	dirPath := this.GetString("path")

	//获取data_dir根目录
	dataDir := utils.GetDataBaseDir()

	dstDir := filepath.Join(dataDir, username.(string), "files", dirPath)

	//开始遍历
	files, err := ioutil.ReadDir(dstDir)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	imgList := make([]ImageFileInfo, 0)

	for _, imgFile := range files {

		if reader, err := os.Open(filepath.Join(dstDir, imgFile.Name())); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			//获得下载链接
			token, err := models.GenerateTmpDownloadToken(username.(string), filepath.Join(dirPath, imgFile.Name()), 2)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			info := ImageFileInfo{
				Name:   imgFile.Name(),
				Width:  im.Width,
				Height: im.Height,
				Url:    utils.GetTmpDownloadLink(token)}
			imgList = append(imgList, info)
			//fmt.Printf("%s %d %d\n", imgFile.Name(), im.Width, im.Height)
		} else {
			fmt.Println("Impossible to open the file:", err)
		}
	}

	ret["success"] = "success"
	ret["files"] = imgList
	this.Data["json"] = &ret
	this.ServeJSON()
	return
}
