package config

import (
	"sync"

	"gower/app/providers"
	"gower/config"

	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	*config.App
	*config.Log
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

func build() {
	cfg = new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
}
