package providers

import (
	"github.com/falling-ts/gower/app/responses"
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/response"
)

var _ services.ResponseService = (*response.Service)(nil)

func init() {
	P.Register("response", Depends{"auth", "cookie", "util", "config", "db", "cache", "exception"}, func(ss ...services.Service) services.Service {
		r := new(responses.Response)
		return response.Mount(r).Init(ss...)
	})
}
