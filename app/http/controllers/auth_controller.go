package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gower/services"
	"net/http"

	"gower/app/http/requests"
	"gower/services/route"
)

type AuthController struct {
	Controllers
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm(c route.Context) services.Response {
	return a.response(http.StatusOK, "auth/register", data{
		"Title": "注册",
	})
}

// Register 执行注册
func (a *AuthController) Register(req *requests.RegisterRequest) services.Response {
	fmt.Println(req)
	return a.response("注册成功")
}

// LoginForm 登录页面
func (a *AuthController) LoginForm(c route.Context) (services.Response, error) {
	return a.response("auth/login", data{
		"Title": "登录",
	}), nil
}

// Login 执行登录
func (a *AuthController) Login(c *gin.Context) services.Response {
	return a.response("auth/login", data{
		"Title": "登录",
	})
}

// Me 获取个人信息
func (a *AuthController) Me(c *gin.Context) (services.Response, error) {
	return a.response("auth/me", data{
		"Title": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout(c *gin.Context) services.Response {
	return a.response("auth/login", data{
		"Title": "退出",
	})
}
