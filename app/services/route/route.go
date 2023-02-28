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

func (r *Route) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.Use(middleware...)
}

func (r *Route) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.Handle(httpMethod, relativePath, handlers...)
}
func (r *Route) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.Any(relativePath, handlers...)
}
func (r *Route) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.GET(relativePath, toGinHandlerFunc(handlers)...)
}
func (r *Route) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.POST(relativePath, handlers...)
}
func (r *Route) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.DELETE(relativePath, handlers...)
}
func (r *Route) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.PATCH(relativePath, handlers...)
}
func (r *Route) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.PUT(relativePath, handlers...)
}
func (r *Route) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.OPTIONS(relativePath, handlers...)
}
func (r *Route) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.HEAD(relativePath, handlers...)
}
func (r *Route) Match(methods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Engine.RouterGroup.Match(methods, relativePath, handlers...)
}

func (r *Route) StaticFile(relativePath, filepath string) gin.IRoutes {
	return r.Engine.RouterGroup.StaticFile(relativePath, filepath)
}
func (r *Route) StaticFileFS(relativePath, filepath string, fs http.FileSystem) gin.IRoutes {
	return r.Engine.RouterGroup.StaticFileFS(relativePath, filepath, fs)
}
func (r *Route) Static(relativePath, root string) gin.IRoutes {
	return r.Engine.RouterGroup.Static(relativePath, root)
}
func (r *Route) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return r.Engine.RouterGroup.StaticFS(relativePath, fs)
}

func toGinHandlerFunc(hfs []HandlerFunc) []gin.HandlerFunc {
	ret := make([]gin.HandlerFunc, len(hfs))
	for i, hf := range hfs {
		ret[i] = func(c *gin.Context) {
			hf(c)
		}
	}
	return ret
}
