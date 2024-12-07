package route

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"reflect"

	"gitee.com/falling-ts/gower/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// Service 路由服务主结构体
type Service struct {
	*gin.Engine
	delims render.Delims
}

var (
	config    services.Config
	exception services.Exception
	db        services.DBService
	response  services.Response
	util      services.UtilService
)

func New() services.RouteService {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	errors.As(args[1].(services.Exception), &exception)
	db = args[2].(services.DBService)
	response = args[3].(services.Response)
	util = args[4].(services.UtilService)

	mode := map[string]string{
		"development": "debug",
		"production":  "release",
		"test":        "debug",
	}[config.Get("app.mode", "test").(string)]
	gin.SetMode(mode)

	s.Engine = gin.New()
	s.delims = render.Delims{Left: "{{", Right: "}}"}

	return s
}

// SecureJsonPrefix 设置Context.SecureJSON中使用的SecureJsonPrefix.
func (s *Service) SecureJsonPrefix(prefix string) services.RouteService {
	s.Engine.SecureJsonPrefix(prefix)
	return s
}

// Delims 设置模板的左右界限, 并返回一个引擎实例.
func (s *Service) Delims(left, right string) services.RouteService {
	s.delims = render.Delims{Left: left, Right: right}
	s.Engine.Delims(left, right)
	return s
}

// LoadHTMLGlobs 加载多层目录结构的模板
func (s *Service) LoadHTMLGlobs(patterns ...string) error {
	var filenames []string
	left := s.delims.Left
	right := s.delims.Right

	for _, pattern := range patterns {
		list, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}
		if len(list) != 0 {
			filenames = append(filenames, list...)
		}
	}

	tmpl, err := template.New("").Delims(left, right).Funcs(s.Engine.FuncMap).ParseFiles(filenames...)
	if err != nil {
		return err
	}

	s.Engine.SetHTMLTemplate(tmpl)
	return nil
}

// LoadHTMLFS 从嵌入文件加载模板
func (s *Service) LoadHTMLFS(fs fs.FS, patterns ...string) {
	left := s.delims.Left
	right := s.delims.Right
	tmpl := template.Must(template.New("").
		Delims(left, right).
		Funcs(s.Engine.FuncMap).
		ParseFS(fs, patterns...))

	s.Engine.SetHTMLTemplate(tmpl)
}

// UseBefore 将中间件前插进组中
func (s *Service) UseBefore(middleware ...services.Handler) services.RouteService {
	handlers := s.Engine.RouterGroup.Handlers
	s.Engine.RouterGroup.Handlers = append(toGinHandlers(middleware), handlers...)
	return s
}

// Use 将中间件添加到组中, 参见GitHub中的示例代码.
func (s *Service) Use(middleware ...services.Handler) services.RouteService {
	handlers := s.Engine.RouterGroup.Handlers
	s.Engine.RouterGroup.Handlers = append(handlers, toGinHandlers(middleware)...)
	return s
}

// NoRoute 找不到路由
func (s *Service) NoRoute(handlers ...services.Handler) {
	s.Engine.NoRoute(toGinHandlers(handlers)...)
}

// NoMethod 不允许方法
func (s *Service) NoMethod(handlers ...services.Handler) {
	s.Engine.NoMethod(toGinHandlers(handlers)...)
}

// Group 创建一个新的路由器组, 您应该添加所有具有公共中间件或相同路径前缀的路由.
// 例如, 所有使用公共中间件进行授权的路由都可以分组.
func (s *Service) Group(relativePath string, handlers ...services.Handler) services.IRouter {
	group := s.Engine.Group(relativePath, toGinHandlers(handlers)...)

	route := &Service{Engine: gin.New()}
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

// Resource 注册一个资源路由.
func (s *Service) Resource(resource string, controller any, handlers ...services.Handler) services.IRoutes {
	controllerValue := reflect.ValueOf(controller)

	indexMethodValue := controllerValue.MethodByName("Index")
	if indexMethodValue.IsValid() {
		s.GET(fmt.Sprintf("/%s", resource), append(handlers, indexMethodValue.Interface())...)
	}

	createMethodValue := controllerValue.MethodByName("Create")
	if createMethodValue.IsValid() {
		s.GET(fmt.Sprintf("/%s/create", resource), append(handlers, createMethodValue.Interface())...)
	}

	storeMethodValue := controllerValue.MethodByName("Store")
	if storeMethodValue.IsValid() {
		s.POST(fmt.Sprintf("/%s", resource), append(handlers, storeMethodValue.Interface())...)
	}

	editMethodValue := controllerValue.MethodByName("Edit")
	if editMethodValue.IsValid() {
		s.GET(fmt.Sprintf("/%s/:id/edit", resource), append(handlers, editMethodValue.Interface())...)
	}

	updateMethodValue := controllerValue.MethodByName("Update")
	if updateMethodValue.IsValid() {
		s.PUT(fmt.Sprintf("/%s/:id", resource), append(handlers, updateMethodValue.Interface())...)
	}

	showMethodValue := controllerValue.MethodByName("Show")
	if showMethodValue.IsValid() {
		s.GET(fmt.Sprintf("/%s/:id", resource), append(handlers, showMethodValue.Interface())...)
	}

	destroyMethodValue := controllerValue.MethodByName("Destroy")
	if destroyMethodValue.IsValid() {
		s.DELETE(fmt.Sprintf("/%s/:id", resource), append(handlers, destroyMethodValue.Interface())...)
	}
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
