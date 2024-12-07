package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	app.Controller
}

var Api = new(ApiController)

// Image 上传图片
func (*ApiController) Image(c *gin.Context) (services.Response, error) {
	path, url, err := upload.Image(c)
	if err != nil {
		return nil, exc.BadRequest(err)
	}

	return res.Created("上传成功", app.Data{
		"path": path,
		"url":  url,
	}), nil
}
