package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/services"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// LoginForm 获取登录页面
func (*AuthController) LoginForm() (services.Response, error) {
	return res.Ok("admin/login", app.Data{}), nil
}
