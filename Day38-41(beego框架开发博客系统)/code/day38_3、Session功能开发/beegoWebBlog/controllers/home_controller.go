package controllers

import (
	"fmt"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func (this *HomeController)Get(){
	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
	this.TplName="home.html"
}