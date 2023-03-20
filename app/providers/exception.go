package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ services.ExceptionService = (*exception.Service)(nil)

func init() {
	P.Register("exception", Depends{"config", "cache"}, func(ss ...services.Service) services.Service {
		e := new(exceptions.Exception)
		return exception.Mount(e).Init(ss...)
	})
}
