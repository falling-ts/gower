package middlewares

import "gower/app"

var (
	config = app.Config()
	logger = app.Logger()
	route  = app.Route()
)

func init() {
	route.Use(Recovery()).
		Use(Logger())
}
