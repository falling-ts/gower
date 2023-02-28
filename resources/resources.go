package resources

import (
	"gower/app/services"
)

var route = services.Route

func init() {
	route.LoadHTMLGlob("resources/views/**/*.tmpl")
}
