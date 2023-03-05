package exception

import "gower/services"

type Exception struct {
}

// Exception 服务名称
func (e *Exception) Exception() {}

func (e *Exception) Register(s services.Services) {
	s.SetService(e)
}
