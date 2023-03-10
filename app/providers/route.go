package providers

import (
	"html/template"
	"net"
	"net/http"

	"gower/services"

	"gower/services/route"
)

var _ Route = (*route.Struct)(nil)

type Route interface {
	services.Service

	Delims(left, right string) *route.Struct
	LoadHTMLGlob(pattern string)
	LoadHTMLFiles(files ...string)
	SetHTMLTemplate(tmpl *template.Template)
	SetFuncMap(funcMap template.FuncMap)

	Run(addr ...string) (err error)
	RunTLS(addr, certFile, keyFile string) (err error)
	RunUnix(file string) (err error)
	RunFd(fd int) (err error)
	RunListener(listener net.Listener) (err error)

	Use(middleware ...route.Handler) route.IRoutes
	Group(relativePath string, handlers ...route.Handler) *route.Struct
	Handle(httpMethod, relativePath string, handlers ...route.Handler) route.IRoutes
	Any(relativePath string, handlers ...route.Handler) route.IRoutes
	GET(relativePath string, handlers ...route.Handler) route.IRoutes
	POST(relativePath string, handlers ...route.Handler) route.IRoutes
	DELETE(relativePath string, handlers ...route.Handler) route.IRoutes
	PATCH(relativePath string, handlers ...route.Handler) route.IRoutes
	PUT(relativePath string, handlers ...route.Handler) route.IRoutes
	OPTIONS(relativePath string, handlers ...route.Handler) route.IRoutes
	HEAD(relativePath string, handlers ...route.Handler) route.IRoutes
	Match(methods []string, relativePath string, handlers ...route.Handler) route.IRoutes

	StaticFile(relativePath, filepath string) route.IRoutes
	StaticFileFS(relativePath, filepath string, fs http.FileSystem) route.IRoutes
	Static(relativePath, root string) route.IRoutes
	StaticFS(relativePath string, fs http.FileSystem) route.IRoutes
}

func init() {
	route.Entity.Init()
	Services.Register("route", route.Entity)
}
