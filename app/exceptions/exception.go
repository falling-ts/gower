package exceptions

import (
	"gower/services"
	"gower/services/exception"
)

var _ services.Exception = (*Exception)(nil)
var _ services.ExceptionService = (*Exception)(nil)

// Exception 异常响应体
type Exception struct {
	*exception.Service `json:"-"`
	Code               int    `json:"code"`
	Msg                string `json:"msg"`
	Data               any    `json:"data"`
}

// Set 通用设置内容
func (e *Exception) Set(arg any) {
	switch arg.(type) {
	case *exception.Service:
		e.Service = arg.(*exception.Service)
	case int:
		e.Code = arg.(int)
	case string:
		e.Msg = arg.(string)
	default:
		e.Data = arg
	}
}

// 通用错误方法
func (e *Exception) Error() string {
	return e.Msg
}

// Get 获取异常内容
func (e *Exception) Get(arg string) any {
	switch arg {
	case "RawErr":
		return e.Service.RawErr
	default:
		return nil
	}
}

// New 抛出异常
func (e *Exception) New(code int, args ...any) services.Exception {
	return e.new(code, args...)
}

func (e *Exception) new(code int, args ...any) *Exception {
	temp := *e
	newE := &temp

	newE.Set(code)
	newE.Set(exception.New())
	newE.Service.Exception = newE

	return newE.Service.Build(args...).(*Exception)
}
