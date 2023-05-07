package options

import (
	"os"
	"path/filepath"
)

type Project struct {
	Name                     string
	Framework                string
	ZendeskLocation          string
	ZendeskLocationSanitized string
	ProductionUrl            string
	DevUrl                   string
	Tailwind                 bool
	Debug                    bool
	PackageManager           string
	PackageManagerVersion    string
	WorkspacePackageSyntax   string

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

var projZendeskLocationsMapping = map[string]string{
	"Ticket sidebar":       "ticket_sidebar",
	"New ticket sidebar":   "new_ticket_sidebar",
	"Organization sidebar": "organization_sidebar",
	"User sidebar":         "user_sidebar",
	"Top bar":              "top_bar",
	"Nav bar":              "nav_bar",
	"Modal":                "modal",
	"Ticket editor":        "ticket_editor",
	"Background":           "background",
}

func (p *Project) GetZendeskLocation() string {
	return projZendeskLocationsMapping[p.ZendeskLocation]
}

func (p *Project) GetProductionUrl() string {
	return "https://<<PLACE YOUR PRODUCTION URL HERE>>/"
}

func (p *Project) GetDevUrl() string {
	return "http://localhost:5173"
}
