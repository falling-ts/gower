package route

import (
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// Handler 将 gin 中间件使用的处理程序定义为返回值.
type Handler any

// Handlers HandlersChain 定义 HandlerFunc 的切片.
type Handlers []Handler

// IRouter 定义所有路由器句柄接口, 包括单路由器和组路由器.
type IRouter interface {
	IRoutes
	Group(string, ...Handler) *Struct
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

// Struct 路由服务主结构体
type Struct struct {
	*gin.Engine
}

var Entity = &Struct{
	gin.New(),
}

// Init 初始化
func (r *Struct) Init(args ...any) services.Service {
	setLogger(r.Engine)
	setRecovery(r.Engine)

	return r
}

// Delims 设置模板的左右界限, 并返回一个引擎实例.
func (r *Struct) Delims(left, right string) *Struct {
	r.Engine.Delims(left, right)
	return r
}

// Use 将中间件添加到组中, 参见GitHub中的示例代码.
func (r *Struct) Use(middleware ...Handler) IRoutes {
	r.Engine.Use(toGinHandlers(middleware)...)
	return r
}

// Group 创建一个新的路由器组, 您应该添加所有具有公共中间件或相同路径前缀的路由.
// 例如, 所有使用公共中间件进行授权的路由都可以分组.
func (r *Struct) Group(relativePath string, handlers ...Handler) *Struct {
	group := r.Engine.Group(relativePath, toGinHandlers(handlers)...)
	r.Engine.RouterGroup = *group
	return r
}

// Handle 用给定的路径和方法注册一个新的请求句柄和中间件.
// 最后一个处理程序应该是真正的处理程序, 其他的应该是中间件, 可以并且应该在不同的路由之间共享.
// 参见GitHub中的示例代码.
//
// 对于 GET, POST, PUT, PATCH 和 DELETE 请求各自的快捷方式可以使用函数.
//
// 此功能用于批量加载, 并允许使用不常用的、非标准化的或自定义的方法(例如, 用于内部与代理的通信).
func (r *Struct) Handle(httpMethod, relativePath string, handlers ...Handler) IRoutes {
	r.Engine.Handle(httpMethod, relativePath, toGinHandlers(handlers)...)
	return r
}

// Any 注册一个匹配所有HTTP方法的路由。
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (r *Struct) Any(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.Any(relativePath, toGinHandlers(handlers)...)
	return r
}

// GET 是 route.Handle("GET", path, handlers) 的短语形式.
func (r *Struct) GET(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.GET(relativePath, toGinHandlers(handlers)...)
	return r
}

// POST 是 route.Handle("POST", path, handlers) 的短语形式.
func (r *Struct) POST(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.POST(relativePath, toGinHandlers(handlers)...)
	return r
}

// DELETE 是 route.Handle("DELETE", path, handlers) 的短语形式.
func (r *Struct) DELETE(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.DELETE(relativePath, toGinHandlers(handlers)...)
	return r
}

// PATCH 是 route.Handle("PATCH", path, handlers) 的短语形式.
func (r *Struct) PATCH(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.PATCH(relativePath, toGinHandlers(handlers)...)
	return r
}

// PUT 是 route.Handle("PUT", path, handlers) 的短语形式.
func (r *Struct) PUT(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.PUT(relativePath, toGinHandlers(handlers)...)
	return r
}

// OPTIONS 是 route.Handle("OPTIONS", path, handlers) 的短语形式.
func (r *Struct) OPTIONS(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.OPTIONS(relativePath, toGinHandlers(handlers)...)
	return r
}

// HEAD 是 route.Handle("HEAD", path, handlers) 的短语形式.
func (r *Struct) HEAD(relativePath string, handlers ...Handler) IRoutes {
	r.Engine.HEAD(relativePath, toGinHandlers(handlers)...)
	return r
}

// Match 注册与您声明的指定方法匹配的路由.
func (r *Struct) Match(methods []string, relativePath string, handlers ...Handler) IRoutes {
	r.Engine.Match(methods, relativePath, toGinHandlers(handlers)...)
	return r
}

// StaticFile 注册一个路由, 以便为本地文件系统的一个文件服务.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (r *Struct) StaticFile(relativePath, filepath string) IRoutes {
	r.Engine.StaticFile(relativePath, filepath)
	return r
}

// StaticFileFS 与 StaticFile 类似, 但可以使用自定义的 http.FileSystem.
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin 默认使用: gin.Dir()
func (r *Struct) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {
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
func (r *Struct) Static(relativePath, root string) IRoutes {
	r.Engine.Static(relativePath, root)
	return r
}

// StaticFS 方法与 Static() 类似, 但是可以使用自定义的 http.FileSystem.
// 默认情况下, Gin 使用 gin.Dir().
func (r *Struct) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	r.Engine.StaticFS(relativePath, fs)
	return r
}

func toGinHandlers(handlers Handlers) gin.HandlersChain {
	ginHandlers := make(gin.HandlersChain, len(handlers))
	for i, handler := range handlers {
		ginHandlers[i] = transHandler(handler)
	}

	return ginHandlers
}
