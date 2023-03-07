package exceptions

import "net/http"

func (e *Exception) BadRequest(args ...any) *Exception {
	return e.throw(http.StatusBadRequest, args...)
}
