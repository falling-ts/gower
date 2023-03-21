package responses

import (
	"gower/services"
	"net/http"
)

// Ok 200 成功, 通用
func (r *Response) Ok(args ...any) services.Response {
	return r.New(http.StatusOK, args...)
}

// Created 201 创建了新资源, POST Create 方法使用
func (r *Response) Created(args ...any) services.Response {
	return r.New(http.StatusCreated, args...)
}

// Accepted 202 已接受, 正在处理
func (r *Response) Accepted(args ...any) services.Response {
	return r.New(http.StatusAccepted, args...)
}

// NonAuthoritative 203 非权威的第三方数据
func (r *Response) NonAuthoritative(args ...any) services.Response {
	return r.New(http.StatusNonAuthoritativeInfo, args...)
}

// NoContent 204 成功处理, 通常在 DELETE 请求成功后使用.
func (r *Response) NoContent(args ...any) services.Response {
	return r.New(http.StatusNoContent, args...)
}

// ResetContent 205 成功处理请求, 但需要客户端重置页面上的表单.
func (r *Response) ResetContent(args ...any) services.Response {
	return r.New(http.StatusResetContent, args...)
}
