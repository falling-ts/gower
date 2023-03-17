package route

import (
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// Service 路由服务主结构体
type Service struct {
	*gin.Engine
}

var (
	config    services.Config
	exception services.Exception
	db        services.DBService
	response  services.Response
)

func New() services.RouteService {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...any) {
	if len(args) < 4 {
		panic("路由服务初始化参数不全.")
	}
	config = args[0].(services.Config)
	exception = args[1].(services.Exception)
	db = args[2].(services.DBService)
	response = args[3].(services.Response)

	gin.SetMode(config.Get("app.mode", "test").(string))

	s.Engine = gin.New()
	setRecovery(s.Engine)
	setLogger(s.Engine)
}

// Delims 设置模板的左右界限, 并返回一个引擎实例.
func (s *Service) Delims(left, right string) services.RouteService {
	s.Engine.Delims(left, right)
	return s
}

// Use 将中间件添加到组中, 参见GitHub中的示例代码.
func (s *Service) Use(middleware ...services.Handler) services.IRoutes {
	s.Engine.Use(toGinHandlers(middleware)...)
	return s
}

// Group 创建一个新的路由器组, 您应该添加所有具有公共中间件或相同路径前缀的路由.
// 例如, 所有使用公共中间件进行授权的路由都可以分组.
func (s *Service) Group(relativePath string, handlers ...services.Handler) services.RouteService {
	group := s.Engine.Group(relativePath, toGinHandlers(handlers)...)

	route := &Service{gin.New()}
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
func (s *Service) Handle(httpMethod, relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.Handle(httpMethod, relativePath, toGinHandlers(handlers)...)
	return s
}

// Any 注册一个匹配所有HTTP方法的路由。
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (s *Service) Any(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.Any(relativePath, toGinHandlers(handlers)...)
	return s
}

// GET 是 route.Handle("GET", path, handlers) 的短语形式.
func (s *Service) GET(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.GET(relativePath, toGinHandlers(handlers)...)
	return s
}

// POST 是 route.Handle("POST", path, handlers) 的短语形式.
func (s *Service) POST(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.POST(relativePath, toGinHandlers(handlers)...)
	return s
}

// DELETE 是 route.Handle("DELETE", path, handlers) 的短语形式.
func (s *Service) DELETE(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.DELETE(relativePath, toGinHandlers(handlers)...)
	return s
}

// PATCH 是 route.Handle("PATCH", path, handlers) 的短语形式.
func (s *Service) PATCH(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.PATCH(relativePath, toGinHandlers(handlers)...)
	return s
}

// PUT 是 route.Handle("PUT", path, handlers) 的短语形式.
func (s *Service) PUT(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.PUT(relativePath, toGinHandlers(handlers)...)
	return s
}

// OPTIONS 是 route.Handle("OPTIONS", path, handlers) 的短语形式.
func (s *Service) OPTIONS(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.OPTIONS(relativePath, toGinHandlers(handlers)...)
	return s
}

// HEAD 是 route.Handle("HEAD", path, handlers) 的短语形式.
func (s *Service) HEAD(relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.HEAD(relativePath, toGinHandlers(handlers)...)
	return s
}

// Match 注册与您声明的指定方法匹配的路由.
func (s *Service) Match(methods []string, relativePath string, handlers ...services.Handler) services.IRoutes {
	s.Engine.Match(methods, relativePath, toGinHandlers(handlers)...)
	return s
}

// StaticFile 注册一个路由, 以便为本地文件系统的一个文件服务.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (s *Service) StaticFile(relativePath, filepath string) services.IRoutes {
	s.Engine.StaticFile(relativePath, filepath)
	return s
}

// StaticFileFS 与 StaticFile 类似, 但可以使用自定义的 http.FileSystem.
// router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})
// Gin 默认使用: gin.Dir()
func (s *Service) StaticFileFS(relativePath, filepath string, fs http.FileSystem) services.IRoutes {
	s.Engine.StaticFileFS(relativePath, filepath, fs)
	return s
}

// Static 方法从给定的文件系统根目录中提供文件服务.
// 它内部使用了 http.FileServer, 因此使用了 http.NotFound 而不是路由器的 NotFound 处理程序.
// 如果要使用操作系统的文件系统实现，请使用以下方法
// router.Static("/static", "/var/www")
//
// 其中, 第一个参数是 URL 的前缀，第二个参数是文件系统根目录的路径.
// 这将把以 /static 开头的 URL 映射到 /var/www 目录下的文件系统.
func (s *Service) Static(relativePath, root string) services.IRoutes {
	s.Engine.Static(relativePath, root)
	return s
}

// StaticFS 方法与 Static() 类似, 但是可以使用自定义的 http.FileSystem.
// 默认情况下, Gin 使用 gin.Dir().
func (s *Service) StaticFS(relativePath string, fs http.FileSystem) services.IRoutes {
	s.Engine.StaticFS(relativePath, fs)
	return s
}

func toGinHandlers(handlers services.Handlers) gin.HandlersChain {
	ginHandlers := make(gin.HandlersChain, len(handlers))
	for i, handler := range handlers {
		ginHandlers[i] = transHandler(handler)
	}

	return ginHandlers
}
