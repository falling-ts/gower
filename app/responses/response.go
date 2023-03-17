package responses

import (
	"gower/services"
	"gower/services/response"
)

// Response 非 HTML 请求的成功响应体
type Response struct {
	*response.Service `json:"-"`
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	Data              any    `json:"data"`
}

// Set 设置响应体内容, 不设置 Code, 默认就是 0, 表示成功
func (r *Response) Set(arg any) services.Response {
	switch arg.(type) {
	case *response.Service:
		r.Service = arg.(*response.Service)
	case *Response:
		r.Service.Response = r
	case string:
		r.Msg = arg.(string)
	default:
		r.Data = arg
	}
	return r
}

func (r *Response) New(code int, args ...any) services.Response {
	temp := *r
	newR := &temp

	return newR.Set(response.New()).
		Set(newR).
		Build(code, args...)
}
