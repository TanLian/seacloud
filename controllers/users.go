package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"seacloud/models"
	"net/http"
	"seacloud/utils"
	"os"
	"path/filepath"
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

	//this.Redirect("/#/apps/files", 302)
	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *UserController)UserLogout() {
	ret := make(map[string]string)
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(this.Ctx.ResponseWriter, &cookie)
	//设置session
	this.SetSession("username", "")
	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
}

type tokenForm struct {
	Token string `json:"token"`
}
func (this *UserController)IsTokenValid() {
	ret := make(map[string]string)
	var form tokenForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &form)

	if models.CheckToken(form.Token) {
		ret["success"] = "success"
	}else {
		ret["error"] = "error"
	}
	this.Data["json"] = &ret
	this.ServeJSON()
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

type newUserForm struct {
	UserName string `json:"username"`
	Password string	`json:"password"`
	Source string `json:"source"`
	IsAdmin bool `json:"is_admin"`
}
func (this *UserController)AddUser() {
	ret := make(map[string]string)

	var form newUserForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if form.UserName == "" || form.Password == "" || form.Source == "" {
		ret["error"] = "Username and password can not be null"
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	username := this.GetSession("username")
	if username == nil {
		ret["error"] = "Session is expired, you may need to relogin."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	//查找数据库，根据用户名获取user对象
	user, err := models.GetUserByName(username.(string))
	if err != nil {
		fmt.Println(err)
		ret["error"] = "User does not exist."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	//只有管理员才有权添加新用户
	if user.IsAdmin == false {
		ret["error"] = "Only admin can add user."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	isAdmin := false
	if form.IsAdmin == true {
		isAdmin = true
	}

	err = models.InsertUser(form.UserName, form.Password, isAdmin, "", "", "", form.Source)
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	dataDir := utils.GetDataBaseDir()
	p := filepath.Join(dataDir, form.UserName)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777)
	}

	p = filepath.Join(dataDir, form.UserName, "files")
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777)
	}

	p = filepath.Join(dataDir, form.UserName, "Trash")
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777)
	}

	ret["success"] = "success"
	this.Data["json"] = &ret
	this.ServeJSON()
	return
}

func (this *UserController)GetUserName() {
	ret := make(map[string]string)
	username := this.GetSession("username")
	if username == nil {
		ret["error"] = "Session is expired, you may need to relogin."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}
	ret["success"] = "success"
	ret["username"] = username.(string)
	this.Data["json"] = &ret
	this.ServeJSON()
	return
}

type userProfile struct {
	UserName string
	Avatar	string
	Motto string
	Intro string
	Depart 	string
	Tele string
}
func (this *UserController)GetAllUsers() {
	ret := make(map[string]interface{})
	username := this.GetSession("username")
	if username == nil {
		ret["error"] = "Session is expired, you may need to relogin."
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	userlist := make([]userProfile, 0)

	usernameList, err := models.GetAllUserNames()
	if err != nil {
		ret["error"] = err.Error()
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	for _, name := range usernameList {
		user := userProfile{
			UserName:name,
			Avatar:  models.GetAvatarDataByUsername(name)}
		profile, _ := models.GetProfileByUsername(name)
		if profile != nil {
			user.Motto = profile.Motto
			user.Intro = profile.Intro
			user.Depart = profile.Depart
			user.Tele = profile.Tele
		}
		userlist = append(userlist, user)
	}

	ret["success"] = "success"
	ret["userlist"] = userlist
	this.Data["json"] = &ret
	this.ServeJSON()
	return
}