package configs

import (
	"gower/services"
	"gower/services/config"
)

// Configs 总配置
type Configs struct {
	*config.Config
	App
	Log
	Cache
}

var _ services.Config = (*Configs)(nil)
var _ services.Configs = (*Configs)(nil)

// Set 通用设置内容
func (c *Configs) Set(arg any) {
	switch arg.(type) {
	case *config.Config:
		c.Config = arg.(*config.Config)
	}
}
