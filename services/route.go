package services

import (
	"html/template"
	"net"
	"net/http"
)

// Handler 将 gin 中间件使用的处理程序定义为返回值.
type Handler any

// Handlers HandlersChain 定义 HandlerFunc 的切片.
type Handlers []Handler

// IRouter 定义所有路由器句柄接口, 包括单路由器和组路由器.
type IRouter interface {
	IRoutes
	Group(string, ...Handler) *Route
}

// IRoutes 定义所有路由器句柄接口.
type IRoutes interface {
	Use(...Handler) IRoutes

	Handle(string, string, ...Handler) IRoutes
	Any(string, ...Handler) IRoutes
	GET(string, ...Handler) IRoutes
	POST(string, ...Handler) IRoutes
	DELETE(string, ...Handler) IRoutes
	PATCH(string, ...Handler) IRoutes
	PUT(string, ...Handler) IRoutes
	OPTIONS(string, ...Handler) IRoutes
	HEAD(string, ...Handler) IRoutes
	Match([]string, string, ...Handler) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

type Route interface {
	Service

	Delims(left, right string) Route
	LoadHTMLGlob(pattern string)
	LoadHTMLFiles(files ...string)
	SetHTMLTemplate(tmpl *template.Template)
	SetFuncMap(funcMap template.FuncMap)

	Run(addr ...string) (err error)
	RunTLS(addr, certFile, keyFile string) (err error)
	RunUnix(file string) (err error)
	RunFd(fd int) (err error)
	RunListener(listener net.Listener) (err error)

	Use(middleware ...Handler) IRoutes
	Group(relativePath string, handlers ...Handler) Route
	Handle(httpMethod, relativePath string, handlers ...Handler) IRoutes
	Any(relativePath string, handlers ...Handler) IRoutes
	GET(relativePath string, handlers ...Handler) IRoutes
	POST(relativePath string, handlers ...Handler) IRoutes
	DELETE(relativePath string, handlers ...Handler) IRoutes
	PATCH(relativePath string, handlers ...Handler) IRoutes
	PUT(relativePath string, handlers ...Handler) IRoutes
	OPTIONS(relativePath string, handlers ...Handler) IRoutes
	HEAD(relativePath string, handlers ...Handler) IRoutes
	Match(methods []string, relativePath string, handlers ...Handler) IRoutes

	StaticFile(relativePath, filepath string) IRoutes
	StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes
	Static(relativePath, root string) IRoutes
	StaticFS(relativePath string, fs http.FileSystem) IRoutes

	Response(data ResponseData, args ...any) Response
}

type Response interface {
	DecideType(data ResponseData, arg any)
}

// ResponseData 正常响应数据接口
type ResponseData interface {
	Set(any) ResponseData
}
