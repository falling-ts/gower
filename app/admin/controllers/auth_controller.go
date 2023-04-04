package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/admin/requests"
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// LoginForm 获取登录页面
func (*AuthController) LoginForm() (services.Response, error) {
	return res.Ok("admin/login", app.Data{
		"app_title": "后台登录",
	}), nil
}

// Login 执行登录
func (a *AuthController) Login(req *requests.AuthRequest, admin *models.AdminUser) (services.Response, error) {
	if err := admin.From(*req.Username); err != nil {
		return nil, excp.BadRequest(err)
	}

	err := passwd.Check(req.Password, admin.Password)
	if err != nil {
		return nil, excp.BadRequest(err, "密码错误")
	}

	token, err := admin.Login(req.RemoteIP())
	if err != nil {
		return nil, excp.BadRequest(err, "登录失败")
	}

	return res.Ok("登录成功", token), nil
}

// Logout 执行退出
func (a *AuthController) Logout(c *gin.Context) (services.Response, error) {
	c.Set("admin-token", nil)
	token, _ := cookie.Get(c, "admin-token")
	if token == "" {
		token = c.GetHeader("Admin-Authorization")
	}

	if err := auth.Black(token); err != nil {
		return nil, excp.BadRequest(err)
	}
	return res.Ok("退出成功"), nil
}
