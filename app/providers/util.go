package providers

import (
	"gower/services"
	"gower/services/util"
)

var _ services.UtilService = (*util.Service)(nil)

func init() {
	ss.Util = util.New()
}
