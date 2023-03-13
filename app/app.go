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

// 核心结构体
type app struct {
	Name     string
	Version  string
	Services *providers.Services
}

// App 核心实体 app.Entity
var App = app{
	Services: providers.InitServices(),
}

func init() {
	c := App.Configs()
	App.Name = c.App.Name
	App.Version = c.App.Version
}

// Run 启动 App
func Run(addr ...string) {
	if err := App.Route().Run(addr...); err != nil {
		panic(err)
	}
}

// Configs 获取配置服务
func (a *app) Configs() *configs.Configs {
	return a.Services.Configs
}

// Cache 获取缓存服务
func (a *app) Cache() services.Cache {
	return a.Services.Cache
}

// Exceptions 获取异常服务
func (a *app) Exceptions() *exceptions.Exceptions {
	return a.Services.Exceptions
}

// Route 获取路由服务
func (a *app) Route() services.Route {
	return a.Services.Route
}
