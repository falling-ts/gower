package controllers

import (
	"gower/app/services/route"
	"net/http"
)

type HomeController struct {
	Controllers
}

var Home = new(HomeController)

// Index 主页
func (h *HomeController) Index(c route.Context) {
	c.HTML(http.StatusOK, "home/index", data{
		"Title": "欢迎来到 Gower",
	})
}

// Test 测试页面
func (h *HomeController) Test(c route.Context) {
	c.HTML(http.StatusOK, "home/test", data{
		"Title": "欢迎来到 Gower",
	})
}
