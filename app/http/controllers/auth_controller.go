package controllers

import (
	"gower/app"
	"gower/app/http/requests"
	"gower/app/models"
	"gower/services"
	"reflect"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm(user *models.User) (string, app.Data) {
	return "auth/register", app.Data{
		"Title": "注册",
	}
}

// Register 执行注册
func (a *AuthController) Register(req *requests.RegisterRequest, user *models.User) (services.Response, error) {
	model, err := user.In(req, app.Rule{
		"password": func(arg any) (string, error) {
			return passwd.Hash(reflect.ValueOf(arg).FieldByName("Password").String())
		},
		"_other": struct{}{},
	})
	if err != nil {
		return nil, excp.BadRequest(err)
	}

	if err = model.(*models.User).Register(); err != nil {
		return nil, excp.BadRequest(err)
	}

	return res.Ok("注册成功"), nil
}

// LoginForm 登录页面
func (a *AuthController) LoginForm() (string, app.Data) {
	return "auth/login", app.Data{
		"Title": "登录",
	}
}

// Login 执行登录
func (a *AuthController) Login(req *requests.LoginRequest, user *models.User) (services.Response, error) {
	if err := user.From(*req.Username); err != nil {
		return nil, excp.BadRequest(err)
	}

	err := passwd.Check(req.Password, user.Password)
	if err != nil {
		return nil, excp.Unauthorized(err, "密码错误")
	}

	token, err := user.Login(req.ClientIP())
	if err != nil {
		return nil, excp.Unauthorized(err, "登录失败")
	}

	return res.Ok("登录成功", token), nil
}

// Me 获取个人信息
func (a *AuthController) Me() (services.Response, error) {
	return res.Ok("auth/me", app.Data{
		"Title": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout() services.Response {
	return res.Ok("auth/login", app.Data{
		"Title": "退出",
	})
}
