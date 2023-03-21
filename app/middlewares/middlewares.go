package middlewares

import "gower/app"

var (
	config = app.Config()
	logger = app.Logger()
	route  = app.Route()
	excp   = app.Exception()
)

func init() {
	route.Use(Recovery()).Use(Logger())
}
