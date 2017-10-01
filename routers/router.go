package routers

import (
	"seacloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/api/files", &controllers.FileController{})
	beego.Router("/test_json", &controllers.TestJsoncontroller{})
}
