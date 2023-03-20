package providers

import (
	"gower/services"
	"gower/services/util"
)

var _ services.UtilService = (*util.Service)(nil)

func init() {
	P.Register("util", func(...services.Service) services.Service {
		return util.New().Init()
	})
}
