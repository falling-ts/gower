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

		user := ar.Group("/user", mws.Auth())
		{
			user.GET("/", admin.Admin.Index)
		}

		role := ar.Group("/role", mws.Auth())
		{
			role.GET("/", admin.Role.Index)
		}

		menu := ar.Group("/menu", mws.Auth())
		{
			menu.GET("/", admin.Menu.Index)
		}

		permission := ar.Group("/permission", mws.Auth())
		{
			permission.GET("/", admin.Permission.Index)
		}
	}
}
