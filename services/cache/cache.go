package cache

import (
	"github.com/patrickmn/go-cache"
	"gower/services"
	"gower/services/config"
	"os"
	"path"
	"time"
)

// Struct 缓存核心结构
type Struct struct {
	*cache.Cache
}

var Entity = new(Struct)

// Init 初始化
func (c *Struct) Init(args ...any) services.Service {
	c.Cache = cache.New(
		config.Entity.Get("cache.expire", 300).(time.Duration),
		config.Entity.Get("cache.clean", 600).(time.Duration))

	interval := config.Entity.Get("cache.interval", 600).(time.Duration)
	if interval != 0 {
		dir := config.Entity.Get("cache.dir", "storage/caches").(string)
		file := config.Entity.Get("cache.file", "go.cache").(string)
		filename := path.Join(dir, file)
		if _, err := os.Stat(filename); err == nil {
			if err = c.LoadFile(filename); err != nil {
				panic(err)
			}
		}

		go func() {
			for {
				time.Sleep(interval)
				if err := c.SaveFile(filename); err != nil {
					panic(err)
				}
			}
		}()
	}

	return c
}

// Flash 闪存取值
func (c *Struct) Flash(k string) (any, bool) {
	value, ok := c.Get(k)
	c.Delete(k)
	return value, ok
}
