package exceptions

import (
	"gitee.com/falling-ts/gower/services"
	"net/http"
)

// BadRequest 400 通用异常请求
func (e *Exception) BadRequest(args ...any) services.Exception {
	return e.New(http.StatusBadRequest, args...)
}

// Unauthorized 401 未授权的请求
func (e *Exception) Unauthorized(args ...any) services.Exception {
	return e.New(http.StatusUnauthorized, args...)
}

// Forbidden 403 请求禁止
func (e *Exception) Forbidden(args ...any) services.Exception {
	return e.New(http.StatusForbidden, args...)
}

// NotFound 404 没有找到请求的资源
func (e *Exception) NotFound(args ...any) services.Exception {
	return e.New(http.StatusNotFound, args...)
}

// MethodNotAllowed 405 请求方法不可行
func (e *Exception) MethodNotAllowed(args ...any) services.Exception {
	return e.New(http.StatusMethodNotAllowed, args...)
}

// NotAcceptable 406 不可接受的请求
func (e *Exception) NotAcceptable(args ...any) services.Exception {
	return e.New(http.StatusNotAcceptable, args...)
}

// RequestTimeout 408 请求超时
func (e *Exception) RequestTimeout(args ...any) services.Exception {
	return e.New(http.StatusRequestTimeout, args...)
}

// RequestEntityTooLarge 413 请求实体太大
func (e *Exception) RequestEntityTooLarge(args ...any) services.Exception {
	return e.New(http.StatusRequestEntityTooLarge, args...)
}

// RequestURITooLong 414 请求 URL 太长
func (e *Exception) RequestURITooLong(args ...any) services.Exception {
	return e.New(http.StatusRequestURITooLong, args...)
}

// UnsupportedMediaType 415 请求内容类型不接受
func (e *Exception) UnsupportedMediaType(args ...any) services.Exception {
	return e.New(http.StatusUnsupportedMediaType, args...)
}

// UnprocessableEntity 422 请求参数错误
func (e *Exception) UnprocessableEntity(args ...any) services.Exception {
	return e.New(http.StatusUnprocessableEntity, args...)
}

// TooManyRequests 429 过多的请求
func (e *Exception) TooManyRequests(args ...any) services.Exception {
	return e.New(http.StatusTooManyRequests, args...)
}
