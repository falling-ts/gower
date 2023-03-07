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

type Application struct {
	Name    string
	Version string
	providers.Services
}

var App *Application

func init() {
	App = new(Application)
	App.Services.Mount()

	cfg := Configs()
	App.Name = cfg.App.Name
	App.Version = cfg.App.Version
}

// Run 运行系统
func Run(addr ...string) {
	if err := App.Run(addr...); err != nil {
		panic(err)
	}
}

// Configs 获取 config 配置
func Configs() configs.Configs {
	return App.Services.ConfigService.Configs()
}

// Route 获取 route 服务
func Route() providers.RouteService {
	return App.Services.RouteService
}

// Exception 获取异常服务
func Exception() exceptions.Exception {
	return App.Services.ExceptionService.ExceptionEntity()
}
