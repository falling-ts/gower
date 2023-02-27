package routes

import (
	"gower/app"
	"gower/app/services"
)

var route = app.Get("route").(*services.Route)
