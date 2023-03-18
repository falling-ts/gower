package providers

import (
	"gower/services"
	"gower/services/logger"
)

var _ services.LoggerService = (*logger.Service)(nil)

func init() {
	ss.Logger = logger.New()
}
