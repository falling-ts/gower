package exceptions

import "gower/services/exception"

// Exception 异常响应体
type Exception struct {
	*exception.Struct
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var _ exception.Content = (*Exception)(nil)

// 通用错误方法
func (e *Exception) Error() string {
	return e.Msg
}

// SetException 设置异常服务
func (e *Exception) SetException(exception *exception.Struct) {
	e.Struct = exception
}

// SetMsg 设置异常消息
func (e *Exception) SetMsg(msg string) {
	e.Msg = msg
}

// SetData 设置数据
func (e *Exception) SetData(data any) {
	e.Data = data
}

// Throw 抛出异常
func (e *Exception) Throw(code uint, args ...any) exception.Content {
	return e.throw(code, args...)
}

// HandleBy 处理异常
func (e *Exception) HandleBy(arg any) {
	e.Struct.HandleBy(arg)
}

func (e *Exception) throw(code uint, args ...any) *Exception {
	temp := *e
	newE := &temp

	newE.Struct = exception.New()
	newE.Struct.Content = newE
	newE.Code = code
	return newE.Struct.Build(code, args...).(*Exception)
}
