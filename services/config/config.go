package config

import (
	"gower/services"
	"sync"

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

// Config 服务名称
func (c *Config) Config() {}

func (c *Config) Register(s services.Services) {
	s.SetService(c)
}

func (c *Config) Configs() *configs.Configs {
	return c.configs
}

func build() {
	cfg = &Config{
		new(configs.Configs),
	}
	if err := env.Parse(cfg.configs); err != nil {
		panic(err)
	}
}
