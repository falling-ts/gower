package providers

import (
	"gitee.com/falling-ts/gower/app/exceptions"
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/exception"
)

var _ services.ExceptionService = (*exception.Service)(nil)

func init() {
	P.Register("exception", Depends{"config", "cache", "util", "cookie"}, func(ss ...services.Service) services.Service {
		e := new(exceptions.Exception)
		return exception.Mount(e).Init(ss...)
	})
}
