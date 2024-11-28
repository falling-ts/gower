//go:build !cli

package providers

import (
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/logger"
)

var _ services.LoggerService = (*logger.Service)(nil)

func init() {
	P.Register("logger", func() (Depends, Resolve) {
		return Depends{"config", "util"}, func(ss ...services.Service) services.Service {
			return logger.New().Init(ss...)
		}
	})
}
