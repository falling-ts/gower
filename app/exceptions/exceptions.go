package exceptions

import "net/http"

// BadRequest 400 通用异常请求
func (e *Exception) BadRequest(args ...any) *Exception {
	return e.new(http.StatusBadRequest, args...)
}

// Unauthorized 401 未授权的请求
func (e *Exception) Unauthorized(args ...any) *Exception {
	return e.new(http.StatusUnauthorized, args...)
}

// Forbidden 403 请求禁止
func (e *Exception) Forbidden(args ...any) *Exception {
	return e.new(http.StatusForbidden, args...)
}

// NotFound 404 没有找到请求的资源
func (e *Exception) NotFound(args ...any) *Exception {
	return e.new(http.StatusNotFound, args...)
}

// MethodNotAllowed 405 请求方法不可行
func (e *Exception) MethodNotAllowed(args ...any) *Exception {
	return e.new(http.StatusMethodNotAllowed, args...)
}

// NotAcceptable 406 不可接受的请求
func (e *Exception) NotAcceptable(args ...any) *Exception {
	return e.new(http.StatusNotAcceptable, args...)
}

// RequestTimeout 408 请求超时
func (e *Exception) RequestTimeout(args ...any) *Exception {
	return e.new(http.StatusRequestTimeout, args...)
}

// RequestEntityTooLarge 413 请求实体太大
func (e *Exception) RequestEntityTooLarge(args ...any) *Exception {
	return e.new(http.StatusRequestEntityTooLarge, args...)
}

// RequestURITooLong 414 请求 URL 太长
func (e *Exception) RequestURITooLong(args ...any) *Exception {
	return e.new(http.StatusRequestURITooLong, args...)
}

// UnsupportedMediaType 415 请求内容类型不接受
func (e *Exception) UnsupportedMediaType(args ...any) *Exception {
	return e.new(http.StatusUnsupportedMediaType, args...)
}

// UnprocessableEntity 422 请求参数错误
func (e *Exception) UnprocessableEntity(args ...any) *Exception {
	return e.new(http.StatusUnprocessableEntity, args...)
}

// TooManyRequests 429 过多的请求
func (e *Exception) TooManyRequests(args ...any) *Exception {
	return e.new(http.StatusTooManyRequests, args...)
}
