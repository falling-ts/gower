package providers

import (
	"gower/services"
)

// ServicesMap 服务集合
type ServicesMap map[string]services.Service

// Services 在内存分配服务集合
var Services = make(ServicesMap)

// Register 服务集合注册服务
func (s ServicesMap) Register(key string, service services.Service) {
	s[key] = service
}
