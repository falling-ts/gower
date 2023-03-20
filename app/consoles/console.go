package consoles

import (
	"time"

	"gower/app"

	"github.com/urfave/cli/v2"
)

var (
	App     = initApp()
	configs = app.Config()
	route   = app.Route()
	util    = app.Util()
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
	var port string
	return &cli.App{
		Name:     configs.App.Name,
		Version:  configs.App.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Falling TS",
				Email: "zgh.yuanshang@gmail.com",
			},
		},
		Copyright: "(c) 2023 falling ts",
		HelpName:  configs.App.Cli,
		Usage:     "命令行工具.",
		UsageText: "辅助开发的命令工具, 在项目根目录下使用 go install 安装.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       "8080",
				Usage:       "启动应用, 并监听端口.",
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			if err := route.Run(":" + port); err != nil {
				return err
			}
			return nil
		},
		Commands: []*cli.Command{},
	}
}
