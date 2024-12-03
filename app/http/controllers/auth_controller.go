package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/http/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// RegisterForm 注册页面
func (a *AuthController) RegisterForm() (string, app.Data) {
	return "auth/register", app.Data{
		"appTitle": "注册",
	}
}

// Register 执行注册
func (a *AuthController) Register(req *requests.RegisterRequest, user *models.User) (services.Response, error) {
	model, err := user.In(req, app.Rule{
		"password": func(req requests.RegisterRequest) (string, error) {
			return passwd.Hash(req.Password)
		},
		"_skips": app.Skips{},
	})
	if err != nil {
		return nil, exc.BadRequest(err)
	}

	if err = model.(*models.User).Register(); err != nil {
		return nil, exc.BadRequest(err)
	}

	return res.Ok("注册成功"), nil
}

// LoginForm 登录页面
func (a *AuthController) LoginForm() (string, app.Data) {
	return "auth/login", app.Data{
		"appTitle": "登录",
	}
}

// Login 执行登录
func (a *AuthController) Login(req *requests.LoginRequest, user *models.User) (services.Response, error) {
	if err := user.From(*req.Username); err != nil {
		return nil, exc.BadRequest(err)
	}

	err := passwd.Check(req.Password, user.Password)
	if err != nil {
		return nil, exc.BadRequest(err, "密码错误")
	}

	token, err := user.Login(req.RemoteIP())
	if err != nil {
		return nil, exc.BadRequest(err, "登录失败")
	}

	return res.Ok("登录成功", token), nil
}

// Me 获取个人信息
func (a *AuthController) Me() (services.Response, error) {
	return res.Ok("auth/me", app.Data{
		"appTitle": "我",
	}), nil
}

// Logout 执行退出
func (a *AuthController) Logout(c *gin.Context) (services.Response, error) {
	c.Set("token", nil)
	token, _ := cookie.Get(c, "token")
	if token == "" {
		token = c.GetHeader("Authorization")
	}

	if err := auth.Black(token); err != nil {
		return nil, exc.BadRequest(err)
	}
	return res.Ok("退出成功"), nil
}
