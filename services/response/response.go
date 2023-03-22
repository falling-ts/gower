package response

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

// Service 响应结构体
type Service struct {
	services.Response
	HttpStatus int
	config     gin.Negotiate
}

var (
	config services.Config
	auth   services.AuthService
)

// Mount 挂载响应体
func Mount(r services.Response) services.Response {
	return r.Set(new(Service)).Set(r)
}

// New 新建响应服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	auth = args[1].(services.AuthService)
	return s.Response
}

// Build 构建每个请求的异常
func (s *Service) Build(code int, args ...any) services.Response {
	s.config.Offered = accepts

	s.decideType("success")
	argsNum := len(args)
	for i := 0; i < argsNum; i++ {
		s.decideType(args[i])
	}

	s.HttpStatus = code

	return s.Response
}

// Handle 处理响应
func (s *Service) Handle(c *gin.Context) bool {
	if c.NegotiateFormat(s.config.Offered...) != gin.MIMEHTML {
		c.Set("body-logger", s.Response)
	} else {
		c.Set("body-logger", "html body")
	}

	var token string
	if tokenAny, ok := c.Get("token"); !ok {
		tokenAny, _ = s.Response.Get("token")
		token = tokenAny.(string)
	}

	if token != "" {
		c.SetCookie("token",
			token,
			100000000,
			"/",
			config.Get("app.domain", "localhost").(string),
			false,
			false)

		s.Set(token)
	}

	c.Negotiate(s.HttpStatus, s.config)
	return true
}

func (s *Service) decideType(arg any) {
	switch arg.(type) {
	case int:
		code := arg.(int)
		if code >= http.StatusOK && code < http.StatusMultipleChoices {
			s.HttpStatus = code
		} else {
			s.Response.Set(arg)
		}
	case string:
		s.Response.Set(arg)
		s.config.HTMLName = arg.(string)
	default:
		s.Response.Set(arg)
		s.config.HTMLData = arg
	}

	s.config.Data = s.Response
}

// IsToken 判断是否是 Token
func (s *Service) IsToken(token string) bool {
	return auth.IsToken(token)
}
