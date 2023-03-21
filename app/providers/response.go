package providers

import (
	"gower/responses"
	"gower/services"
	"gower/services/response"
)

var _ services.ResponseService = (*response.Service)(nil)

func init() {
	P.Register("response", Depends{"config", "token"}, func(ss ...services.Service) services.Service {
		r := new(responses.Response)
		return response.Mount(r).Init(ss...)
	})
}
