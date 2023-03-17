package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ services.ExceptionService = (*exception.Service)(nil)

func init() {
	e := new(exceptions.Exception)
	ss.Exception = exception.Mount(e).(*exceptions.Exception)
}
