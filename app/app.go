/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"os"

	"gower/app/exceptions"
	"gower/app/providers"
	"gower/app/responses"
	"gower/configs"
	"gower/services"

	"github.com/urfave/cli/v2"
)

// 核心结构体
type app struct {
	*cli.App
	*providers.Provider

	Name    string
	Version string
}

// App 核心实体
var App = new(app)

func init() {
	App.Provider = providers.P

	c := Config()
	App.Name = c.App.Name
	App.Version = c.App.Version
}

// Run 启动 App
func Run() {
	if err := App.Run(os.Args); err != nil {
		panic(err)
	}
}

// SetApp 设置 Cli App
func SetApp(cliApp *cli.App) {
	App.App = cliApp
}

// Get 通用获取服务方法
func Get(key string) services.Service {
	return App.Get(key)
}

// Config 获取配置服务
func Config() *configs.Config {
	return Get("config").(*configs.Config)
}

// Cache 获取缓存服务
func Cache() services.CacheService {
	return Get("cache").(services.CacheService)
}

// Exception 获取异常服务
func Exception() *exceptions.Exception {
	return Get("exception").(*exceptions.Exception)
}

// Route 获取路由服务
func Route() services.RouteService {
	return Get("route").(services.RouteService)
}

// Validator 获取验证器
func Validator() services.ValidatorService {
	return Get("validator").(services.ValidatorService)
}

// DB 获取数据库服务
func DB() services.DBService {
	return Get("db").(services.DBService)
}

// Response 获取响应体
func Response() *responses.Response {
	return Get("response").(*responses.Response)
}

// Logger 获取日志服务
func Logger() services.LoggerService {
	return Get("logger").(services.LoggerService)
}

// Passwd 获取密码服务
func Passwd() services.PasswdService {
	return Get("passwd").(services.PasswdService)
}

// Util 获取工具服务
func Util() services.UtilService {
	return Get("util").(services.UtilService)
}

// Translate 获取翻译服务
func Translate() services.TranslateService {
	return Get("translate").(services.TranslateService)
}
