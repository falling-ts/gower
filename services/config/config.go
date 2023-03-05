package config

import (
	"sync"

	"gower/app/providers"
	"gower/configs"

	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	configs *configs.Configs
}

var (
	cfg  *Config
	once sync.Once
)

func New() *Config {
	once.Do(func() {
		build()
	})

	return cfg
}

func (c *Config) Register(services *providers.Services) {
	services.ConfigService = c
}

func (c *Config) Configs() *configs.Configs {
	return c.configs
}

func build() {
	cfg = new(Config)
	if err := env.Parse(cfg.configs); err != nil {
		panic(err)
	}
}
