package npm

import (
	"fmt"

	fsutils "github.com/OlegGulevskyy/create-zaf-app/pkg/fs-utils"
)

// Update existing manifest JSON to include workspaces property
func AddWorkspacesToPackageJson(path string) error {
	jsonData, err := fsutils.GetJsonFromFile(path)
	if err != nil {
		return fmt.Errorf("error getting JSON from file: %w", err)
	}

	// Add the key-value pair to the map
	jsonData["workspaces"] = []string{"apps/*", "packages/*"}

	fsutils.WriteJsonToFile(jsonData, path)

	return nil

}
