package routes

import (
	api "github.com/falling-ts/gower/app/api/controllers"
)

func init() {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/hello", api.Hello.Index)
	}
}
