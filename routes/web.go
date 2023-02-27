package routes

import "gower/app/http/controllers"

func init() {
	route.GET("/", controllers.Pong.Ping)
}
