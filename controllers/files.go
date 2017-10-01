package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"seacloud/utils"
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