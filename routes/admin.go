package routes

import (
	adm "gitee.com/falling-ts/gower/app/admin/controllers"
	mws "gitee.com/falling-ts/gower/app/admin/middlewares"
)

func init() {
	admin := route.Group("/admin", mws.Default(), mws.Menus())
	{
		admin.GET("/", mws.Auth(), adm.Home.Index)

		upload := admin.Group("/upload", mws.Auth())
		{
			upload.POST("/image", adm.Api.Image)
		}

		auth := admin.Group("/auth")
		{
			auth.GET("/login", adm.Auth.LoginForm)
			auth.POST("/login", adm.Auth.Login)
			auth.POST("/logout", mws.Auth(), adm.Auth.Logout)
		}

		setting := admin.Group("/system", mws.Auth())
		{
			setting.Resource("user", adm.Admin)
			setting.Resource("role", adm.Role)
			setting.Resource("permission", adm.Permission)
			setting.Resource("menu", adm.Menu)
		}
	}
}
