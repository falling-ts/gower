//go:build static

package public

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
//go:embed static/images/*
var static embed.FS

func init() {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	Static = http.FS(sub)
}
