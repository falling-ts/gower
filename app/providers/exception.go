package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ ExceptionService = (*exception.Exception)(nil)

type ExceptionService interface {
	services.Service

	BindContent(exceptions exception.Exceptions)
	Build(code uint, args ...any) exception.Exceptions
	Excp() exception.Exceptions
}

func buildExceptions() *exceptions.Exceptions {
	return new(exceptions.Exceptions)
}
