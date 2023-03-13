package cache

import (
	"os"
	"path"
	"time"

	"gower/services"

	"github.com/patrickmn/go-cache"
)

// Cache 服务
type Cache struct {
	*cache.Cache
}

var configs services.Configs

// New 创建服务
func New() *Cache {
	return new(Cache)
}

// Init 初始化
func (c *Cache) Init(args ...any) {
	if len(args) == 0 {
		panic("缓存服务初始化参数不存在.")
	}
	configs = args[0].(services.Configs)

	c.Cache = cache.New(
		configs.Get("cache.expire", 300).(time.Duration),
		configs.Get("cache.clean", 600).(time.Duration))

	interval := configs.Get("cache.interval", 600).(time.Duration)
	if interval != 0 {
		dir := configs.Get("cache.dir", "storage/caches").(string)
		file := configs.Get("cache.file", "go.cache").(string)
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
}

// Flash 闪存取值
func (c *Cache) Flash(k string) (any, bool) {
	value, ok := c.Get(k)
	c.Delete(k)
	return value, ok
}
