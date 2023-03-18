package providers

import (
	"gower/app/exceptions"
	"gower/app/responses"
	"gower/configs"
	"gower/services"
)

type Services struct {
	Config    *configs.Config
	Cache     services.CacheService
	Exception *exceptions.Exception
	Route     services.RouteService
	Validator services.ValidatorService
	DB        services.DBService
	Response  *responses.Response
	Logger    services.LoggerService
}

// Services 在内存分配服务集合
var ss = new(Services)

// InitServices 初始化服务集合
func InitServices() *Services {
	ss.Cache.Init(ss.Config)
	ss.DB.Init(ss.Config)
	ss.Exception.Init(ss.Config, ss.Cache)
	ss.Logger.Init(ss.Config)
	ss.Route.Init(ss.Config, ss.Exception, ss.DB, ss.Response)
	ss.Validator.Init()

	return ss
}
