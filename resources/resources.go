package resources

import (
	"gower/app"
)

var route = app.Route()

func init() {
	route.LoadHTMLGlob("resources/views/**/*.tmpl")
}
