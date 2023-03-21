package providers

import (
	"gower/services"
	"gower/services/token"
)

var _ services.TokenService = (*token.Service)(nil)

func init() {
	P.Register("token", func() (Depends, Resolve) {
		return Depends{"config", "util", "cache"}, func(ss ...services.Service) services.Service {
			return token.New().Init(ss...)
		}
	})
}
