package providers

import (
	"gower/services"
	"gower/services/config"
)

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	services.Service

	Cfg() config.Configs
}
