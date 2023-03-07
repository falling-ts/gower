package exceptions

import (
	"gower/services"
	"gower/services/exception"
)

// Exceptions 异常响应体
type Exceptions struct {
	*exception.Exception
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 通用错误方法
func (e *Exceptions) Error() string {
	return e.Msg
}

// Link 链接异常服务
func (e *Exceptions) Link(s services.Service) {
	s.BindAbility(e)
}

// SetException 设置异常服务
func (e *Exceptions) SetException(exception *exception.Exception) {
	e.Exception = exception
}

// SetMsg 设置异常消息
func (e *Exceptions) SetMsg(msg string) {
	e.Msg = msg
}

// SetData 设置数据
func (e *Exceptions) SetData(data any) {
	e.Data = data
}

func (e *Exceptions) throw(code uint, args ...any) *Exceptions {
	temp := *e
	newE := &temp

	newE.Exception = exception.New()
	newE.Exception.Exceptions = newE
	newE.Code = code
	return newE.Exception.Build(code, args).(*Exceptions)
}
