package turborepo

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/OlegGulevskyy/create-zaf-app/internal/options"
	fsutils "github.com/OlegGulevskyy/create-zaf-app/pkg/fs-utils"
	"github.com/leaanthony/gosod"
)

//go:embed all:packages/**
//go:embed all:**.js
//go:embed all:.gitignore
//go:embed all:**.md
//go:embed all:**.json
var TurborepoStaticFiles embed.FS

func Create(options *options.Project) {
	fmt.Println("Creating turborepo project at", options)
	g := gosod.New(TurborepoStaticFiles)
	g.Extract(options.TargetDir(), options)

	fmt.Println("Creating apps folder at", options.AppsDir())
	fsutils.CreateFolderIfNotExists(options.AppsDir())

	executeViteCli(options)
}

func executeViteCli(opts *options.Project) {
	// if pkg manager is npm, we postfix Cli with @latest
	viteCliEntryPoint := "vite"
	if opts.PackageManager == "npm" {
		viteCliEntryPoint = "vite@latest"
	}
	cmd := exec.Command(opts.PackageManager, "create", viteCliEntryPoint, "addon", "--template", opts.Framework)
	cmd.Dir = opts.AppsDir()

	// Attach standard output and standard error for logging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Failed to create the Vite app: %s", err)
	}
}
