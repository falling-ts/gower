//go:build prod

package envs

import (
	"embed"
	"github.com/joho/godotenv"
)

//go:embed .env.prod
var prod embed.FS

func init() {
	Envs = &prod

	if err := godotenv.Overload("envs/.env.prod"); err != nil {
		if err = godotenv.Overload(".env.prod"); err != nil {
			if err = loadFile(".env.prod", true); err != nil {
				panic("环境加载失败")
			}
		}
	}
}
