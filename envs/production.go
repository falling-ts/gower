//go:build prod

package envs

import "github.com/joho/godotenv"

func init() {
	if err := godotenv.Overload("envs/.env.production"); err != nil {
		if err = loadFile(".env.production", true); err != nil {
			if err := godotenv.Overload(".env.production"); err != nil {
				panic("环境加载失败")
			}
		}
	}
}
