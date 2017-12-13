package controllers

import (
	"encoding/base64"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/widuu/goini"
	"seacloud/utils"
	"strconv"
	"seacloud/models"
	//"encoding/json"
)

type AvatarController struct {
	beego.Controller
}

func (this *AvatarController)UploadAvatar() {
	ret := make(map[string]string)
	username := this.GetSession("username")

	f, h, err := this.GetFile("file")                  //获取上传的头像
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	defer f.Close()

	confPath := utils.GetConfPath()
	conf := goini.SetConfig(confPath)
	avatar_max_size := conf.GetValue("GENERA", "avatar_max_size")
	avatar_max_size_int, err :=  strconv.ParseInt(avatar_max_size, 10, 64)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	if h.Size > avatar_max_size_int {
		ret["error"] = "Avatar's size cannot exceed " + avatar_max_size
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	content := make([]byte, 10300)
	_, err = f.Read(content)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	dstData := make([]byte, 20000)
	base64.StdEncoding.Encode(dstData, content)


	err = models.InsertOrUpdateAvatar(username.(string), string(dstData))
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

func (this *AvatarController)GetAvatar() {
	ret := make(map[string]string)
	username := this.GetSession("username")

	if username == nil {
		ret["error"] = "not login"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	data:= models.GetAvatarDataByUsername(username.(string))
	
	ret["success"] = "success"
	ret["data"] = data
	this.Data["json"] = &ret
	this.ServeJSON()
}