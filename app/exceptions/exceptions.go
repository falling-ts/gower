package exceptions

import "net/http"

// BadRequest 400 通用异常请求
func (e *Exceptions) BadRequest(args ...any) *Exceptions {
	return e.new(http.StatusBadRequest, args...)
}

// Unauthorized 401 未授权的请求
func (e *Exceptions) Unauthorized(args ...any) *Exceptions {
	return e.new(http.StatusUnauthorized, args...)
}

// Forbidden 403 请求禁止
func (e *Exceptions) Forbidden(args ...any) *Exceptions {
	return e.new(http.StatusForbidden, args...)
}

// NotFound 404 没有找到请求的资源
func (e *Exceptions) NotFound(args ...any) *Exceptions {
	return e.new(http.StatusNotFound, args...)
}

// MethodNotAllowed 405 请求方法不可行
func (e *Exceptions) MethodNotAllowed(args ...any) *Exceptions {
	return e.new(http.StatusMethodNotAllowed, args...)
}

// NotAcceptable 406 不可接受的请求
func (e *Exceptions) NotAcceptable(args ...any) *Exceptions {
	return e.new(http.StatusNotAcceptable, args...)
}

// RequestTimeout 408 请求超时
func (e *Exceptions) RequestTimeout(args ...any) *Exceptions {
	return e.new(http.StatusRequestTimeout, args...)
}

// RequestEntityTooLarge 413 请求实体太大
func (e *Exceptions) RequestEntityTooLarge(args ...any) *Exceptions {
	return e.new(http.StatusRequestEntityTooLarge, args...)
}

// RequestURITooLong 414 请求 URL 太长
func (e *Exceptions) RequestURITooLong(args ...any) *Exceptions {
	return e.new(http.StatusRequestURITooLong, args...)
}

// UnsupportedMediaType 415 请求内容类型不接受
func (e *Exceptions) UnsupportedMediaType(args ...any) *Exceptions {
	return e.new(http.StatusUnsupportedMediaType, args...)
}

// UnprocessableEntity 422 请求参数错误
func (e *Exceptions) UnprocessableEntity(args ...any) *Exceptions {
	return e.new(http.StatusUnprocessableEntity, args...)
}

// TooManyRequests 429 过多的请求
func (e *Exceptions) TooManyRequests(args ...any) *Exceptions {
	return e.new(http.StatusTooManyRequests, args...)
}
