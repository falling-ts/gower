package providers

import "gower/services/config"
import "gower/configs"

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	Service
	Configs() *configs.Configs
}
