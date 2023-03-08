package providers

import (
	"errors"
	"gower/services"
	"gower/services/config"
	"gower/services/exception"
	"gower/services/route"
)

// Services 核心结构体, 所有服务的挂载结构
type Services struct {
	ConfigService
	ExceptionService
	RouteService
}

var _ services.Services = (*Services)(nil)

// Mount 挂载注册服务
func (s *Services) Mount() services.Services {
	config.Build().Register(s)
	exception.Build().Register(s)
	route.Build().Register(s)
	return s
}

// BindContent 绑定服务的内容
func (s *Services) BindContent() services.Services {
	config.Build().BindContent(buildConfigs())
	exception.Build().BindContent(buildExceptions())
	return s
}

// SetService 实际挂载操作
func (s *Services) SetService(service services.Service) {
	switch service.(type) {
	case ConfigService:
		s.ConfigService = service.(ConfigService)
	case ExceptionService:
		s.ExceptionService = service.(ExceptionService)
	case RouteService:
		s.RouteService = service.(RouteService)
	default:
		panic(errors.New("未知服务"))
	}
}
