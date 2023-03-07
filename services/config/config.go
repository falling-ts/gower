package config

import (
	"sync"

	"gower/services"

	_ "github.com/joho/godotenv/autoload"
)

// Configs 配置能力
type Configs interface {
	services.Ability
	Get(fieldStr string, args ...string) any
}

// Config 配置主结构体
type Config struct {
	Configs
}

var (
	cfg  *Config
	once sync.Once
)

// Build 构建单例模式
func Build() *Config {
	once.Do(func() {
		build()
	})

	return cfg
}

// Register 注册服务
func (c *Config) Register(s services.Services) {
	s.SetService(c)
}

// BindAbility 绑定配置能力
func (c *Config) BindAbility(a services.Ability) {
	c.Configs = a.(Configs)
}

// Cfg 获取内部配置
func (c *Config) Cfg() Configs {
	return c.Configs
}

func build() {
	cfg = new(Config)
}
