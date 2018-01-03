package routers

import (
	"seacloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/is_user_login/", &controllers.UserController{}, "post:IsTokenValid")
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/api/user/logout", &controllers.UserController{}, "post:UserLogout")
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
	beego.Router("/api/profile/get", &controllers.ProfileController{}, "get:GetProfile")
	beego.Router("/api/profile/post", &controllers.ProfileController{}, "post:PostProfile")
	beego.Router("/api/avatar/upload", &controllers.AvatarController{}, "post:UploadAvatar")
	beego.Router("/api/avatar/get", &controllers.AvatarController{}, "get:GetAvatar")
	beego.Router("/api/user/changepw", &controllers.UserController{}, "post:ChangePassword")
	beego.Router("/api/user/add", &controllers.UserController{}, "post:AddUser")
	beego.Router("/api/user/getusername", &controllers.UserController{}, "get:GetUserName")
	beego.Router("/api/user/getallusers", &controllers.UserController{}, "get:GetAllUsers")
	beego.Router("/api/group/new", &controllers.GroupController{}, "post:NewGroup")
	//beego.Router("/api/group/getallgroups", &controllers.GroupController{}, "get:GetAllGroups")
	beego.Router("/api/files/get_dir_img_files", &controllers.ImageFileController{})
	beego.Router("/api/files/share/generate_download_link", &controllers.ShareController{}, "post:GenerateDownloadLink")
	beego.Router("/api/files/share/get_download_link_info", &controllers.ShareController{}, "get:GetDownloadLinkInfo")
	beego.Router("/api/files/share/delete_download_link", &controllers.ShareController{}, "post:DeleteDownloadLink")
	beego.Router("/test_json", &controllers.TestJsoncontroller{})
}
