package controllers

import (
	"fmt"
	"gower/app"
	"gower/app/models"
	"gower/services"
)

type HomeController struct {
	app.Controller
}

var Home = new(HomeController)

// Index 主页
func (h *HomeController) Index(auth *models.Auth) (services.Response, error) {
	var (
		raw  any
		data app.Data
		err  error
	)
	username := auth.Username
	fmt.Println(username)
	if auth.ID != 0 {
		raw, err = auth.SetModel(auth).Out(app.Rule{
			"name": func() string {
				name := *auth.Nickname
				if name == "" {
					name = *auth.Username
				}
				if name == "" {
					name = "无名者"
				}

				return name
			},
			"avatar": func() string {
				path := *auth.Avatar
				if path == "" {
					path = "/static/images/avatar.png"
				}

				return config.App.Url + path
			},
		})
		if err != nil {
			return nil, excp.BadRequest(err)
		}

		data, _ = raw.(map[string]any)
	}
	if data == nil {
		data = make(app.Data)
	}

	data["title"] = "欢迎来到 Gower"
	return res.Ok("home/index", data), nil
}
