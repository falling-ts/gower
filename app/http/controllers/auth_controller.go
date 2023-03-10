package controllers

import (
	"fmt"
	"gower/app/http/requests"
	"gower/services"
)

type AuthController struct {
	Controllers
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm() (string, Data) {
	return "auth/register", Data{
		"Title": "注册",
	}
}

// Register 执行注册
func (a *AuthController) Register(req requests.RegisterRequest) services.Response {
	fmt.Println(req)
	return a.response("注册成功")
}

// LoginForm 登录页面
func (a *AuthController) LoginForm() (services.Response, error) {
	return a.response("auth/login", Data{
		"Title": "登录",
	}), nil
}

// Login 执行登录
func (a *AuthController) Login() services.Response {
	return a.response("auth/login", Data{
		"Title": "登录",
	})
}

// Me 获取个人信息
func (a *AuthController) Me() (services.Response, error) {
	return a.response("auth/me", Data{
		"Title": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout() services.Response {
	return a.response("auth/login", Data{
		"Title": "退出",
	})
}
