package responses

import (
	"gower/services"
	"net/http"
)

// Ok 200 成功
func (r *Response) Ok(args ...any) services.Response {
	return r.New(http.StatusOK, args...)
}
