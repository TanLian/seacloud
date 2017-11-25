package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"seacloud/models"
	"net/http"
)

type UserController struct {
	beego.Controller
}

type loginForm struct {
	UserName string	`json:"username"`
	Password string `json:"password"`
}
func (this *UserController)Post() {
	//拿到用户名和密码
	var form loginForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	fmt.Println(form)

	ret := make(map[string]string)

	//查找数据库，根据用户名获取user对象
	user, err := models.GetUserByName(form.UserName)
	if err != nil {
		fmt.Println(err)
		ret["error"] = "User does not exist."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//校验用户名和密码
	result, err := user.CheckUserPass(form.Password)
	if err != nil || result == false {
		fmt.Println(err)
		ret["error"] = "Password is not correct."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//用户名和密码均正确
	fmt.Println("Username and password are both correct.")

	//设置session
	this.SetSession("username", form.UserName)

	//生成token
	token := models.GenToken()
	cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: 3600}
	http.SetCookie(this.Ctx.ResponseWriter, &cookie)
	this.Redirect("/#/apps/files", 302)
}

type changePasswordForm struct {
	Password string	`json:"password"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
func (this *UserController)ChangePassword() {
	ret := make(map[string]string)

	var form changePasswordForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &form)

	if form.Password == "" || form.NewPassword == "" {
		ret["error"] = "Password and newpassword can not be null"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	if form.NewPassword != form.ConfirmPassword {
		ret["error"] = "Inconsistent Password and confirmation password"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	username := this.GetSession("username")

	//查找数据库，根据用户名获取user对象
	user, err := models.GetUserByName(username.(string))
	if err != nil {
		fmt.Println(err)
		ret["error"] = "User does not exist."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	result, err := user.CheckUserPass(form.Password)
	if err != nil || result == false {
		fmt.Println(err)
		ret["error"] = "Password is not correct."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	err = user.ChangePassWord(form.NewPassword)
	if err != nil {
		fmt.Println(err)
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
}