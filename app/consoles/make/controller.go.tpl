package controllers

import (
	"{{.Module}}/app"
	"{{.Module}}/app/http/requests"
	"{{.Module}}/services"
)

type {{.UpCamel}}Controller struct {
	app.Controller
}

var {{.UpCamel}} = new({{.UpCamel}}Controller)

// Index 获取页面
func (*{{.UpCamel}}Controller) Index(req *requests.{{.UpCamel}}Request) (services.Response, error) {
	return res.Ok("模板的 define 名称", app.Data{
	    "name": "",
	}), nil
}
