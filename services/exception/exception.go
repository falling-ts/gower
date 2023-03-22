package exception

import (
	"net/http"

	"gower/services"

	"github.com/gin-gonic/gin"
)

// 接受的响应数据类型
var accepts = []string{
	gin.MIMEJSON,
	gin.MIMEHTML,
	gin.MIMEXML,
	gin.MIMEYAML,
	gin.MIMETOML,
	gin.MIMEPlain,
}

// Service 异常服务
type Service struct {
	services.Exception
	RawErr error
}

var (
	cache  services.CacheService
	config services.Config
	util   services.UtilService
)

// Mount 挂载异常内容
func Mount(e services.Exception) services.Exception {
	return e.Set(new(Service)).Set(e)
}

// New 新建异常服务
func New() *Service {
	return new(Service)
}

// Init 服务初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	cache = args[1].(services.CacheService)
	util = args[2].(services.UtilService)

	return s.Exception
}

// Build 构建每个请求的异常
func (s *Service) Build(args ...any) services.Exception {
	argsNum := len(args)

	s.decideType("未知异常")
	for i := 0; i < argsNum; i++ {
		s.decideType(args[i])
	}

	if s.RawErr == nil {
		s.RawErr = s.Exception
	}

	return s.Exception
}

// Handle 处理异常
func (s *Service) Handle(c *gin.Context) bool {
	_ = c.Error(s.RawErr)

	c.Set("body-logger", s.Exception)
	switch c.NegotiateFormat(accepts...) {
	case gin.MIMEJSON:
		c.JSON(http.StatusOK, s.Exception)
	case gin.MIMEHTML:
		key := util.ExcpKey()
		cache.SetDefault(key, s.Exception)
		c.SetCookie(
			"exception",
			key,
			300,
			"/",
			config.Get("app.domain", "localhost").(string),
			false,
			true)

		referer := c.Request.Referer()
		if referer != "" {
			c.Redirect(http.StatusFound, referer)
			break
		}

		code, _ := s.Exception.Get("code")
		host := config.Get("app.url", "http://localhost:8080").(string)
		switch code.(int) {
		case http.StatusUnauthorized:
			c.Redirect(http.StatusFound, host+"/auth/login")
		}
	case gin.MIMEXML:
		c.XML(http.StatusOK, s.Exception)
	case gin.MIMEYAML:
		c.YAML(http.StatusOK, s.Exception)
	case gin.MIMETOML:
		c.TOML(http.StatusOK, s.Exception)
	default:
		c.String(http.StatusOK, s.Exception.Error())
	}

	c.Abort()
	return true
}

func (s *Service) decideType(arg any) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		_ = s.Exception.Set(err.Error())
		s.RawErr = err.(error)
	case string:
		_ = s.Exception.Set(arg.(string))
	default:
		_ = s.Exception.Set(arg)
	}
}
