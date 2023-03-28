package routes

import (
	web "github.com/falling-ts/gower/app/http/controllers"
	mws "github.com/falling-ts/gower/app/http/middlewares"
	"github.com/falling-ts/gower/public"
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

	if public.Static == nil {
		route.StaticFile("/favicon.ico", "public/static/images/favicon.ico")
		route.Static("/static", "public/static")
	} else {
		route.StaticFileFS("/favicon.ico", "images/favicon.ico", public.Static)
		route.StaticFS("/static", public.Static)
	}

	route.Static("/uploads", "public/uploads")
}
