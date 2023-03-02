package controllers

import (
	"gower/app/services/route"
	"net/http"
)

type AuthController struct {
	Controllers
}

var Auth = new(AuthController)

// RegisterForm get register form
func (a *AuthController) RegisterForm(c route.Context) {
	c.HTML(http.StatusOK, "auth/register", data{
		"Title": "register",
	})
}

// Register exec register
func (a *AuthController) Register(c route.Context) {
	c.JSON(http.StatusOK, data{
		"code": 0,
		"msg":  "SUCCESS",
		"data": nil,
	})
}

// LoginForm get login form
func (a *AuthController) LoginForm(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "login",
	})
}

// Login exec login
func (a *AuthController) Login(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "login",
	})
}

// Me get personal center
func (a *AuthController) Me(c route.Context) {
	c.HTML(http.StatusOK, "auth/me", data{
		"Title": "me",
	})
}

// Logout exec logout
func (a *AuthController) Logout(c route.Context) {
	c.HTML(http.StatusOK, "auth/login", data{
		"Title": "login",
	})
}
