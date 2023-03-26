package resources

import (
	"github.com/falling-ts/gower/app"
)

var route = app.Route()

func init() {
	route.LoadHTMLGlob("resources/views/**/*.tmpl")
}
