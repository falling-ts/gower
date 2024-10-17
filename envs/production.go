//go:build prod

package envs

import (
	"embed"
	"github.com/joho/godotenv"
)

//go:embed .env.production
var prod embed.FS

func init() {
	Envs = &prod

	if err := godotenv.Overload("envs/.env.production"); err != nil {
		if err = godotenv.Overload(".env.production"); err != nil {
			if err = loadFile(".env.production", true); err != nil {
				panic("环境加载失败")
			}
		}
	}
}
