package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"seacloud/models"
	"encoding/json"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController)GetProfile() {
	ret := make(map[string]interface{})
	username := this.GetSession("username")

	profile, _ := models.GetProfileByUsername(username.(string))
	if profile != nil {
		ret["username"] = profile.UserName
		ret["motto"] = profile.Motto
		ret["intro"] = profile.Intro
		ret["depart"] = profile.Depart
		ret["tele"] = profile.Tele
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	ret["username"] = ""
	ret["motto"] = ""
	ret["intro"] = ""
	ret["depart"] = ""
	ret["tele"] = ""
	this.Data["json"] = &ret
	this.ServeJSON()
	return
}

type postProfile struct {
	Motto string `json:"motto"`
	Intro string `json:"intro"`
	Depart string `json:"depart"`
	Tele string `json:"tele"`
}
func (this *ProfileController)PostProfile() {
	ret := make(map[string]interface{})
	username := this.GetSession("username")

	var params postProfile

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//fmt.Println(params)
	err = models.InsertOrUpdateProfile(username.(string), params.Motto, params.Intro, params.Depart, params.Tele)
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