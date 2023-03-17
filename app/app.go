/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"gower/app/responses"
	"os"

	"gorm.io/gorm"
	"gower/app/exceptions"
	"gower/app/providers"
	"gower/configs"
	"gower/services"

	"github.com/urfave/cli/v2"
)

// 核心结构体
type app struct {
	Name     string
	Version  string
	Services *providers.Services
	Cli      *cli.App
}

// App 核心实体
var App = &app{
	Services: providers.InitServices(),
}

func init() {
	c := Config()
	App.Name = c.App.Name
	App.Version = c.App.Version
}

// Run 启动 App
func Run() {
	if err := App.Cli.Run(os.Args); err != nil {
		panic(err)
	}
}

// SetCli 设置 Cli
func SetCli(cliApp *cli.App) {
	App.Cli = cliApp
}

// Config 获取配置服务
func Config() *configs.Config {
	return App.Services.Config
}

// Cache 获取缓存服务
func Cache() services.CacheService {
	return App.Services.Cache
}

// Exception 获取异常服务
func Exception() *exceptions.Exception {
	return App.Services.Exception
}

// Route 获取路由服务
func Route() services.RouteService {
	return App.Services.Route
}

// Validator 获取验证器
func Validator() services.ValidatorService {
	return App.Services.Validator
}

// DB 获取数据库服务
func DB() services.DBService {
	return App.Services.DB
}

// GormDB 获取 gorm DB
func GormDB() *gorm.DB {
	return DB().GormDB()
}

// Response 获取响应体
func Response() *responses.Response {
	return App.Services.Response
}
