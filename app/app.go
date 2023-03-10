/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"gower/app/exceptions"
	"gower/app/providers"
	"gower/configs"
	"gower/services"
)

// Struct 核心结构体 app.Struct
type Struct struct {
	Name     string
	Version  string
	Services providers.ServicesMap
}

var App = Struct{
	Services: providers.Services,
}

func init() {
	cfg := Cfg()
	App.Name = cfg.App.Name
	App.Version = cfg.App.Version
}

// Run 启动 App
func Run(addr ...string) {
	if err := App.Route().Run(addr...); err != nil {
		panic(err)
	}
}

// Cfg 获取配置功能
func Cfg() *configs.All {
	if cfg, ok := App.Config().Configs().(*configs.All); ok {
		return cfg
	}

	panic("配置服务错误")
}

// Route 获取路由服务
func Route() providers.Route {
	return App.Route()
}

// Excp 获取异常功能
func Excp() *exceptions.Exception {
	if excp, ok := App.Exception().Exception().(*exceptions.Exception); ok {
		return excp
	}

	panic("异常服务错误")
}

// Get 通用获取服务方法
func (a *Struct) Get(key string) services.Service {
	if service, ok := a.Services[key]; ok {
		return service
	}

	panic("服务不存在")
}

// Config 获取配置服务
func (a *Struct) Config() providers.Config {
	if config, ok := a.Get("config").(providers.Config); ok {
		return config
	}

	panic("配置服务适配失败")
}

// Exception 获取异常服务
func (a *Struct) Exception() providers.Exception {
	if exception, ok := a.Get("exception").(providers.Exception); ok {
		return exception
	}

	panic("异常服务适配失败")
}

// Route 获取路由服务
func (a *Struct) Route() providers.Route {
	if route, ok := a.Get("route").(providers.Route); ok {
		return route
	}

	panic("路由服务适配失败")
}
