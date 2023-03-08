package exception

import (
	"sync"

	"gower/services"
)

// Exceptions 异常内容
type Exceptions interface {
	SetException(exception *Exception)
	SetMsg(msg string)
	SetData(data any)
	Throw(code uint, args ...any) Exceptions
}

// Exception 异常主结构体
type Exception struct {
	Exceptions
	RawErr error
}

var (
	exception *Exception
	once      sync.Once
)

// Build 构建单例模式
func Build() *Exception {
	once.Do(func() {
		build()
	})

	return exception
}

// New 创建新异常服务
func New() *Exception {
	return new(Exception)
}

// Register 注册服务
func (e *Exception) Register(s services.Services) {
	s.SetService(e)
}

// BindContent 绑定异常内容
func (e *Exception) BindContent(exceptions Exceptions) {
	e.Exceptions = exceptions
	e.Exceptions.SetException(e)
}

// Build 构建每个请求的异常
func (e *Exception) Build(code uint, args ...any) Exceptions {
	e.Exceptions.SetMsg("未知异常")
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

// Excp 获取异常实体
func (e *Exception) Excp() Exceptions {
	return e.Exceptions
}

func build() {
	exception = new(Exception)
}
