package cache

import (
	"os"
	"path"
	"time"

	"gower/services"

	"github.com/patrickmn/go-cache"
)

// Service 服务
type Service struct {
	*cache.Cache
}

var config services.Config

// New 创建服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...any) {
	if len(args) == 0 {
		panic("缓存服务初始化参数不存在.")
	}
	config = args[0].(services.Config)

	s.Cache = cache.New(
		config.Get("cache.expire", "300s").(time.Duration),
		config.Get("cache.clean", "600s").(time.Duration))

	interval := config.Get("cache.interval", "600s").(time.Duration)
	if interval != 0 {
		dir := config.Get("cache.dir", "storage/caches").(string)
		file := config.Get("cache.file", "go.cache").(string)
		filename := path.Join(dir, file)
		if _, err := os.Stat(filename); err == nil {
			if err = s.LoadFile(filename); err != nil {
				panic(err)
			}
		}

		go func() {
			for {
				time.Sleep(interval)
				if err := s.SaveFile(filename); err != nil {
					panic(err)
				}
			}
		}()
	}
}

// Flash 闪存取值
func (s *Service) Flash(k string) (any, bool) {
	value, ok := s.Get(k)
	s.Delete(k)
	return value, ok
}
