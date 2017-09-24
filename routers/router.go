package routers

import (
	"seacloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test_json", &controllers.TestJsoncontroller{})
}
