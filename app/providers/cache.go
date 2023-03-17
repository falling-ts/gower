package providers

import (
	"gower/services"
	"gower/services/cache"
)

var _ services.CacheService = (*cache.Service)(nil)

func init() {
	ss.Cache = cache.New()
}
