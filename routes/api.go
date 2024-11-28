package routes

import (
	api "gitee.com/falling-ts/gower/app/api/controllers"
)

func init() {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/hello", api.Hello.Index)
		v1.POST("/upload/image", api.Upload.Image)
	}
}
