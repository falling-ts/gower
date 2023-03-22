package configs

import (
	"gower/services"
	"gower/services/config"
)

var _ services.Config = (*Config)(nil)
var _ services.ConfigService = (*Config)(nil)

// Config 总配置
type Config struct {
	*config.Service
	App
	Log
	Cache
	DB
	Passwd
	Jwt
}

// Set 通用设置内容
func (c *Config) Set(arg any) services.Config {
	switch arg.(type) {
	case *config.Service:
		c.Service = arg.(*config.Service)
	case *Config:
		c.Service.Config = c
	}

	return c
}
