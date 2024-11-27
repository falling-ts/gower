package routes

import (
	web "github.com/falling-ts/gower/app/http/controllers"
	mws "github.com/falling-ts/gower/app/http/middlewares"
)

func init() {
	route.GET("/", mws.Default(), web.Home.Index)

	// 注册与登录
	auth := route.Group("/auth")
	{
		auth.GET("/register", web.Auth.RegisterForm)
		auth.POST("/register", web.Auth.Register)
		auth.GET("/login", web.Auth.LoginForm)
		auth.POST("/login", web.Auth.Login)
		auth.GET("/me", mws.Auth(), web.Auth.Me)
		auth.POST("/logout", mws.Auth(), web.Auth.Logout)
	}

	route.GET("/400", web.Excp.BadRequest)
	route.GET("/404", web.Excp.NotFound)
}
