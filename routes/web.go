package routes

import (
	. "gower/app/http/controllers"
	"gower/app/http/middlewares"
)

func init() {
	route.StaticFile("/favicon.ico", "public/static/images/favicon.ico")
	route.Static("/static", "public/static")

	route.GET("/", Home.Index)

	// 注册与登录
	auth := route.Group("/auth")
	{
		auth.GET("/register", Auth.RegisterForm)
		auth.POST("/register", Auth.Register)
		auth.GET("/login", Auth.LoginForm)
		auth.POST("/login", Auth.Login)
		auth.Use(middlewares.Auth()).GET("/me", Auth.Me)
		auth.Use(middlewares.Auth()).POST("/logout", Auth.Logout)
	}
}
