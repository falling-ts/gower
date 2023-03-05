package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"
)

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	Config() // 用以区分其它服务
	services.Service

	Configs() *configs.Configs
}
