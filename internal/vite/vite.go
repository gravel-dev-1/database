//go:build !dev

package vite

import (
	"embed"
	"io/fs"
)

var FS fs.FS

//go:embed build/* build/**/*
var productionFS embed.FS

func init() { FS, _ = fs.Sub(productionFS, "build") }
