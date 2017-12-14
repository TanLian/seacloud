package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	"seacloud/models"
	"encoding/json"
)

type GroupController struct {
	beego.Controller
}

type newGroupForm struct {
	GroupName string	`json:"groupname"`
}
func (this *GroupController)NewGroup() {
	ret := make(map[string]string)

	username := this.GetSession("username")
	if username == nil {
		ret["error"] = "Session is expired, you need to relogin."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var form newGroupForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &form)

	//群组名不能为空
	if form.GroupName == "" {
		ret["error"] = "Group name is required."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	err := models.CreateGroup(form.GroupName, username.(string))
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
	return

}