package config

import (
	"reflect"
	"strings"
	"sync"

	"gower/services"
	"gower/utils/str"
)

// Service 配置服务
type Service struct {
	services.Config
	mu sync.RWMutex
}

// Mount 挂载配置内容
func Mount(c services.Config) services.Config {
	return c.Set(new(Service)).Set(c)
}

// Init 服务初始化
func (s *Service) Init(...services.Service) services.Service {
	return s.Config
}

// Get 获取配置参数, 包含默认值
func (s *Service) Get(fieldStr string, args ...any) any {
	s.mu.RLock()
	defer s.mu.RUnlock()

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

		fields[i] = str.Conv(field).Uppercase()
	}

	var cfgValue reflect.Value
	for i, field := range fields {
		if i == 0 {
			cfgValue = reflect.ValueOf(s.Config).Elem().FieldByName(field)
		} else {
			cfgValue = cfgValue.FieldByName(field)
		}

		if !cfgValue.IsValid() {
			return def
		}
	}

	return cfgValue.Interface()
}
