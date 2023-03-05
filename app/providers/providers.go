package providers

import (
	"gower/services/config"
	"gower/services/route"
)

// Service 所有服务的通用接口
type Service interface {
	Register(services *Services)
}

// Services 核心结构体, 所有服务的挂载结构
type Services struct {
	ConfigService
	ExceptionService
	RouteService
}

// Mount 挂载注册服务
func (s *Services) Mount() {
	config.New().Register(s)
	route.New().Register(s)
}
