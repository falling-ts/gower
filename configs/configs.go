package configs

import (
	"gower/services"
	"gower/services/config"

	"github.com/caarlos0/env/v7"
)

type Configs struct {
	App
	Log
}

var _ config.Configs = (*Configs)(nil)

// Link 配置能力链接到配置服务上, 然后由服务绑定配置能力.
func (c *Configs) Link(s services.Service) {
	s.BindAbility(c)
	if err := env.Parse(c); err != nil {
		panic(err)
	}
}
