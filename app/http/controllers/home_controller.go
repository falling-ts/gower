package controllers

import (
	"gower/services"
)

type HomeController struct {
	Controllers
}

var Home = new(HomeController)

// Index 主页
func (h *HomeController) Index() (services.Response, error) {
	return h.response("home/index", Data{
		"Title": "欢迎来到 Gower",
	}), nil
}
