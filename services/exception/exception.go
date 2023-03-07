package exception

import (
	"sync"

	"gower/app/exceptions"
	"gower/app/providers"
	"gower/services"
)

type Exception struct {
	exceptions.Exception
	RawErr error
}

var (
	exception *Exception
	once      sync.Once
)

// New 简单工厂与单例创建
func New() *Exception {
	once.Do(func() {
		build()
	})

	return exception
}

// Register 注册服务
func (e *Exception) Register(s services.Services) {
	s.SetService(e)
}

// Clone 每个请求一份异常
func (e *Exception) Clone(code uint, args ...any) providers.ExceptionService {
	temp := *e
	newE := &temp
	newE.Exception.Exception = newE
	newE.Exception.Code = code
	newE.Exception.Msg = "未知异常"

	argsNum := len(args)
	if argsNum > 0 {
		decideType(args[0], newE)
	}
	if argsNum > 1 {
		decideType(args[1], newE)
	}
	if argsNum > 2 {
		decideType(args[2], newE)
	}
	if argsNum > 3 {
		decideType(args[3], newE)
	}
	if argsNum > 4 {
		decideType(args[4], newE)
	}
	if argsNum > 5 {
		decideType(args[5], newE)
	}

	return newE
}

func decideType(arg any, newE *Exception) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		newE.Exception.Msg = err.Error()
		newE.RawErr = err.(error)
	case string:
		newE.Exception.Msg = arg.(string)
	default:
		newE.Exception.Data = arg
	}
}

// Excp 获取异常实体
func (e *Exception) Excp() *exceptions.Exception {
	return &e.Exception
}

func build() {
	exception = new(Exception)
	exception.Exception.Exception = exception
}
