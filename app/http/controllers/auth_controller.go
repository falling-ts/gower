package controllers

import (
	"github.com/gin-gonic/gin"
	"gower/app"
	"gower/app/http/requests"
	"gower/app/models"
	"gower/services"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm(user *models.User) (string, app.Data) {
	return "auth/register", app.Data{
		"title": "注册",
	}
}

// Register 执行注册
func (a *AuthController) Register(req *requests.RegisterRequest, user *models.User) (services.Response, error) {
	model, err := user.In(req, app.Rule{
		"password": func() (string, error) {
			return passwd.Hash(req.Password)
		},
		"_skips": app.Skips{},
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
		"title": "登录",
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

	token, err := user.Login(req.RemoteIP())
	if err != nil {
		return nil, excp.Unauthorized(err, "登录失败")
	}

	return res.Ok("登录成功", token), nil
}

// Me 获取个人信息
func (a *AuthController) Me(auth models.Auth) (services.Response, error) {
	return res.Ok("auth/me", app.Data{
		"title": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout(c *gin.Context) services.Response {
	c.Set("token", nil)
	token, _ := c.Cookie("token")
	if token == "" {
		token = c.GetHeader("Authorization")
	}

	auth.Black(token)
	return res.Ok("退出成功")
}
