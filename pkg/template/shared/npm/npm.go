package npm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Update existing manifest JSON to include workspaces property
func UpdateManifestJson(path string) error {
	// Read the JSON file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Parse the JSON content into a map
	var jsonData map[string]interface{}
	err = json.Unmarshal(content, &jsonData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// Add the key-value pair to the map
	jsonData["workspaces"] = []string{"apps/*", "packages/*"}

	// Convert the map back to JSON
	modifiedContent, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting JSON to bytes: %w", err)
	}

	// Write the modified JSON back to the file
	err = ioutil.WriteFile(path, modifiedContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil

}
