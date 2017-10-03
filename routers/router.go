package routers

import (
	"seacloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/api/files", &controllers.FileController{})
	beego.Router("/api/file/get_tmp_download_link", &controllers.LinkController{}, "get:GetTmpDownloadLink")
	beego.Router("/api/file/get_tmp_upload_link", &controllers.LinkController{}, "get:GetTmpUploadLink")
	beego.Router("/api/file/download", &controllers.FileController{}, "get:DownloadFile")
	beego.Router("/api/file/upload", &controllers.FileController{}, "post:UploadFile")
	beego.Router("/api/file/delete", &controllers.FileController{}, "get:DeleteFile")
	beego.Router("/test_json", &controllers.TestJsoncontroller{})
}
