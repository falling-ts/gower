package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"

	"github.com/caarlos0/env/v7"
)

var _ ConfigService = (*config.Config)(nil)

type ConfigService interface {
	services.Service

	BindContent(configs config.Configs)
	Get(fieldStr string, args ...string) any
	Cfg() config.Configs
}

func buildConfigs() *configs.Configs {
	c := new(configs.Configs)
	if err := env.Parse(c); err != nil {
		panic(err)
	}

	return c
}
