package consoles

import "github.com/urfave/cli/v2"

func init() {
	var port string
	App.Commands = append(App.Commands, &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "启动应用",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       "8080",
				Usage:       "启动应用, 并监听端口.",
				Destination: &port,
			},
		},
		Action: func(*cli.Context) error {
			if err := route.Run(":" + port); err != nil {
				return err
			}
			return nil
		},
	})
}
