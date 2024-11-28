package providers

import (
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/translate"
	"gitee.com/falling-ts/gower/trans"
)

var _ services.TranslateService = (*translate.Service)(nil)

func init() {
	P.Register("translate", func() (Depends, Resolve) {
		return Depends{"config"}, func(ss ...services.Service) services.Service {
			return translate.Mount(trans.All).Init(ss...)
		}
	})
}
