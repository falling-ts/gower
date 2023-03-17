package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"

	"github.com/caarlos0/env/v7"
)

var _ services.Config = (*config.Service)(nil)

func init() {
	c := new(configs.Config)
	if err := env.Parse(c); err != nil {
		panic(err)
	}

	ss.Config = config.Mount(c).(*configs.Config)
}
