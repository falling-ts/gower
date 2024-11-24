package consoles

import (
	"bufio"
	"fmt"
	"github.com/falling-ts/gower/utils/str"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"strings"
	"text/template"
)

func init() {
	App.Commands = append(App.Commands, &cli.Command{
		Name:    "make",
		Aliases: []string{"m"},
		Usage:   "初始化创建文件",
		UsageText: `
例如:
$ gower make TestAa TestAb TestAc

创建 3 套 Web 控制器, 请求, 模型
`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "rest",
				Aliases: []string{"r"},
				Usage:   "创建 RestApi 控制器, 以及请求模型.",
				Action: func(c *cli.Context, r string) error {
					var err error
					args := strings.Split(r, ",")
					argsNum := len(args)
					for i := 0; i < argsNum; i++ {
						content := args[i]
						if err = makeControllerRest(content); err != nil {
							return err
						}
						if err = makeRequest(content); err != nil {
							return err
						}
						if err = makeModel(content); err != nil {
							return err
						}

						fmt.Println(str.Conv(content).UpCamel() + " Rest 控制器,请求,模型创建成功")
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:    "api",
				Aliases: []string{"i"},
				Usage:   "创建 Api 控制器, 以及请求模型.",
				Action: func(c *cli.Context, api string) error {
					var err error
					args := strings.Split(api, ",")
					argsNum := len(args)
					for i := 0; i < argsNum; i++ {
						content := args[i]
						if err = makeApiController(content); err != nil {
							return err
						}
						if err = makeApiRequest(content); err != nil {
							return err
						}
						if err = makeModel(content); err != nil {
							return err
						}

						fmt.Println(str.Conv(content).UpCamel() + " Api 控制器,请求,模型创建成功")
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:    "admin",
				Aliases: []string{"a"},
				Usage:   "创建 Admin 控制器, 以及请求模型.",
				Action: func(c *cli.Context, api string) error {
					var err error
					args := strings.Split(api, ",")
					argsNum := len(args)
					for i := 0; i < argsNum; i++ {
						content := args[i]
						if err = makeAdminController(content); err != nil {
							return err
						}
						if err = makeAdminRequest(content); err != nil {
							return err
						}
						if err = makeModel(content); err != nil {
							return err
						}

						fmt.Println(str.Conv(content).UpCamel() + " Admin 控制器,请求,模型创建成功")
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:    "admin:rest",
				Aliases: []string{"ar"},
				Usage:   "创建 Admin Rest 控制器, 以及请求模型.",
				Action: func(c *cli.Context, api string) error {
					var err error
					args := strings.Split(api, ",")
					argsNum := len(args)
					for i := 0; i < argsNum; i++ {
						content := args[i]
						if err = makeAdminControllerRest(content); err != nil {
							return err
						}
						if err = makeAdminRequest(content); err != nil {
							return err
						}
						if err = makeModel(content); err != nil {
							return err
						}

						fmt.Println(str.Conv(content).UpCamel() + " Admin Rest 控制器,请求,模型创建成功")
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:    "controller",
				Aliases: []string{"c"},
				Usage:   "创建 Web 控制器.",
				Action: func(ctx *cli.Context, c string) error {
					if err := makeController(c); err != nil {
						return err
					}

					fmt.Println(str.Conv(c).UpCamel() + "Controller 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "controller-rest",
				Aliases: []string{"cr"},
				Usage:   "创建 Web Rest 控制器.",
				Action: func(ctx *cli.Context, c string) error {
					if err := makeControllerRest(c); err != nil {
						return err
					}

					fmt.Println(str.Conv(c).UpCamel() + "Controller 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "api:controller",
				Aliases: []string{"ic"},
				Usage:   "创建 Api Rest 控制器.",
				Action: func(ctx *cli.Context, c string) error {
					if err := makeApiController(c); err != nil {
						return err
					}

					fmt.Println(str.Conv(c).UpCamel() + "Controller 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "admin:controller",
				Aliases: []string{"ac"},
				Usage:   "创建 Admin 控制器.",
				Action: func(ctx *cli.Context, c string) error {
					if err := makeAdminController(c); err != nil {
						return err
					}

					fmt.Println(str.Conv(c).UpCamel() + " Admin Controller 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "admin:controller-rest",
				Aliases: []string{"acr"},
				Usage:   "创建 Admin Rest 控制器.",
				Action: func(ctx *cli.Context, c string) error {
					if err := makeAdminControllerRest(c); err != nil {
						return err
					}

					fmt.Println(str.Conv(c).UpCamel() + " Admin Controller 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "request",
				Aliases: []string{"req"},
				Usage:   "创建 Web 请求.",
				Action: func(ctx *cli.Context, req string) error {
					if err := makeRequest(req); err != nil {
						return err
					}

					fmt.Println(str.Conv(req).UpCamel() + "Request 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "api:request",
				Aliases: []string{"ir"},
				Usage:   "创建 Api 请求.",
				Action: func(ctx *cli.Context, req string) error {
					if err := makeApiRequest(req); err != nil {
						return err
					}

					fmt.Println(str.Conv(req).UpCamel() + "Request 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "admin:request",
				Aliases: []string{"are"},
				Usage:   "创建 Admin 请求.",
				Action: func(ctx *cli.Context, req string) error {
					if err := makeAdminRequest(req); err != nil {
						return err
					}

					fmt.Println(str.Conv(req).UpCamel() + "Request 创建成功")
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "model",
				Aliases: []string{"m"},
				Usage:   "创建模型.",
				Action: func(ctx *cli.Context, m string) error {
					if err := makeModel(m); err != nil {
						return err
					}

					fmt.Println(str.Conv(m).UpCamel() + " Model 创建成功")
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			var err error

			argsNum := c.Args().Len()
			for i := 0; i < argsNum; i++ {
				content := c.Args().Get(i)
				if err = makeController(content); err != nil {
					return err
				}
				if err = makeRequest(content); err != nil {
					return err
				}
				if err = makeModel(content); err != nil {
					return err
				}

				fmt.Println(str.Conv(content).UpCamel() + " 控制器,请求,模型创建成功")
			}

			return nil
		},
	})
}

func makeController(c string) error {
	file, err := gowerMake(c, "app/http/controllers", "_controller.go", "make/controller.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeRequest(r string) error {
	file, err := gowerMake(r, "app/http/requests", "_request.go", "make/request.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeModel(m string) error {
	file, err := gowerMake(m, "app/models", ".go", "make/model.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeControllerRest(c string) error {
	file, err := gowerMake(c, "app/http/controllers", "_controller.go", "make/controller_rest.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeApiController(c string) error {
	file, err := gowerMake(c, "app/api/controllers", "_controller.go", "make/api.controller.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeApiRequest(r string) error {
	file, err := gowerMake(r, "app/api/requests", "_request.go", "make/request.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeAdminController(c string) error {
	file, err := gowerMake(c, "app/admin/controllers", "_controller.go", "make/admin.controller.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeAdminControllerRest(c string) error {
	file, err := gowerMake(c, "app/admin/controllers", "_controller.go", "make/admin.controller_rest.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func makeAdminRequest(r string) error {
	file, err := gowerMake(r, "app/admin/requests", "_request.go", "make/request.go.tpl")
	if err != nil {
		_ = os.Remove(file)
		return err
	}

	return nil
}

func gowerMake(content, dirStr, suffix, tplFile string) (string, error) {
	conv := str.Conv(content)
	dir := util.CreateDir(dirStr)
	filename := conv.Snake() + suffix
	file := path.Join(dir, filename)
	if util.IsExist(file) {
		return "", nil
	}

	f, err := os.Create(file)
	if err != nil {
		return file, err
	}

	defer func() {
		_ = f.Close()
	}()

	data, err := tplFS.ReadFile(tplFile)
	if err != nil {
		return file, err
	}
	tpl, err := template.New(content).Parse(string(data))
	if err != nil {
		return file, err
	}

	module, err := readModule()
	if err != nil {
		return file, err
	}

	if err = tpl.Execute(f, map[string]any{
		"UpCamel": conv.UpCamel(),
		"Module":  module,
	}); err != nil {
		return file, err
	}

	return "", nil
}

func readModule() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module") {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				return parts[1], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("go.mod not found")
}
