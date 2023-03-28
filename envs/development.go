package envs

import (
	"embed"
	"io/fs"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Envs *embed.FS

func init() {
	if err := godotenv.Load("envs/.env.development"); err != nil {
		if err = loadFile(".env.development", false); err != nil {
			if err := godotenv.Load(".env.development"); err != nil {
				panic("环境加载失败")
			}
		}
	}
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := Envs.Open(filename)
	if err != nil {
		return
	}
	defer func(file fs.File) {
		_ = file.Close()
	}(file)

	return godotenv.Parse(file)
}

func loadFile(filename string, overload bool) error {
	envMap, err := readFile(filename)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			_ = os.Setenv(key, value)
		}
	}

	return nil
}
