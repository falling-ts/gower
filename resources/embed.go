//go:build tmpl

package resources

import "embed"

//go:embed views/*
//go:embed views/**/*
var tmplFS embed.FS

func init() {
	TmplFS = &tmplFS
}
