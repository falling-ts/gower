package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
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

	user := auth.User
	if user.ID != 0 {
		raw, err = user.SetModel(&user).Out(app.Rule{
			"name": func() string {
				name := *user.Nickname
				if name == "" {
					name = *user.Username
				}
				if name == "" {
					name = "无名者"
				}

				return name
			},
			"avatar": func() string {
				path := *user.Avatar
				if path == "" {
					path = "/public/static/images/avatar.png"
				}

				return config.App.Url + path
			},
		})
		if err != nil {
			return nil, exc.BadRequest(err)
		}

		data, _ = raw.(map[string]any)
	}
	if data == nil {
		data = make(app.Data)
	}

	return res.Ok("home/index", data), nil
}
