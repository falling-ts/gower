package consoles

import (
	"time"

	"gitee.com/falling-ts/gower/app"

	"github.com/urfave/cli/v2"
)

var (
	App    = initApp()
	config = app.Config()
	route  = app.Route()
	util   = app.Util()
)

func init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "输出软件版本.",
	}

	app.SetApp(App)
}

func initApp() *cli.App {
	return &cli.App{
		Name:     config.App.Name,
		Version:  config.App.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Falling TS",
				Email: "zgh.yuanshang@gmail.com",
			},
		},
		Copyright: "(c) 2023 falling ts",
		HelpName:  config.App.Cli,
		Usage:     "命令行工具.",
		UsageText: "辅助开发的命令工具, 在项目根目录下使用 go install 安装.",
		Commands:  []*cli.Command{},
	}
}
