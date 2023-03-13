package configs

import "gower/services/config"

// Configs 总配置
type Configs struct {
	*config.Config
	App
	Log
	Cache
}

// Set 通用设置内容
func (c *Configs) Set(arg any) {
	switch arg.(type) {
	case *config.Config:
		c.Config = arg.(*config.Config)
	}
}

// Get 获取配置内容
func (c *Configs) Get(fieldStr string, args ...any) any {
	return c.Config.Get(fieldStr, args...)
}
