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

	data, err := models.GetAvatarDataByUsername(username.(string))
	if err != nil {
		//说明未设置头像，返回默认的
		ret["success"] = "success"
		ret["data"] = "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"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	ret["success"] = "success"
	ret["data"] = data
	this.Data["json"] = &ret
	this.ServeJSON()
}