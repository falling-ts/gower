package resources

import (
	"gower/app"
)

var route = app.Entity.Route()

func init() {
	route.LoadHTMLGlob("resources/views/**/*.tmpl")
}
