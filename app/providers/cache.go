package providers

import (
	"gower/services"
	"gower/services/cache"
)

var _ services.Cache = (*cache.Cache)(nil)

func init() {
	s.Cache = cache.New()
}
