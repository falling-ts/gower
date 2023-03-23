package routes

import api "gower/app/api/controllers"

func init() {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/hello", api.Hello.Index)
	}
}
