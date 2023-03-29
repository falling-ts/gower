package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type UploadController struct {
	app.Controller
}

var Upload = new(UploadController)

// Image 上传图片
func (*UploadController) Image(c *gin.Context) (services.Response, error) {
	path, url, err := upload.Image(c)
	if err != nil {
		return nil, excp.BadRequest(err)
	}

	return res.Created("上传成功", app.Data{
		"path": path,
		"url":  url,
	}), nil
}
