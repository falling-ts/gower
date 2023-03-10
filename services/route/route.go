package route

import (
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// Route 路由服务主结构体
type Route struct {
	*gin.Engine
}

var (
	configs    services.Configs
	cache      services.Cache
	exceptions services.Exceptions
)

func New() services.Route {
	return new(Route)
}

// Init 初始化
func (r *Route) Init(args ...any) {
	if len(args) < 3 {
		panic("路由服务初始化参数不全.")
	}
	configs = args[0].(services.Configs)
	cache = args[1].(services.Cache)
	exceptions = args[2].(services.Exceptions)

	r.Engine = gin.New()
	setLogger(r.Engine)
	setRecovery(r.Engine)
}

// Delims 设置模板的左右界限, 并返回一个引擎实例.
func (r *Route) Delims(left, right string) services.Route {
	r.Engine.Delims(left, right)
	return r
}

// Use 将中间件添加到组中, 参见GitHub中的示例代码.
func (r *Route) Use(middleware ...services.Handler) services.IRoutes {
	r.Engine.Use(toGinHandlers(middleware)...)
	return r
}

// Group 创建一个新的路由器组, 您应该添加所有具有公共中间件或相同路径前缀的路由.
// 例如, 所有使用公共中间件进行授权的路由都可以分组.
func (r *Route) Group(relativePath string, handlers ...services.Handler) services.Route {
	group := r.Engine.Group(relativePath, toGinHandlers(handlers)...)

	route := &Route{gin.New()}
	route.Engine.RouterGroup = *group
	return route
}

// Handle 用给定的路径和方法注册一个新的请求句柄和中间件.
// 最后一个处理程序应该是真正的处理程序, 其他的应该是中间件, 可以并且应该在不同的路由之间共享.
// 参见GitHub中的示例代码.
//
// 对于 GET, POST, PUT, PATCH 和 DELETE 请求各自的快捷方式可以使用函数.
//
// 此功能用于批量加载, 并允许使用不常用的、非标准化的或自定义的方法(例如, 用于内部与代理的通信).
func (r *Route) Handle(httpMethod, relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.Handle(httpMethod, relativePath, toGinHandlers(handlers)...)
	return r
}

// Any 注册一个匹配所有HTTP方法的路由。
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (r *Route) Any(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.Any(relativePath, toGinHandlers(handlers)...)
	return r
}

// GET 是 route.Handle("GET", path, handlers) 的短语形式.
func (r *Route) GET(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.GET(relativePath, toGinHandlers(handlers)...)
	return r
}

// POST 是 route.Handle("POST", path, handlers) 的短语形式.
func (r *Route) POST(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.POST(relativePath, toGinHandlers(handlers)...)
	return r
}

// DELETE 是 route.Handle("DELETE", path, handlers) 的短语形式.
func (r *Route) DELETE(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.DELETE(relativePath, toGinHandlers(handlers)...)
	return r
}

// PATCH 是 route.Handle("PATCH", path, handlers) 的短语形式.
func (r *Route) PATCH(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.PATCH(relativePath, toGinHandlers(handlers)...)
	return r
}

// PUT 是 route.Handle("PUT", path, handlers) 的短语形式.
func (r *Route) PUT(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.PUT(relativePath, toGinHandlers(handlers)...)
	return r
}

// OPTIONS 是 route.Handle("OPTIONS", path, handlers) 的短语形式.
func (r *Route) OPTIONS(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.OPTIONS(relativePath, toGinHandlers(handlers)...)
	return r
}

// HEAD 是 route.Handle("HEAD", path, handlers) 的短语形式.
func (r *Route) HEAD(relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.HEAD(relativePath, toGinHandlers(handlers)...)
	return r
}

// Match 注册与您声明的指定方法匹配的路由.
func (r *Route) Match(methods []string, relativePath string, handlers ...services.Handler) services.IRoutes {
	r.Engine.Match(methods, relativePath, toGinHandlers(handlers)...)
	return r
}

// StaticFile 注册一个路由, 以便为本地文件系统的一个文件服务.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *Route) StaticFile(relativePath, filepath string) services.IRoutes {
	r.Engine.StaticFile(relativePath, filepath)
	return r
}

// StaticFileFS 与 StaticFile 类似, 但可以使用自定义的 http.FileSystem.
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin 默认使用: gin.Dir()
func (r *Route) StaticFileFS(relativePath, filepath string, fs http.FileSystem) services.IRoutes {
	r.Engine.StaticFileFS(relativePath, filepath, fs)
	return r
}

// Static 方法从给定的文件系统根目录中提供文件服务.
// 它内部使用了 http.FileServer, 因此使用了 http.NotFound 而不是路由器的 NotFound 处理程序.
// 如果要使用操作系统的文件系统实现，请使用以下方法
// router.Static("/static", "/var/www")
//
// 其中, 第一个参数是 URL 的前缀，第二个参数是文件系统根目录的路径.
// 这将把以 /static 开头的 URL 映射到 /var/www 目录下的文件系统.
func (r *Route) Static(relativePath, root string) services.IRoutes {
	r.Engine.Static(relativePath, root)
	return r
}

// StaticFS 方法与 Static() 类似, 但是可以使用自定义的 http.FileSystem.
// 默认情况下, Gin 使用 gin.Dir().
func (r *Route) StaticFS(relativePath string, fs http.FileSystem) services.IRoutes {
	r.Engine.StaticFS(relativePath, fs)
	return r
}

func toGinHandlers(handlers services.Handlers) gin.HandlersChain {
	ginHandlers := make(gin.HandlersChain, len(handlers))
	for i, handler := range handlers {
		ginHandlers[i] = transHandler(handler)
	}

	return ginHandlers
}
