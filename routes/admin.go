package routes

import admin "github.com/falling-ts/gower/app/admin/controllers"

func init() {
	ar := route.Group("/admin")
	{
		auth := ar.Group("/auth")
		{
			auth.GET("/login", admin.Auth.LoginForm)
		}
	}
}
