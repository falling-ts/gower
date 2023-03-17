package console

import (
	"github.com/urfave/cli/v2"
	"gower/app"
	"time"
)

var (
	cliApp  = initCli()
	configs = app.Config()
	route   = app.Route()
)

func init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "输出软件版本.",
	}

	app.SetCli(cliApp)
}

func initCli() *cli.App {
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
