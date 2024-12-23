package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	app.Controller
}

var Auth = new(AuthController)

// LoginForm 获取登录页面
func (*AuthController) LoginForm() (services.Response, error) {
	return res.Ok("admin/login", app.Data{
		"appTitle": "后台登录",
	}), nil
}

// Login 执行登录
func (a *AuthController) Login(req *requests.AuthRequest, admin *models.AdminUser) (services.Response, error) {
	if err := admin.From(*req.Username); err != nil {
		return nil, exc.BadRequest(err)
	}

	err := passwd.Check(req.Password, admin.Password)
	if err != nil {
		return nil, exc.BadRequest(err, "密码错误")
	}

	token, err := admin.Login(req.RemoteIP())
	if err != nil {
		return nil, exc.BadRequest(err, "登录失败")
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
		return nil, exc.BadRequest(err)
	}
	return res.Ok("退出成功"), nil
}
