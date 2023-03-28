//go:build test

package envs

import "github.com/joho/godotenv"

func init() {
	if err := godotenv.Overload(".env.test"); err != nil {
		if err = loadFile(".env.test", true); err != nil {
			panic("环境加载失败")
		}
	}
}
