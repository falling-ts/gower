package exception

import "gower/services"

// Exception 异常服务
type Exception struct {
	services.Exceptions
	RawErr error
}

// Mount 挂载异常内容
func Mount(e services.Exceptions) services.Exceptions {
	exception := new(Exception)
	exception.Exceptions = e
	e.Set(exception)

	return e
}

// New 创建新异常服务
func New() *Exception {
	return new(Exception)
}

// Init 服务初始化
func (e *Exception) Init(...any) {}

// Build 构建每个请求的异常
func (e *Exception) Build(args ...any) services.Exceptions {
	e.Exceptions.Set("未知异常")
	argsNum := len(args)

	if argsNum > 0 {
		decideType(args[0], e)
	}
	if argsNum > 1 {
		decideType(args[1], e)
	}
	if argsNum > 2 {
		decideType(args[2], e)
	}
	if argsNum > 3 {
		decideType(args[3], e)
	}
	if argsNum > 4 {
		decideType(args[4], e)
	}
	if argsNum > 5 {
		decideType(args[5], e)
	}

	return e.Exceptions
}
