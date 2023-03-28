package providers

import (
	"github.com/falling-ts/gower/configs"
	_ "github.com/falling-ts/gower/envs"
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/config"

	"github.com/caarlos0/env/v7"
)

var _ services.Config = (*config.Service)(nil)

func init() {
	P.Register("config", func(...services.Service) services.Service {
		c := new(configs.Config)
		if err := env.Parse(c); err != nil {
			panic(err)
		}

		return config.Mount(c).Init()
	})
}
