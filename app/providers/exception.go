package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ services.Exception = (*exception.Exception)(nil)

func init() {
	e := new(exceptions.Exceptions)
	s.Exceptions = exception.Mount(e).(*exceptions.Exceptions)
}
