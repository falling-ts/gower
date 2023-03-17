package providers

import (
	"gower/app/responses"
	"gower/services"
	"gower/services/response"
)

var _ services.ResponseService = (*response.Service)(nil)

func init() {
	r := new(responses.Response)
	ss.Response = response.Mount(r).(*responses.Response)
}
