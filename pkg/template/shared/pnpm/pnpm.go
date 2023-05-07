package pnpm

import (
	"embed"

	"github.com/OlegGulevskyy/create-zaf-app/internal/options"
	"github.com/leaanthony/gosod"
)

//go:embed all:**.yaml
var PnpmFiles embed.FS

func CreateWorkspaceRcFile(options *options.Project) {
	g := gosod.New(PnpmFiles)
	g.Extract(options.TargetDir(), options)
}
