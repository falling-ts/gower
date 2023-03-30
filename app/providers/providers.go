package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/utils/slice"
)

// Depends 服务依赖
type Depends []string

// Resolve 服务提供者
type Resolve func(...services.Service) services.Service

// 服务注册器
type register struct {
	key     string
	depends Depends
	resolve func(...services.Service) services.Service
	handler func() (Depends, Resolve)
}

// 设置注册器
func (r *register) set(arg any) {
	switch arg.(type) {
	case string:
		r.key = arg.(string)
	case Depends:
		r.depends = arg.(Depends)
	case []string:
		r.depends = arg.([]string)
	case func(...services.Service) services.Service:
		r.resolve = arg.(func(...services.Service) services.Service)
	case func() (Depends, Resolve):
		r.handler = arg.(func() (Depends, Resolve))
	}
}

type providers map[string]any
type caches map[string]services.Service

// Provider 服务提供者
type Provider struct {
	providers
	caches
}

var P = &Provider{
	make(providers),
	make(caches),
}

// Register 登记服务提供者
func (p *Provider) Register(args ...any) {
	r := new(register)
	for _, arg := range args {
		r.set(arg)
	}

	if r.key == "" {
		panic("请设置服务标识符")
	}
	if r.handler != nil {
		p.providers[r.key] = r.handler
		return
	}

	if r.resolve == nil {
		panic("请设置服务提供者")
	}
	if len(r.depends) == 0 {
		p.providers[r.key] = r.resolve
	} else {
		p.providers[r.key] = func() (Depends, Resolve) {
			return r.depends, r.resolve
		}
	}
}

// Get 获取服务
func (p *Provider) Get(key string) services.Service {
	if s, ok := p.caches[key]; ok {
		return s
	}

	var fn any
	var ok bool
	if fn, ok = p.providers[key]; !ok {
		return services.Service(nil)
	}

	switch fn.(type) {
	case func(...services.Service) services.Service:
		return p.Cache(key, fn.(func(...services.Service) services.Service)())
	case func() (Depends, Resolve):
		depends, resolve := fn.(func() (Depends, Resolve))()
		ss := make([]services.Service, len(depends))
		for i, depend := range depends {
			ss[i] = p.Get(depend)
		}

		return p.Cache(key, resolve(ss...))
	}

	panic("服务注册异常")
}

// Cache 服务缓存
func (p *Provider) Cache(key string, s services.Service) services.Service {
	p.caches[key] = s
	return s
}

// Clear 清空缓存
func (p *Provider) Clear() {
	p.caches = make(map[string]services.Service)
}

// Del 删除指定缓存
func (p *Provider) Del(keys ...string) {
	for _, key := range keys {
		delete(p.caches, key)
	}
}

// DelExcept 删除指定除外的缓存
func (p *Provider) DelExcept(keys ...string) {
	for key, _ := range p.caches {
		if !slice.Strings(keys).Has(key) {
			delete(p.caches, key)
		}
	}
}
