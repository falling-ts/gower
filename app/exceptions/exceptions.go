package exceptions

import (
	"gower/services/exception"
)

type Exception struct {
	*exception.Exception
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 通用错误方法
func (e *Exception) Error() string {
	return e.Msg
}

func (e *Exception) throw(code uint, args ...any) *Exception {
	newE := e.Clone(code, args...)
	return newE.Excp()
}
