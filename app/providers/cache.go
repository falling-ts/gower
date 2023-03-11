package providers

import (
	"gower/services"
	"gower/services/cache"
	"time"
)

var _ Cache = (*cache.Struct)(nil)

type Cache interface {
	services.Service

	SetDefault(k string, x any)
	Set(k string, x any, d time.Duration)
	Add(k string, x any, d time.Duration) error
	Replace(k string, x any, d time.Duration) error
	Increment(k string, n int64) error
	Decrement(k string, n int64) error

	Get(k string) (any, bool)
	GetWithExpiration(k string) (any, time.Time, bool)
	Flash(k string) (any, bool)

	Delete(k string)
	Flush()

	SaveFile(filename string) error
	LoadFile(filename string) error

	ItemCount() int
}

func initCache() {
	cache.Entity.Init()
	Services.Register("cache", cache.Entity)
}
