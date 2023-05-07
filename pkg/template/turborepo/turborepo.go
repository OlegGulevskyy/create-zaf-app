package turborepo

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/OlegGulevskyy/create-zaf-app/internal/options"
	"github.com/OlegGulevskyy/create-zaf-app/pkg/env"
	fsutils "github.com/OlegGulevskyy/create-zaf-app/pkg/fs-utils"
	"github.com/OlegGulevskyy/create-zaf-app/pkg/template/shared/npm"
	"github.com/OlegGulevskyy/create-zaf-app/pkg/template/shared/pnpm"

	"github.com/leaanthony/gosod"
)

//go:embed all:packages/**
//go:embed all:**.js
//go:embed all:.gitignore
//go:embed all:**.md
//go:embed all:**.json
var TurborepoStaticFiles embed.FS

func Create(options *options.Project) {
	fmt.Printf("[turborepo.go]: %+v", options)
	g := gosod.New(TurborepoStaticFiles)
	g.Extract(options.TargetDir(), options)

	fsutils.CreateFolderIfNotExists(options.AppsDir())

	executeViteCli(options)
	setWorkspacesRc(options)
}

func setWorkspacesRc(opts *options.Project) {
	if opts.PackageManager == "pnpm" {
		pnpm.CreateWorkspaceRcFile(opts)
	} else if opts.PackageManager == "npm" {
		packageJsonPath := path.Join(opts.TargetDir(), "package.json")
		npm.AddWorkspacesToPackageJson(packageJsonPath)
	}
}

func executeViteCli(opts *options.Project) {
	// if pkg manager is npm, we postfix Cli with @latest
	viteCliEntryPoint := "vite"
	if opts.PackageManager == "npm" {
		viteCliEntryPoint = "vite@latest"
	}

	viteCliargs := []string{
		opts.PackageManager,
		"create",
		viteCliEntryPoint,
		"addon",
	}

	if opts.PackageManager == "npm" {
		npmV := env.PkgManagerVersion(opts.PackageManager)
		if npmV.Major() >= 7 {
			viteCliargs = append(viteCliargs, "--")

		}
	}
	viteCliargs = append(viteCliargs, "--template", opts.Framework)

	cmd := exec.Command(viteCliargs[0], viteCliargs[1:]...)
	cmd.Dir = opts.AppsDir()

	// Attach standard output and standard error for logging if debug is enabled
	if opts.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// Execute the command
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Failed to create the Vite app: %s", err)
	}
}
