package config

import (
	"reflect"
	"strings"
	"sync"
	"unicode"

	"gower/services"

	_ "github.com/joho/godotenv/autoload"
)

// Config 配置主结构体
type Config struct {
	services.Configs
	mu sync.RWMutex
}

func Mount(c services.Configs) services.Configs {
	config := new(Config)
	config.Configs = c
	c.Set(config)

	return c
}

// Init 服务初始化
func (c *Config) Init(...any) {}

// Get 获取配置参数, 包含默认值
func (c *Config) Get(fieldStr string, args ...any) any {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var def any
	if len(args) > 0 {
		def = args[0]
	}

	fields := strings.Split(fieldStr, ".")
	if len(fields) == 0 {
		return def
	}

	for i, field := range fields {
		if field == "db" {
			fields[i] = strings.ToUpper(field)
			continue
		}

		r := []rune(field)
		r[0] = unicode.ToUpper(r[0])
		fields[i] = string(r)
	}

	var cfgValue reflect.Value
	for i, field := range fields {
		if i == 0 {
			cfgValue = reflect.ValueOf(c.Configs).Elem().FieldByName(field)
		} else {
			cfgValue = cfgValue.FieldByName(field)
		}

		if !cfgValue.IsValid() {
			return def
		}
	}

	return cfgValue.Interface()
}
