package consoles

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	App.Commands = append(App.Commands, &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "创建项目",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return nil
			}

			return initProject(c.Args().Get(0))
		},
	})
}

func initProject(project string) error {
	err := create(project)
	if err != nil {
		return err
	}

	err = initEnv(project)
	if err != nil {
		return err
	}

	err = initEnv(filepath.Join(project, "envs"))
	if err != nil {
		return err
	}

	err = initKey(project)
	if err != nil {
		return err
	}

	err = jwtKey(project)
	if err != nil {
		return err
	}

	err = goModTidy(project)
	if err != nil {
		return err
	}

	err = pnpmInstall(project)
	if err != nil {
		return err
	}

	err = initGit(project)
	if err != nil {
		return err
	}

	err = addAll(project)
	if err != nil {
		return err
	}

	err = commitM(project)
	if err != nil {
		return err
	}

	err = buildDev(project)
	if err != nil {
		return err
	}

	err = execTest(project)
	if err != nil {
		return err
	}

	err = overrideGower(project)
	if err != nil {
		return err
	}

	return nil
}

func create(project string) error {
	if util.IsExist(project) {
		return errors.New("目录已存在")
	}

	file, err := gower.Open("create/gower.zip")
	if err != nil {
		return err
	}
	defer func(file fs.File) {
		_ = file.Close()
	}(file)

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, file); err != nil {
		return err
	}

	reader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		err = func() error {
			path := filepath.Join(project, file.Name)

			if file.FileInfo().IsDir() {
				err = os.MkdirAll(path, os.ModePerm)
				if err != nil {
					return err
				}
			} else {
				if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
					return err
				}

				destFile, err := os.Create(path)
				if err != nil {
					return err
				}
				defer func(destFile *os.File) {
					_ = destFile.Close()
				}(destFile)

				srcFile, err := file.Open()
				if err != nil {
					return err
				}
				defer func(srcFile io.ReadCloser) {
					_ = srcFile.Close()
				}(srcFile)

				if _, err = io.Copy(destFile, srcFile); err != nil {
					return err
				}
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	fmt.Printf("%s 项目创建成功\n", project)
	return nil
}

func initEnv(dir string) error {
	dev := filepath.Join(dir, ".env.dev")
	test := filepath.Join(dir, ".env.test")
	prod := filepath.Join(dir, ".env.prod")

	err := copyFile(dev, test)
	if err != nil {
		return err
	}
	err = util.SetEnv(test, "APP_MODE", "test")
	if err != nil {
		return err
	}

	err = copyFile(dev, prod)
	if err != nil {
		return err
	}
	err = util.SetEnv(prod, "APP_MODE", "production")
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(sourceFile *os.File) {
		_ = sourceFile.Close()
	}(sourceFile)

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(destinationFile *os.File) {
		_ = destinationFile.Close()
	}(destinationFile)

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync()
}

func initKey(project string) error {
	return command("gower", []string{"init", "key"}, project, "初始化 APP 密钥...")
}

func jwtKey(project string) error {
	return command("gower", []string{"jwt", "key"}, project, "初始化 JWT 密钥...")
}

func goModTidy(project string) error {
	return command("go", []string{"mod", "tidy"}, project, "下载 Go 依赖包...")
}

func pnpmInstall(project string) error {
	return command("pnpm", []string{"install"}, project, "下载前端依赖包...")
}

func initGit(project string) error {
	return command("git", []string{"init", "."}, project, "初始化 Git 仓库...")
}

func addAll(project string) error {
	return command("git", []string{"add", "."}, project, "添加所有文件...")
}

func commitM(project string) error {
	return command("git", []string{"commit", "-m", "init commit"}, project, "初始化 commit...")
}

func buildDev(project string) error {
	return command("npm", []string{"run", "dev"}, project, "构建前端库文件...")
}

func execTest(project string) error {
	return command("go", []string{"test", "-bench=Benchmark", "-tags", "tmpl,static"}, project, "执行基准测试...")
}

func overrideGower(project string) error {
	return command("go", []string{"install", "-tags", "cli"}, project, "本地化命令行工具...")
}

func command(c string, args []string, project string, hint string) error {
	fmt.Printf("---------------- %s: \n", hint)

	cmd := exec.Command(c, args...)
	cmd.Dir = filepath.Join(".", project)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scannerErr := bufio.NewScanner(stderr)
	go func() {
		for scannerErr.Scan() {
			fmt.Println(scannerErr.Text())
		}
	}()

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
