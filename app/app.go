/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"os"

	"github.com/falling-ts/gower/app/exceptions"
	"github.com/falling-ts/gower/app/providers"
	"github.com/falling-ts/gower/app/responses"
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
	auth, ok := Get("auth").(services.AuthService)
	if ok {
		return auth
	}

	return services.AuthService(nil)
}

// Cache 获取缓存服务
func Cache() services.CacheService {
	cache, ok := Get("cache").(services.CacheService)
	if ok {
		return cache
	}

	return services.CacheService(nil)
}

// Config 获取配置服务
func Config() *configs.Config {
	config, ok := Get("config").(*configs.Config)
	if ok {
		return config
	}

	return (*configs.Config)(nil)
}

// Cookie 获取 Cookie 服务
func Cookie() services.CookieService {
	cookie, ok := Get("cookie").(services.CookieService)
	if ok {
		return cookie
	}

	return services.CookieService(nil)
}

// DB 获取数据库服务
func DB() services.DBService {
	db, ok := Get("db").(services.DBService)
	if ok {
		return db
	}

	return services.DBService(nil)
}

// Exception 获取异常服务
func Exception() *exceptions.Exception {
	excp, ok := Get("exception").(*exceptions.Exception)
	if ok {
		return excp
	}

	return (*exceptions.Exception)(nil)
}

// Logger 获取日志服务
func Logger() services.LoggerService {
	logger, ok := Get("logger").(services.LoggerService)
	if ok {
		return logger
	}

	return services.LoggerService(nil)
}

// Passwd 获取密码服务
func Passwd() services.PasswdService {
	passwd, ok := Get("passwd").(services.PasswdService)
	if ok {
		return passwd
	}

	return services.PasswdService(nil)
}

// Response 获取响应体
func Response() *responses.Response {
	res, ok := Get("response").(*responses.Response)
	if ok {
		return res
	}

	return (*responses.Response)(nil)
}

// Route 获取路由服务
func Route() services.RouteService {
	route, ok := Get("route").(services.RouteService)
	if ok {
		return route
	}

	return services.RouteService(nil)
}

// SymCrypt 获取对称加密服务
func SymCrypt() services.SymCryptService {
	symCrypt, ok := Get("sym-crypt").(services.SymCryptService)
	if ok {
		return symCrypt
	}

	return services.SymCryptService(nil)
}

// Translate 获取翻译服务
func Translate() services.TranslateService {
	trans, ok := Get("translate").(services.TranslateService)
	if ok {
		return trans
	}

	return services.TranslateService(nil)
}

// Util 获取工具服务
func Util() services.UtilService {
	util, ok := Get("util").(services.UtilService)
	if ok {
		return util
	}

	return services.UtilService(nil)
}

// Validator 获取验证器
func Validator() services.ValidatorService {
	valid, ok := Get("validator").(services.ValidatorService)
	if ok {
		return valid
	}

	return services.ValidatorService(nil)
}

// Upload 获取上传服务
func Upload() services.UploadService {
	upload, ok := Get("upload").(services.UploadService)
	if ok {
		return upload
	}

	return services.UploadService(nil)
}
