package response

import (
	"gower/services"

	"github.com/gin-gonic/gin"
)

// Service 响应结构体
type Service struct {
	services.Response
	HttpStatus int
	config     gin.Negotiate
}

// Mount 挂载响应体
func Mount(r services.Response) services.Response {
	return r.Set(new(Service)).Set(r)
}

// New 新建响应服务
func New() *Service {
	return new(Service)
}

func (s *Service) Init(...any) {}

// Build 构建每个请求的异常
func (s *Service) Build(code int, args ...any) services.Response {
	s.config.Offered = services.Accepts

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

	c.Negotiate(s.HttpStatus, s.config)
	return true
}
