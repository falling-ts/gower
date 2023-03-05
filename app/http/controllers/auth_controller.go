package controllers

import (
	"gower/app/http/requests"
	"gower/services/route"
	"net/http"
)

type AuthController struct {
	Controllers
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm(c route.Context) {
	c.HTML(http.StatusOK, "auth/register", data{
		"Title": "注册",
	})
}

// Register 执行注册
func (a *AuthController) Register(req requests.RegisterRequest, test *requests.LoginRequest) {
	req.JSON(http.StatusOK, data{
		"code": 0,
		"msg":  "注册成功",
		"data": nil,
	})
}

// LoginForm 登录页面
func (a *AuthController) LoginForm(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "登录",
	})
}

// Login 执行登录
func (a *AuthController) Login(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "登录",
	})
}

// Me 获取个人信息
func (a *AuthController) Me(c route.Context) {
	c.HTML(http.StatusOK, "auth/me", data{
		"Title": "我",
	})
}

// Logout 执行退出
func (a *AuthController) Logout(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "退出",
	})
}
