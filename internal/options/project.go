package options

import (
	"os"
	"path/filepath"
)

type Project struct {
	Name                   string
	Framework              string
	ZendeskLocation        string
	Tailwind               bool
	Debug                  bool
	PackageManager         string
	PackageManagerVersion  string
	WorkspacePackageSyntax string

	selectedListItem  string
	selectedInputItem string
}

func (p *Project) TargetDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(currentDir, p.Name)
	return path
}

func (p *Project) GetPkgManager() string {
	return p.PackageManager
}

func (p *Project) AppsDir() string {
	return filepath.Join(p.TargetDir(), "apps")
}
