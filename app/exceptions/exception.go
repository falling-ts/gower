package exceptions

import (
	"errors"
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/exception"
	"gitee.com/falling-ts/gower/utils/str"
	"reflect"
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
func (e *Exception) Set(arg any) services.Exception {
	switch arg.(type) {
	case *exception.Service:
		e.Service = arg.(*exception.Service)
	case *Exception:
		e.Service.Exception = e
	case int:
		e.Code = arg.(int)
	case string:
		e.Msg = arg.(string)
	default:
		e.Data = arg
	}

	return e
}

// Get 获取内容
func (e *Exception) Get(field string) (any, error) {
	res := reflect.ValueOf(e).Elem()
	value := res.FieldByName(str.Conv(field).Uppercase())
	if value.IsValid() {
		return value.Interface(), nil
	}

	return nil, errors.New("无效字段")
}

// 通用错误方法
func (e *Exception) Error() string {
	return e.Msg
}

// New 创建异常
func (e *Exception) New(code int, args ...any) services.Exception {
	temp := *e
	newE := &temp

	return newE.Set(exception.New()).
		Set(newE).
		Build(args...).
		Set(code)
}
