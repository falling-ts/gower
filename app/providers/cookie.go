package providers

import (
	"gower/services"
	"gower/services/cookie"
)

var _ services.CookieService = (*cookie.Service)(nil)

func init() {
	P.Register("cookie", Depends{"config", "sym-crypt"}, func(ss ...services.Service) services.Service {
		return cookie.New().Init(ss...)
	})
}
