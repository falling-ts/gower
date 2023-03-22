package util

import (
	"github.com/jaevor/go-nanoid"
	"gower/services"
	"reflect"
)

type Service struct{}

// New 新建 Util 服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(...services.Service) services.Service {
	return s
}

// Nanoid 获取简单唯一 ID
func (s *Service) Nanoid(args ...int) string {
	arg := 21
	if len(args) > 0 {
		arg = args[0]
	}

	genKey, err := nanoid.Standard(arg)
	if err != nil {
		panic(err)
	}
	return genKey()
}

// Direct 获取反射指针类型
func (s *Service) Direct(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Pointer {
		return v.Addr()
	}

	return v
}
