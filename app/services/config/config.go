package config

import (
	"sync"

	"gower/config"

	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
)

var (
	cfg  *config.Config
	once sync.Once
)

func New() *config.Config {
	once.Do(func() {
		build()
	})

	return cfg
}

func build() {
	cfg = &config.Config{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
}
