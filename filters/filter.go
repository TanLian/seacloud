package filters

import (
	"seacloud/models"
	"fmt"

	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func IsLogin(ctx *context.Context) bool {
	cookie, err := ctx.Request.Cookie("Authorization")
	if err != nil || !models.CheckToken(cookie.Value) {
		//http.Redirect(ctx.ResponseWriter, ctx.Request, "/user/login/?next=/", http.StatusMovedPermanently)
		fmt.Println("校验失败")
		return false
	}
	fmt.Println("校验通过")
	fmt.Println(cookie)
	return true
}

var FilterUser = func(ctx *context.Context) {
	ok := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/user/login/?next=/")
	}
}
