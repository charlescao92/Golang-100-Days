package routers

import (
	"beeproject/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

/**
 *基础路由：
 *  beego框架提供了常见的http的请求类型方法的路由方案，比如：get，post，head，options，delete等方法
 */
func init() {
	//固定路由的Get请求
	beego.Router("/GetInfo", &controllers.ChangelessController{})

	//固定路由的Post请求
	beego.Router("/PostInfo", &controllers.ChangelessController{})

	//固定路由的Delete请求
	beego.Router("/DeleteInfo", &controllers.ChangelessController{})

	//固定路由的Options请求
	beego.Router("/OptionsInfo", &controllers.ChangelessController{})
}
