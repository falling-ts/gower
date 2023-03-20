package providers

import (
	"gower/app/responses"
	"gower/services"
	"gower/services/response"
)

var _ services.ResponseService = (*response.Service)(nil)

func init() {
	P.Register("response", func(...services.Service) services.Service {
		r := new(responses.Response)
		return response.Mount(r).Init()
	})
}
