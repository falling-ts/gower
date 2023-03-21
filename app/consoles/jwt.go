package consoles

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func init() {
	App.Commands = append(App.Commands, &cli.Command{
		Name:    "jwt",
		Aliases: []string{"j"},
		Usage:   "操作 JWT 各项参数",
		UsageText: `
key 初始化 JWT 秘钥
`,
		Action: func(c *cli.Context) error {
			argsNum := c.Args().Len()
			for i := 0; i < argsNum; i++ {
				return execJwt(c.Args().Get(i))
			}
			return nil
		},
	})
}

func execJwt(arg string) error {
	switch arg {
	case "key":
		key, err := util.SecretKey(64)
		if err != nil {
			return err
		}
		err = util.SetEnv("JWT_KEY", key)
		if err != nil {
			return err
		}

		fmt.Println("密钥生成成功")
	}

	return nil
}
