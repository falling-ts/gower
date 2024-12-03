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
	"path"
	"path/filepath"
	"regexp"
	"strings"
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
	pattern := `^[a-zA-Z][a-zA-Z0-9-_]*$`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(project) {
		return errors.New("项目名称必须字母开头，只能包含字母、数字、-、_")
	}

	dir := project
	if util.IsExist("src") {
		dir = filepath.Join("src", dir)
	}

	err := create(dir, project)
	if err != nil {
		return err
	}

	err = setGradleProp(dir, project)
	if err != nil {
		return err
	}

	err = setGoMod(dir, project)
	if err != nil {
		return err
	}

	err = setGoLandRun(dir, project)
	if err != nil {
		return err
	}

	err = setService(dir, project)
	if err != nil {
		return err
	}

	err = initEnv(dir)
	if err != nil {
		return err
	}

	err = initEnv(filepath.Join(dir, "envs"))
	if err != nil {
		return err
	}

	err = initKey(dir)
	if err != nil {
		return err
	}

	err = jwtKey(dir)
	if err != nil {
		return err
	}

	err = goModTidy(dir)
	if err != nil {
		return err
	}

	err = pnpmInstall(dir)
	if err != nil {
		return err
	}

	err = initGit(dir)
	if err != nil {
		return err
	}

	err = addAll(dir)
	if err != nil {
		return err
	}

	err = commitM(dir)
	if err != nil {
		return err
	}

	err = buildDev(dir)
	if err != nil {
		return err
	}

	if util.IsExist("go.work") {
		err = addWork(dir)
		if err != nil {
			return err
		}
	}

	if util.IsExist("settings.gradle") {
		err = addGradle(project)
		if err != nil {
			return err
		}
	}

	err = execTest(dir)
	if err != nil {
		return err
	}

	return nil
}

func create(dir string, project string) error {
	if util.IsExist(dir) {
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
			filePath := filepath.Join(dir, file.Name)

			if file.FileInfo().IsDir() {
				err = os.MkdirAll(filePath, os.ModePerm)
				if err != nil {
					return err
				}
			} else {
				if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
					return err
				}

				srcFile, err := file.Open()
				if err != nil {
					return err
				}
				defer func(srcFile io.ReadCloser) {
					_ = srcFile.Close()
				}(srcFile)

				content, err := io.ReadAll(srcFile)
				if err != nil {
					return err
				}

				newContent := strings.ReplaceAll(string(content), "\"gitee.com/falling-ts/gower", fmt.Sprintf("\"%s", project))

				destFile, err := os.Create(filePath)
				if err != nil {
					return err
				}
				defer func(destFile *os.File) {
					_ = destFile.Close()
				}(destFile)

				if _, err = destFile.WriteString(newContent); err != nil {
					return err
				}
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	gowerZip, err := os.Create(filepath.Join(dir, "app/consoles/create/gower.zip"))
	if err != nil {
		return err
	}
	defer func(gowerZip *os.File) {
		_ = gowerZip.Close()
	}(gowerZip)

	_, err = io.Copy(gowerZip, file)
	if err != nil {
		return err
	}

	fmt.Printf("%s 项目创建成功\n", dir)
	return nil
}

func setGradleProp(dir string, project string) error {
	propFile := filepath.Join(dir, "gradle.properties")

	content, err := os.ReadFile(propFile)
	if err != nil {
		return err
	}

	contentStr := string(content)
	newContent := strings.Replace(contentStr, "bin = gower", fmt.Sprintf("bin = %s", project), -1)
	return os.WriteFile(propFile, []byte(newContent), 0644)
}

func setGoMod(dir string, project string) error {
	modFile := filepath.Join(dir, "go.mod")

	content, err := os.ReadFile(modFile)
	if err != nil {
		return err
	}

	contentStr := string(content)
	newContent := strings.Replace(contentStr, "gitee.com/falling-ts/gower", project, -1)
	return os.WriteFile(modFile, []byte(newContent), 0644)
}

func setGoLandRun(dir string, project string) error {
	runFile := filepath.Join(dir, ".run", "gower.run.xml")

	content, err := os.ReadFile(runFile)
	if err != nil {
		return err
	}

	contentStr := string(content)
	newContent := strings.Replace(contentStr, "name=\"gower\"", fmt.Sprintf("name=\"%s\"", project), -1)
	err = os.WriteFile(runFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	newRunFile := filepath.Join(dir, ".run", fmt.Sprintf("%s.run.xml", project))
	err = os.Rename(runFile, newRunFile)
	if err != nil {
		return err
	}

	runFile = filepath.Join(dir, ".run", "gower-dev.run.xml")

	content, err = os.ReadFile(runFile)
	if err != nil {
		return err
	}

	contentStr = string(content)
	newContent = strings.Replace(contentStr, "name=\"gower-dev\"", fmt.Sprintf("name=\"%s-dev\"", project), -1)
	err = os.WriteFile(runFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	newRunFile = filepath.Join(dir, ".run", fmt.Sprintf("%s-dev.run.xml", project))
	err = os.Rename(runFile, newRunFile)
	if err != nil {
		return err
	}

	runFile = filepath.Join(dir, ".run", "gower-test.run.xml")

	content, err = os.ReadFile(runFile)
	if err != nil {
		return err
	}

	contentStr = string(content)
	newContent = strings.Replace(contentStr, "name=\"gower-test\"", fmt.Sprintf("name=\"%s-test\"", project), -1)
	err = os.WriteFile(runFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	newRunFile = filepath.Join(dir, ".run", fmt.Sprintf("%s-test.run.xml", project))
	err = os.Rename(runFile, newRunFile)
	if err != nil {
		return err
	}

	runFile = filepath.Join(dir, ".run", "gower-prod.run.xml")

	content, err = os.ReadFile(runFile)
	if err != nil {
		return err
	}

	contentStr = string(content)
	newContent = strings.Replace(contentStr, "name=\"gower-prod\"", fmt.Sprintf("name=\"%s-prod\"", project), -1)
	err = os.WriteFile(runFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	newRunFile = filepath.Join(dir, ".run", fmt.Sprintf("%s-prod.run.xml", project))
	return os.Rename(runFile, newRunFile)
}

func setService(dir string, project string) error {
	serviceFile := filepath.Join(dir, "gower.service")

	content, err := os.ReadFile(serviceFile)
	if err != nil {
		return err
	}

	contentStr := string(content)
	newContent := strings.Replace(contentStr, "gower", project, -1)
	err = os.WriteFile(serviceFile, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	newServiceFile := filepath.Join(dir, fmt.Sprintf("%s.service", project))
	return os.Rename(serviceFile, newServiceFile)
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

func initKey(dir string) error {
	return command("gower", []string{"init", "key"}, dir, "初始化 APP 密钥...")
}

func jwtKey(dir string) error {
	return command("gower", []string{"jwt", "key"}, dir, "初始化 JWT 密钥...")
}

func goModTidy(dir string) error {
	return command("go", []string{"mod", "tidy"}, dir, "下载 Go 依赖包...")
}

func pnpmInstall(dir string) error {
	return command("pnpm", []string{"install"}, dir, "下载前端依赖包...")
}

func initGit(dir string) error {
	return command("git", []string{"init", "--initial-branch=main"}, dir, "初始化 Git 仓库...")
}

func addAll(dir string) error {
	return command("git", []string{"add", "."}, dir, "添加所有文件...")
}

func commitM(dir string) error {
	return command("git", []string{"commit", "-m", "init commit"}, dir, "初始化 commit...")
}

func buildDev(dir string) error {
	return command("npm", []string{"run", "dev"}, dir, "构建前端库文件...")
}

func addWork(dir string) error {
	return command("go", []string{"work", "use", "./"}, dir, "添加工作目录...")
}

func addGradle(project string) error {
	dir := path.Join("src", project)

	content, err := os.ReadFile("settings.gradle")
	if err != nil {
		return err
	}

	newContent := os.Expand(`
include ':${project}'
project(':${project}').projectDir = new File('${dir}')
`, func(s string) string {
		return map[string]string{
			"project": project,
			"dir":     dir,
		}[s]
	})

	updateContent := string(content) + newContent
	return os.WriteFile("settings.gradle", []byte(updateContent), 0644)

}

func execTest(project string) error {
	return command("go", []string{"test", "-bench=Benchmark", "-tags", "tmpl,static"}, project, "执行基准测试...")
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
