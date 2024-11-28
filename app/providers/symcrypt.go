package providers

import (
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/symcrypt"
)

var _ services.SymCryptService = (*symcrypt.Service)(nil)

func init() {
	P.Register("sym-crypt", Depends{"config"}, func(ss ...services.Service) services.Service {
		return symcrypt.New().Init(ss...)
	})
}
