package providers

import (
	"gower/services"
	"gower/services/auth"
)

var _ services.AuthService = (*auth.Service)(nil)

func init() {
	P.Register("auth", func() (Depends, Resolve) {
		return Depends{"config", "util", "cache"}, func(ss ...services.Service) services.Service {
			return auth.New().Init(ss...)
		}
	})
}
