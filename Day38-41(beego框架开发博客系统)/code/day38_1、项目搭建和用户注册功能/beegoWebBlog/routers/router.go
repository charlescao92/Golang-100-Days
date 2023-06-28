package routers

import (
	"beegoWebBlog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	//注册
	beego.Router("/register", &controllers.RegisterController{})
}
