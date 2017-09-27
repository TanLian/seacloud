package filters

import (
	"seacloud/models/user"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func IsLogin(ctx *context.Context) bool {
	cookie, err := ctx.Request.Cookie("Authorization")
	if err != nil || !user.CheckToken(cookie.Value) {
		//http.Redirect(ctx.ResponseWriter, ctx.Request, "/user/login/?next=/", http.StatusMovedPermanently)
		return false
	}
	return true
}

var FilterUser = func(ctx *context.Context) {
	ok := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/user/login/?next=/")
	}
}
