package routes

import (
	admin "github.com/falling-ts/gower/app/admin/controllers"
	mws "github.com/falling-ts/gower/app/admin/middlewares"
)

func init() {
	ar := route.Group("/admin", mws.Default())
	{
		ar.GET("/", mws.Auth(), admin.Home.Index)

		auth := ar.Group("/auth")
		{
			auth.GET("/login", admin.Auth.LoginForm)
			auth.POST("/login", admin.Auth.Login)
			auth.POST("/logout", mws.Auth(), admin.Auth.Logout)
		}
	}
}
