package controllers

import (
	//"fmt"

	"encoding/json"
	"path/filepath"
	"seacloud/models"
	"seacloud/utils"
	"time"

	"github.com/astaxie/beego"
)

type ShareController struct {
	beego.Controller
}

type GenerateDownloadLinkInfo struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Expired  string `json:"expired"`
}

func (this *ShareController) GenerateDownloadLink() {
	username := this.GetSession("username")
	ret := make(map[string]string)

	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var params GenerateDownloadLinkInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var expired time.Time
	if params.Expired == "" {
		expired = time.Unix(0, 0)
	} else {
		expired, err = time.Parse("2006-01-02", params.Expired)
		if err != nil {
			ret["error"] = err.Error()
			this.Data["json"] = &ret
			this.ServeJSON()
			return
		}
	}

	p := filepath.Clean(params.Path)

	token := utils.GenerateToken(username.(string), p)
	err = models.StoreDownloadLink(username.(string), token, p, params.Password, expired)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	ret["success"] = "success"
	ret["token"] = token
	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *ShareController) GetDownloadLinkInfo() {
	username := this.GetSession("username")
	ret := make(map[string]interface{})

	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	p := this.GetString("p", "")
	if p == "" {
		ret["error"] = "Path is required."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	p = filepath.Clean(p)

	generated, info, err := models.GetDownloadLinkInfo(username.(string), p)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	ret["success"] = "success"
	if generated == false {
		ret["generated"] = false
	} else {
		ret["generated"] = true
		ret["info"] = info
	}
	this.Data["json"] = &ret
	this.ServeJSON()
}

type DeleteDownloadLinkInfo struct {
	Path string `json:"p"`
}

func (this *ShareController) DeleteDownloadLink() {
	username := this.GetSession("username")
	ret := make(map[string]interface{})

	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var params DeleteDownloadLinkInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	p := filepath.Clean(params.Path)

	err = models.DeleteDownloadLink(username.(string), p)
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

type EditDownloadLinkPasswordInfo struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

func (this *ShareController) EditDownloadLinkPassword() {
	username := this.GetSession("username")
	ret := make(map[string]string)

	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var params EditDownloadLinkPasswordInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	p := filepath.Clean(params.Path)

	err = models.EditDownloadLinkPassword(username.(string), p, params.Password)
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

type EditDownloadLinkExpiredInfo struct {
	Path    string `json:"path"`
	Expired string `json:"expired"`
}

func (this *ShareController) EditDownloadLinkExpired() {
	username := this.GetSession("username")
	ret := make(map[string]string)

	if username == nil {
		ret["error"] = "Session is expired."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var params EditDownloadLinkExpiredInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	var expired time.Time
	if params.Expired == "" {
		expired = time.Unix(0, 0)
	} else {
		expired, err = time.Parse("2006-01-02", params.Expired)
		if err != nil {
			ret["error"] = err.Error()
			this.Data["json"] = &ret
			this.ServeJSON()
			return
		}
	}

	p := filepath.Clean(params.Path)

	err = models.EditDownloadLinkExpired(username.(string), p, expired)
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
