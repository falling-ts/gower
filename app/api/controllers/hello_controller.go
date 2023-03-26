package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/api/requests"
)

type HelloController struct {
	app.Controller
}

var Hello = new(HelloController)

func (t *HelloController) Index(req *requests.IndexRequest) (string, any) {
	return "Hello, " + *req.Name, app.Data{
		"key": "value",
	}
}
