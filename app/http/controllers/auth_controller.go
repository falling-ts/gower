package controllers

import (
	"gower/app/http/requests"
	"gower/app/models"
	"gower/services"
)

type AuthController struct {
	Controller
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm() (string, Data) {
	return "auth/register", Data{
		"Title": "注册",
	}
}

// Register 执行注册
func (a *AuthController) Register(req requests.RegisterRequest, user *models.User) (services.Response, error) {
	model, err := user.In(&req, models.Rule{
		"username": "username",
		"password(password)": func(args ...any) (string, error) {
			return passwd.Hash(args[0].(string))
		},
	})
	if err != nil {
		return nil, excp.BadRequest(err)
	}

	if err := model.(*models.User).Register(); err != nil {
		return nil, excp.BadRequest(err)
	}

	return a.ok("注册成功"), nil
}

// LoginForm 登录页面
func (a *AuthController) LoginForm() (services.Response, error) {
	return a.ok("auth/login", Data{
		"Title": "登录",
	}), nil
}

// Login 执行登录
func (a *AuthController) Login() services.Response {
	return a.ok("auth/login", Data{
		"Title": "登录",
	})
}

// Me 获取个人信息
func (a *AuthController) Me() (services.Response, error) {
	return a.ok("auth/me", Data{
		"Title": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout() services.Response {
	return a.ok("auth/login", Data{
		"Title": "退出",
	})
}
