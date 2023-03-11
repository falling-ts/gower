package providers

import (
	"gower/configs"
	"gower/services"
	"gower/services/config"

	"github.com/caarlos0/env/v7"
)

var _ Config = (*config.Struct)(nil)

// Config 适配接口
type Config interface {
	services.Service

	Get(fieldStr string, args ...any) any
	Configs() config.Content
}

func initConfig() {
	c := new(configs.All)
	if err := env.Parse(c); err != nil {
		panic(err)
	}
	config.Entity.Init(c)

	Services.Register("config", config.Entity)
}
