package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

//正则路由控制器
type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {

	//*全匹配
	//logs.Info("全匹配：" + this.Ctx.Input.URL())
	//this.Ctx.Output.Body([]byte("请求URL:" + this.Ctx.Input.URL()))

	//变量匹配
	//id := this.Ctx.Input.Param(":name")
	//logs.Info("Id:" + id)
	//this.Ctx.ResponseWriter.Write([]byte("Id:" + id))

	//*.*匹配
	path := this.Ctx.Input.Param(":path")
	logs.Info(path)
	ext := this.Ctx.Input.Param(":ext")
	logs.Info(ext)
	this.Ctx.ResponseWriter.Write([]byte("filePath: " + path + " , ext: " + ext))

	//int类型匹配 只能匹配int类型的
	//id := this.Ctx.Input.Param(":id")
	//this.Ctx.ResponseWriter.Write([]byte("int类型变量值：" + id))

	//string类型匹配 只能匹配string类型
	userName := this.Ctx.Input.Param(":username")
	this.Ctx.ResponseWriter.Write([]byte("string类型变量值：" + userName))
}
