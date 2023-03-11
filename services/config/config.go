package config

import (
	"reflect"
	"strings"
	"sync"
	"unicode"

	"gower/services"

	_ "github.com/joho/godotenv/autoload"
)

type Content interface{}

// Struct 配置主结构体
type Struct struct {
	Content
	mu sync.RWMutex
}

var Entity = new(Struct)

// Init 服务初始化
func (c *Struct) Init(args ...any) services.Service {
	if len(args) == 0 {
		panic("初始化参数不存在")
	}

	content, ok := args[0].(Content)
	if !ok {
		panic("配置服务初始化失败")
	}
	c.Content = content

	return c
}

// Get 获取配置参数, 包含默认值
func (c *Struct) Get(fieldStr string, args ...any) any {
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
		r := []rune(field)
		r[0] = unicode.ToUpper(r[0])
		fields[i] = string(r)
	}

	var cfgValue reflect.Value
	for i, field := range fields {
		if i == 0 {
			cfgValue = reflect.ValueOf(c.Content).Elem().FieldByName(field)
		} else {
			cfgValue = cfgValue.FieldByName(field)
		}

		if !cfgValue.IsValid() {
			return def
		}
	}

	return cfgValue.Interface()
}

// Configs 获取内部配置
func (c *Struct) Configs() Content {
	return c.Content
}
