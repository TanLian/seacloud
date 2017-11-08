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
	beego.Router("/api/file/rename", &controllers.FileController{}, "post:RenameFile")
	beego.Router("/api/file/new", &controllers.FileController{}, "post:NewFile")
	beego.Router("/api/dir/new", &controllers.FileController{}, "post:NewDir")
	beego.Router("/api/files/trash", &controllers.FileController{}, "get:GetTrashFiles")
	beego.Router("/api/files/trash/clear", &controllers.FileController{}, "delete:ClearTrashFiles")
	beego.Router("/api/files/trash/restore", &controllers.FileController{}, "post:RestoreTrashSingleFile")
	beego.Router("/api/files/favorites", &controllers.FileController{}, "get:GetFavorateFiles")
	beego.Router("/api/files/favorites/add", &controllers.FileController{}, "post:AddFavorateFile")
	beego.Router("/api/files/favorites/delete", &controllers.FileController{}, "post:DeleteFavorateFile")
	beego.Router("/test_json", &controllers.TestJsoncontroller{})
}
