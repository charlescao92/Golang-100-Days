package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：13100000000"
	c.Data["qq"] = "QQ：8888888"
	c.Data["tel"] = "Tel：13100000000"
	c.TplName = "aboultme.html"
}
