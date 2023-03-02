package route

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Route struct {
	*gin.Engine
}

var (
	route *Route
	once  sync.Once
)

// New singleton route service
func New() *Route {
	once.Do(func() {
		route = &Route{
			gin.Default(),
		}
	})

	return route
}

// Use adds middleware to the group, see example code in GitHub.
func (r *Route) Use(middleware ...HandlerFunc) IRoutes {
	r.Engine.Use(toGinHandlers(middleware)...)
	return r
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func (r *Route) Group(relativePath string, handlers ...HandlerFunc) *Route {
	group := r.Engine.Group(relativePath, toGinHandlers(handlers)...)
	r.Engine.RouterGroup = *group
	return r
}

// Handle registers a new request handle and middleware with the given path and method.
// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
// See the example code in GitHub.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (r *Route) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.Handle(httpMethod, relativePath, toGinHandlers(handlers)...)
	return r
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (r *Route) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.Any(relativePath, toGinHandlers(handlers)...)
	return r
}

// GET is a shortcut for route.Handle("GET", path, handlers).
func (r *Route) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.GET(relativePath, toGinHandlers(handlers)...)
	return r
}

// POST is a shortcut for route.Handle("POST", path, handlers).
func (r *Route) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.POST(relativePath, toGinHandlers(handlers)...)
	return r
}

// DELETE is a shortcut for route.Handle("DELETE", path, handlers).
func (r *Route) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.DELETE(relativePath, toGinHandlers(handlers)...)
	return r
}

// PATCH is a shortcut for route.Handle("PATCH", path, handlers).
func (r *Route) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.PATCH(relativePath, toGinHandlers(handlers)...)
	return r
}

// PUT is a shortcut for route.Handle("PUT", path, handlers).
func (r *Route) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.PUT(relativePath, toGinHandlers(handlers)...)
	return r
}

// OPTIONS is a shortcut for route.Handle("OPTIONS", path, handlers).
func (r *Route) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.OPTIONS(relativePath, toGinHandlers(handlers)...)
	return r
}

// HEAD is a shortcut for route.Handle("HEAD", path, handlers).
func (r *Route) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.HEAD(relativePath, toGinHandlers(handlers)...)
	return r
}

// Match registers a route that matches the specified methods that you declared.
func (r *Route) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	r.Engine.Match(methods, relativePath, toGinHandlers(handlers)...)
	return r
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *Route) StaticFile(relativePath, filepath string) IRoutes {
	r.Engine.StaticFile(relativePath, filepath)
	return r
}

// StaticFileFS works just like `StaticFile` but a custom `http.FileSystem` can be used instead..
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin by default uses: gin.Dir()
func (r *Route) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {
	r.Engine.StaticFileFS(relativePath, filepath, fs)
	return r
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//
//	router.Static("/static", "/var/www")
func (r *Route) Static(relativePath, root string) IRoutes {
	r.Engine.Static(relativePath, root)
	return r
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default uses: gin.Dir()
func (r *Route) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	r.Engine.StaticFS(relativePath, fs)
	return r
}

func toGinHandlers(handlers HandlersChain) gin.HandlersChain {
	ginHandlers := make(gin.HandlersChain, len(handlers))
	for i, handler := range handlers {
		ginHandlers[i] = toGinHandler(handler)
	}

	return ginHandlers
}

func toGinHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
	}
}
