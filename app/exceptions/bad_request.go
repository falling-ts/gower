package exceptions

import "net/http"

func (e *Exceptions) BadRequest(args ...any) *Exceptions {
	return e.new(http.StatusBadRequest, args...)
}