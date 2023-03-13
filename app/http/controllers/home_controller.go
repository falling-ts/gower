package controllers

import (
	"gower/services"
	"gower/services/route"
)

type HomeController struct {
	Controllers
}

var Home = new(HomeController)

// Index 主页
func (h *HomeController) Index(c route.Context) (services.Response, error) {
	return h.response("home/index", data{
		"Title": "欢迎来到 Gower",
	}), nil
}

// Test 测试页面
func (h *HomeController) Test(c route.Context) services.Response {
	return h.response("home/test", data{
		"Title": "欢迎来到 Gower",
	})
}
