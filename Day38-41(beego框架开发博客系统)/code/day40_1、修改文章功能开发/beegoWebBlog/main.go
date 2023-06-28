package main

import (
	_ "beegoWebBlog/routers"
	"beegoWebBlog/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true // 打开session
	beego.Run()
}
