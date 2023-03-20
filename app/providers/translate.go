package providers

import (
	"gower/services"
	"gower/services/translate"
	"gower/trans"
)

var _ services.TranslateService = (*translate.Service)(nil)

func init() {
	P.Register("translate", func() (Depends, Resolve) {
		return Depends{"config"}, func(ss ...services.Service) services.Service {
			return translate.Mount(trans.All).Init(ss...)
		}
	})
}
