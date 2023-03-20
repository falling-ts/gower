package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

var _ services.Config = (*config.Service)(nil)

func init() {
	P.Register("config", func(...services.Service) services.Service {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}

		c := new(configs.Config)
		if err := env.Parse(c); err != nil {
			panic(err)
		}

		return config.Mount(c).Init()
	})
}
