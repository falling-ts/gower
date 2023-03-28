package resources

import (
	"embed"
	"github.com/falling-ts/gower/app"
)

var (
	route = app.Route()
	Tmpl  *embed.FS
)

func init() {
	err := route.LoadHTMLGlobs(
		"resources/views/*.tmpl",
		"resources/views/**/*.tmpl")
	if err != nil {
		if Tmpl == nil {
			panic("没有模板内容可加载")
		}
		route.LoadHTMLFS(Tmpl, "views/*.tmpl",
			"views/**/*.tmpl")
	}
}
