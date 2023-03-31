package consoles

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func init() {
	App.Commands = append(App.Commands, &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "初始化应用内容",
		UsageText: `
key 初始化秘钥
`,
		Action: func(c *cli.Context) error {
			argsNum := c.Args().Len()
			for i := 0; i < argsNum; i++ {
				err := execInit(c.Args().Get(i))
				if err != nil {
					return err
				}
			}
			return nil
		},
	})
}

func execInit(arg string) error {
	switch arg {
	case "key":
		key, err := util.SecretKey(32)
		if err != nil {
			return err
		}

		err = util.SetEnv("envs/.env.development", "APP_KEY", key)
		if err != nil {
			return err
		}

		test := "envs/.env.test"
		if util.IsExist(test) {
			err = util.SetEnv(test, "APP_KEY", key)
			if err != nil {
				return err
			}
		}

		prod := "envs/.env.production"
		if util.IsExist(prod) {
			err = util.SetEnv(prod, "APP_KEY", key)
			if err != nil {
				return err
			}
		}

		fmt.Println("APP 密钥生成成功")
	}

	return nil
}
