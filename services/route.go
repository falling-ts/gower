package services

import (
	"github.com/gin-gonic/gin"
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
	Group(string, ...Handler) IRouter
}

// IRoutes 定义所有路由器句柄接口.
type IRoutes interface {
	Use(...Handler) RouteService

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

// RouteService 路由服务接口
type RouteService interface {
	Service

	Handler() http.Handler
	SecureJsonPrefix(prefix string) RouteService

	Delims(left, right string) RouteService
	LoadHTMLGlob(pattern string)
	LoadHTMLFiles(files ...string)
	SetHTMLTemplate(tmpl *template.Template)
	SetFuncMap(funcMap template.FuncMap)

	NoRoute(handlers ...Handler)
	NoMethod(handlers ...Handler)

	Routes() (routes gin.RoutesInfo)

	Run(addr ...string) (err error)
	RunTLS(addr, certFile, keyFile string) (err error)
	RunUnix(file string) (err error)
	RunFd(fd int) (err error)
	RunListener(listener net.Listener) (err error)

	HandleContext(c *gin.Context)

	Group(relativePath string, handlers ...Handler) IRouter
	UseBefore(middleware ...Handler) RouteService
	Use(middleware ...Handler) RouteService
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

	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
