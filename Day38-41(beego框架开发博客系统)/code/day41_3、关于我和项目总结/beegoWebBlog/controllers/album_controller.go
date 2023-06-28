package controllers

import (
	"beegoWebBlog/models"
	"github.com/beego/beego/v2/core/logs"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
