package providers

import (
	"gower/services"
	"gower/services/cache"
)

var _ services.CacheService = (*cache.Service)(nil)

func init() {
	P.Register("cache", []string{"config"}, func(ss ...services.Service) services.Service {
		return cache.New().Init(ss...)
	})
}
