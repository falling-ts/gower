package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/auth"
)

var _ services.AuthService = (*auth.Service)(nil)

func init() {
	P.Register("auth", func() (Depends, Resolve) {
		return Depends{"config", "util", "cache"}, func(ss ...services.Service) services.Service {
			return auth.New().Init(ss...)
		}
	})
}
