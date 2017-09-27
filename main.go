package main

import (
	_ "seacloud/routers"

	"fmt"

	"seacloud/filters"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	//"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//设置日志输出引擎
	beego.BeeLogger.DelLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true) //输出文件名和行号

	//检测用户是否登录，应用于全部的请求，如果未登陆，则重定向到登录页面
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterUser)

	//数据库相关
	dbName := beego.AppConfig.String("dbname")
	dbUser := beego.AppConfig.String("dbuser")
	dbPass := beego.AppConfig.String("dbpass")
	dbHost := beego.AppConfig.String("dbhost")
	dbPort, err := beego.AppConfig.Int("dbport")
	if err != nil {
		dbPort = 3306
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName))

}

func main() {
	beego.Run()
}
