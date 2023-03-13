package exceptions

import (
	"gower/services"
	"gower/services/exception"
)

// Exceptions 异常响应体
type Exceptions struct {
	*exception.Exception
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var _ services.Exception = (*Exceptions)(nil)
var _ services.Exceptions = (*Exceptions)(nil)

// 通用错误方法
func (e *Exceptions) Error() string {
	return e.Msg
}

// Set 通用设置内容
func (e *Exceptions) Set(arg any) {
	switch arg.(type) {
	case *exception.Exception:
		e.Exception = arg.(*exception.Exception)
	case int:
		e.Code = arg.(int)
	case string:
		e.Msg = arg.(string)
	default:
		e.Data = arg
	}
}

// Get 获取异常内容
func (e *Exceptions) Get(arg string) any {
	switch arg {
	case "RawErr":
		return e.Exception.RawErr
	default:
		return nil
	}
}

// New 抛出异常
func (e *Exceptions) New(code int, args ...any) services.Exceptions {
	return e.new(code, args...)
}

func (e *Exceptions) new(code int, args ...any) *Exceptions {
	temp := *e
	newE := &temp

	newE.Set(code)
	newE.Set(exception.New())
	newE.Exception.Exceptions = newE

	return newE.Exception.Build(args...).(*Exceptions)
}
