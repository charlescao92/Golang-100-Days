package routers

import (
	"beegoWebBlog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
    //注册
    beego.Router("/register", &controllers.RegisterController{})
    //登录
    beego.Router("/login", &controllers.LoginController{})
    //退出
	beego.Router("/exit", &controllers.ExitController{})
}
