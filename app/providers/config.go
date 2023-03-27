package providers

import (
	"io/fs"
	"os"
	"strings"

	"github.com/falling-ts/gower/configs"
	"github.com/falling-ts/gower/envs"
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/config"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

var _ services.Config = (*config.Service)(nil)

func init() {
	P.Register("config", func(...services.Service) services.Service {
		if err := godotenv.Load(".env"); err != nil {
			if err = godotenv.Load(".env.example"); err != nil {
				if err := godotenv.Load("envs/.env"); err != nil {
					if envs.FS == nil {
						if err = godotenv.Load("envs/.env.example"); err != nil {
							panic("加载环境失败")
						}
					} else {
						if err = loadFile(".env", true); err != nil {
							if err = loadFile(".env.example", true); err != nil {
								panic("加载环境失败")
							}
						}
					}
				}
			}
		}

		c := new(configs.Config)
		if err := env.Parse(c); err != nil {
			panic(err)
		}

		return config.Mount(c).Init()
	})
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := envs.FS.Open(filename)
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
