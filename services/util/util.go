package util

import (
	"github.com/jaevor/go-nanoid"
	"gower/services"
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
