package consoles

import "embed"

//go:embed make/*
var tplFS embed.FS

//go:embed create/gower.zip
var gower embed.FS
