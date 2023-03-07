package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"
)

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	services.Service

	Configs() configs.Configs
}
