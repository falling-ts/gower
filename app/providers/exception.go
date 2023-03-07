package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ ExceptionService = (*exception.Exception)(nil)

type ExceptionService interface {
	services.Service

	Clone(code uint, args ...any) ExceptionService
	Excp() *exceptions.Exception
}
