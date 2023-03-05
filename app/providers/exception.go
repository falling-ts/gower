package providers

import (
	"gower/services"
	"gower/services/exception"
)

var _ ExceptionService = (*exception.Exception)(nil)

type ExceptionService interface {
	Exception() // 用以区分其它服务
	services.Service
}
