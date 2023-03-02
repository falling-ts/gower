package controllers

import (
	"gower/app/services/route"
	"net/http"
)

type HomeController struct {
	Controllers
}

var Home = new(HomeController)

func (h *HomeController) Index(c route.Context) {
	c.HTML(http.StatusOK, "home/index", data{
		"Title": "Main website",
	})
}

func (h *HomeController) Test(c route.Context) {
	c.HTML(http.StatusOK, "home/test", data{
		"Title": "Main website",
	})
}
