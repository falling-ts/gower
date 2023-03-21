package controllers

import (
	"gower/app"
	"gower/services"
)

type HomeController struct {
	app.Controller
}

var Home = new(HomeController)

// Index 主页
func (h *HomeController) Index() (services.Response, error) {
	return res.Ok("home/index", app.Data{
		"Title": "欢迎来到 Gower",
	}), nil
}
