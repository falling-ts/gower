package providers

import (
	"gitee.com/falling-ts/gower/services"
	"gitee.com/falling-ts/gower/services/cache"
)

var _ services.CacheService = (*cache.Service)(nil)

func init() {
	P.Register("cache", []string{"config"}, func(ss ...services.Service) services.Service {
		return cache.New().Init(ss...)
	})
}
