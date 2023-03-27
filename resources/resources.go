package resources

import (
	"embed"
	"github.com/falling-ts/gower/app"
)

var (
	route  = app.Route()
	TmplFS *embed.FS
)

func init() {
	err := route.LoadHTMLGlobs(
		"resources/views/*.tmpl",
		"resources/views/**/*.tmpl")
	if err != nil {
		if TmplFS == nil {
			panic("没有模板内容可加载")
		}
		route.LoadHTMLFS(TmplFS, "views/*.tmpl",
			"views/**/*.tmpl")
	}
}
