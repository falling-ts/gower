package config

import (
	"gower/services"
	"sync"

	"gower/configs"

	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
)

// Config 服务主结构体
type Config struct {
	configs configs.Configs
}

var (
	cfg  *Config
	once sync.Once
)

// New 简单工厂与单例创建
func New() *Config {
	once.Do(func() {
		build()
	})

	return cfg
}

// Register 注册服务
func (c *Config) Register(s services.Services) {
	s.SetService(c)
}

// Configs 获取内部配置
func (c *Config) Configs() configs.Configs {
	return c.configs
}

func build() {
	cfg = new(Config)
	if err := env.Parse(&cfg.configs); err != nil {
		panic(err)
	}
}
