package configs

import (
	"gower/services"
	"gower/services/config"
	"reflect"
	"strings"
	"unicode"

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

// Get 获取配置参数, 包含默认值
func (c *Configs) Get(fieldStr string, args ...string) any {
	def := ""
	if len(args) > 0 {
		def = args[0]
	}

	fields := strings.Split(fieldStr, ".")
	if len(fields) == 0 {
		return def
	}

	for i, field := range fields {
		r := []rune(field)
		r[0] = unicode.ToUpper(r[0])
		fields[i] = string(r)
	}

	var cfgValue reflect.Value
	for i, field := range fields {
		if i == 0 {
			cfgValue = reflect.ValueOf(c).Elem().FieldByName(field)
		} else {
			cfgValue = cfgValue.FieldByName(field)
		}

		if !cfgValue.IsValid() {
			return def
		}
	}

	return cfgValue
}
