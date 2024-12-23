//go:build !cli

package providers

import (
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/route"
)

var _ services.RouteService = (*route.Service)(nil)

func init() {
	P.Register("route", func() (Depends, Resolve) {
		return Depends{"config", "exception", "db", "response", "util"}, func(ss ...services.Service) services.Service {
			return route.New().Init(ss...)
		}
	})
}
