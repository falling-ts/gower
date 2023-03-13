package providers

import (
	"github.com/caarlos0/env/v7"
	"gower/configs"
	"gower/services"
	"gower/services/config"
)

var _ services.Config = (*config.Config)(nil)

func init() {
	c := new(configs.Configs)
	if err := env.Parse(c); err != nil {
		panic(err)
	}

	s.Configs = config.Mount(c).(*configs.Configs)
}
