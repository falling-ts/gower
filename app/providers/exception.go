package providers

import (
	"gower/app/exceptions"
	"gower/services"
	"gower/services/exception"
)

var _ Exception = (*exception.Struct)(nil)

// Exception 适配接口
type Exception interface {
	services.Service

	Build(code uint, args ...any) exception.Content
	Exception() exception.Content
	HandleBy(any)
}

func initException() {
	e := new(exceptions.Exception)
	exception.Entity.Init(e)

	Services.Register("exception", exception.Entity)
}
