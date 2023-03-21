package responses

import (
	"errors"
	"reflect"

	"gower/services"
	"gower/services/response"
	"gower/utils/str"
)

// Response 非 HTML 请求的成功响应体
type Response struct {
	*response.Service `json:"-"`
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	Data              any    `json:"data"`
	Token             string `json:"token"`
}

// Set 设置响应体内容, 不设置 Code, 默认就是 0, 表示成功
func (r *Response) Set(arg any) services.Response {
	switch arg.(type) {
	case *response.Service:
		r.Service = arg.(*response.Service)
	case *Response:
		r.Service.Response = r
	case string:
		s := arg.(string)
		if r.IsToken(s) {
			r.Token = s
			break
		}
		r.Msg = s
	default:
		r.Data = arg
	}
	return r
}

// Get 获取内容
func (r *Response) Get(field string) (any, error) {
	res := reflect.ValueOf(r).Elem()
	value := res.FieldByName(str.Conv(field).Uppercase())
	if value.IsValid() {
		return value.Interface(), nil
	}

	return nil, errors.New("无效字段")
}

// New 创建新响应
func (r *Response) New(code int, args ...any) services.Response {
	temp := *r
	newR := &temp

	return newR.Set(response.New()).
		Set(newR).
		Build(code, args...)
}
