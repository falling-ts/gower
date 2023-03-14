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
	c := Configs()
	App.Name = c.App.Name
	App.Version = c.App.Version
}

// Run 启动 App
func Run(addr ...string) {
	if err := Route().Run(addr...); err != nil {
		panic(err)
	}
}

// Configs 获取配置服务
func Configs() *configs.Configs {
	return App.Services.Configs
}

// Cache 获取缓存服务
func Cache() services.Cache {
	return App.Services.Cache
}

// Exceptions 获取异常服务
func Exceptions() *exceptions.Exceptions {
	return App.Services.Exceptions
}

// Route 获取路由服务
func Route() services.Route {
	return App.Services.Route
}

// Validator 获取验证器
func Validator() services.Validator {
	return App.Services.Validator
}
