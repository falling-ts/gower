package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
)

type HomeController struct {
	app.Controller
}

var Home = new(HomeController)

// Index 获取页面
func (*HomeController) Index(auth *models.Auth) (services.Response, error) {
	var (
		raw  any
		data app.Data
		err  error
	)

	admin := auth.AdminUser
	if admin.ID != 0 {
		raw, err = admin.SetModel(&admin).Out(app.Rule{
			"name": func() string {
				name := *admin.Nickname
				if name == "" {
					name = *admin.Username
				}
				if name == "" {
					name = "无名者"
				}

				return name
			},
			"avatar": func() string {
				path := *admin.Avatar
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

	return res.Ok("admin/index", data), nil
}
