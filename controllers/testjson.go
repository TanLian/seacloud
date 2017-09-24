package controllers

import (
	"github.com/astaxie/beego"
)

type TestJsoncontroller struct {
	beego.Controller
}

func (this *TestJsoncontroller) Get() {
	ret := make(map[string]string)
	ret["name"] = "TanLian"
	ret["age"] = "24"
	this.Data["json"] = &ret
	this.ServeJSON()
}
