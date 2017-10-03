package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"seacloud/utils"
	"seacloud/models"
)

type LinkController struct {
	beego.Controller
}

func (this *LinkController)GetTmpDownloadLink() {
	username := this.GetSession("username")
	p := this.GetString("path")
	ret := make(map[string]string)
	token, err := models.GenerateTmpDownloadToken(username.(string), p, 2)
	if err != nil {
		ret["error"] = "error"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	link := utils.GetTmpDownloadLink(token)
	
	ret["link"] = link
	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *LinkController)GetTmpUploadLink() {
	username := this.GetSession("username")
	p := this.GetString("path")
	ret := make(map[string]string)
	token := models.GenerateTmpUploadToken(username.(string), p, 2)

	link := utils.GetTmpUploadLink(token)
	
	ret["link"] = link
	this.Data["json"] = &ret
	this.ServeJSON()
}