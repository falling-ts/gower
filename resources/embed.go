//go:build tmpl

package resources

import "embed"

//go:embed views/*
//go:embed views/**/*
//go:embed views/**/**/*
var tmpl embed.FS

func init() {
	Tmpl = &tmpl
}
