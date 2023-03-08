package config

import (
	"reflect"
	"strings"
	"sync"
	"unicode"

	"gower/services"

	_ "github.com/joho/godotenv/autoload"
)

// Configs 配置内容
type Configs interface{}

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

// BindContent 绑定配置内容
func (c *Config) BindContent(configs Configs) {
	c.Configs = configs
}

// Get 获取配置参数, 包含默认值
func (c *Config) Get(fieldStr string, args ...string) any {
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

// Cfg 获取内部配置
func (c *Config) Cfg() Configs {
	return c.Configs
}

func build() {
	cfg = new(Config)
}
