/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"github.com/falling-ts/gower/app/responses"
	"os"

	"github.com/falling-ts/gower/app/exceptions"
	"github.com/falling-ts/gower/app/providers"
	"github.com/falling-ts/gower/configs"
	"github.com/falling-ts/gower/services"

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
var App = &app{
	Provider: providers.P,
}

func init() {
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

// Auth 获取 Auth 服务
func Auth() services.AuthService {
	return Get("auth").(services.AuthService)
}

// Cache 获取缓存服务
func Cache() services.CacheService {
	return Get("cache").(services.CacheService)
}

// Config 获取配置服务
func Config() *configs.Config {
	return Get("config").(*configs.Config)
}

// Cookie 获取 Cookie 服务
func Cookie() services.CookieService {
	return Get("cookie").(services.CookieService)
}

// DB 获取数据库服务
func DB() services.DBService {
	return Get("db").(services.DBService)
}

// Exception 获取异常服务
func Exception() *exceptions.Exception {
	return Get("exception").(*exceptions.Exception)
}

// Logger 获取日志服务
func Logger() services.LoggerService {
	return Get("logger").(services.LoggerService)
}

// Passwd 获取密码服务
func Passwd() services.PasswdService {
	return Get("passwd").(services.PasswdService)
}

// Response 获取响应体
func Response() *responses.Response {
	return Get("response").(*responses.Response)
}

// Route 获取路由服务
func Route() services.RouteService {
	return Get("route").(services.RouteService)
}

// SymCrypt 获取对称加密服务
func SymCrypt() services.SymCryptService {
	return Get("sym-crypt").(services.SymCryptService)
}

// Translate 获取翻译服务
func Translate() services.TranslateService {
	return Get("translate").(services.TranslateService)
}

// Util 获取工具服务
func Util() services.UtilService {
	return Get("util").(services.UtilService)
}

// Validator 获取验证器
func Validator() services.ValidatorService {
	return Get("validator").(services.ValidatorService)
}
