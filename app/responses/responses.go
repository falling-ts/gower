package responses

import "gower/services"

// Responses 非 HTML 请求的成功响应体
type Responses struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Set 设置响应体内容
func (r *Responses) Set(arg any) services.Responses {
	switch arg.(type) {
	case int:
		r.Code = arg.(int)
	case string:
		r.Msg = arg.(string)
	default:
		r.Data = arg
	}
	return r
}
