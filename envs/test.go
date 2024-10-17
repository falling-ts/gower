//go:build test

package envs

import (
	"embed"
	"github.com/joho/godotenv"
)

//go:embed .env.test
var test embed.FS

func init() {
	Envs = &test

	if err := godotenv.Overload("envs/.env.test"); err != nil {
		if err = godotenv.Overload(".env.test"); err != nil {
			if err = loadFile(".env.test", true); err != nil {
				panic("环境加载失败")
			}
		}
	}
}
