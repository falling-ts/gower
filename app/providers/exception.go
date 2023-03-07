package providers

import (
	"gower/services"
	"gower/services/exception"
)

var _ ExceptionService = (*exception.Exception)(nil)

type ExceptionService interface {
	services.Service
	Build(code uint, args ...any) exception.Exceptions
	Excp() exception.Exceptions
}
