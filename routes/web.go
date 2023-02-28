package routes

import . "gower/app/http/controllers"

func init() {
	route.Static("/public/static", "public/static")

	route.GET("/ping", Pong.Ping)

	route.GET("/", Home.Index)

	route.GET("/test", Home.Test)
}
