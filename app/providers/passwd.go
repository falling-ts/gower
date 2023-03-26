package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/passwd"
)

var _ services.PasswdService = (*passwd.Service)(nil)

func init() {
	P.Register("passwd", func() (Depends, Resolve) {
		return Depends{"config", "exception"}, func(ss ...services.Service) services.Service {
			return passwd.New().Init(ss...)
		}
	})
}
