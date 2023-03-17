package exception

import "gower/services"

// Service 异常服务
type Service struct {
	services.Exception
	RawErr error
}

// Mount 挂载异常内容
func Mount(e services.Exception) services.Exception {
	s := new(Service)
	s.Exception = e
	e.Set(s)

	return e
}

// New 创建新异常服务
func New() *Service {
	return new(Service)
}

// Init 服务初始化
func (s *Service) Init(...any) {}

// Build 构建每个请求的异常
func (s *Service) Build(args ...any) services.Exception {
	s.Exception.Set("未知异常")
	argsNum := len(args)

	for i := 0; i < argsNum; i++ {
		decideType(args[i], s)
	}

	return s.Exception
}
