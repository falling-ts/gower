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
)

type App struct {
	Name    string
	Version string
	providers.Services
}

var Gower *App

func init() {
	Gower = new(App)
	Gower.Services.Mount().BindContent()

	cfg := Cfg()
	Gower.Name = cfg.App.Name
	Gower.Version = cfg.App.Version
}

// Run 启动 App
func Run(addr ...string) {
	if err := Gower.Run(addr...); err != nil {
		panic(err)
	}
}

// Cfg 获取配置功能
func Cfg() *configs.Configs {
	if cfg, ok := Gower.Config().Cfg().(*configs.Configs); ok {
		return cfg
	}

	panic("配置服务错误")
}

// Route 获取路由服务
func Route() providers.RouteService {
	return Gower.Route()
}

// Excp 获取异常功能
func Excp() *exceptions.Exceptions {
	if excp, ok := Gower.Exception().Excp().(*exceptions.Exceptions); ok {
		return excp
	}

	panic("异常服务错误")
}

// Config 获取配置服务
func (a *App) Config() providers.ConfigService {
	return a.Services.ConfigService
}

// Route 获取路由服务
func (a *App) Route() providers.RouteService {
	return a.Services.RouteService
}

// Exception 获取异常服务
func (a *App) Exception() providers.ExceptionService {
	return a.Services.ExceptionService
}
