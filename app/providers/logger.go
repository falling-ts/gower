package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/logger"
)

var _ services.LoggerService = (*logger.Service)(nil)

func init() {
	P.Register("logger", func() (Depends, Resolve) {
		return Depends{"config"}, func(ss ...services.Service) services.Service {
			return logger.New().Init(ss...)
		}
	})
}
