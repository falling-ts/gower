package routes

import (
	admin "gitee.com/falling-ts/gower/app/admin/controllers"
	mws "gitee.com/falling-ts/gower/app/admin/middlewares"
)

func init() {
	ar := route.Group("/admin", mws.Default(), mws.Menus())
	{
		ar.GET("/", mws.Auth(), admin.Home.Index)

		auth := ar.Group("/auth")
		{
			auth.GET("/login", admin.Auth.LoginForm)
			auth.POST("/login", admin.Auth.Login)
			auth.POST("/logout", mws.Auth(), admin.Auth.Logout)
		}

		ar.Restful("user", admin.Admin)
		ar.Restful("role", admin.Role)
		ar.Restful("permission", admin.Permission)
		ar.Restful("menu", admin.Menu)
	}
}
