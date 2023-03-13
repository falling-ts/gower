package resources

import (
	"gower/app"
)

var route = app.App.Route()

func init() {
	route.LoadHTMLGlob("resources/views/**/*.tmpl")
}
