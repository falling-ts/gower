package providers

import (
	"gower/app/exceptions"
	"gower/configs"
	"gower/services"
)

type Services struct {
	services.Cache
	*configs.Configs
	*exceptions.Exceptions
	services.Route
	services.Validator
	services.DB
}

// Services 在内存分配服务集合
var s = new(Services)

// InitServices 初始化服务集合
func InitServices() *Services {
	s.Validator.Init()
	s.Cache.Init(s.Configs)
	s.DB.Init(s.Configs)
	s.Route.Init(s.Configs, s.Cache, s.Exceptions, s.DB)

	return s
}
