//go:build static

package public

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
//go:embed static/images/*
var public embed.FS

func init() {
	sub, err := fs.Sub(public, "static")
	if err != nil {
		panic(err)
	}
	FS = http.FS(sub)
}
