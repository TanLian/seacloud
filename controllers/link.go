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

/*func (this *LinkController)GetUploadLink() {
	username := this.GetSession("username")
	ret := make(map[string]string)

	
}*/

func (this *LinkController)GetTmpDownloadLink() {
	username := this.GetSession("username")
	p := this.GetString("path")
	ret := make(map[string]string)
	token := models.GenerateTmpDownloaToken(username.(string), p, 2)

	link := utils.GetTmpDownloadLink(token)
	
	ret["link"] = link
	this.Data["json"] = &ret
	this.ServeJSON()
}
