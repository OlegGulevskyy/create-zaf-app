package fsutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) == true {
		// Create the targetDirectory
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetJsonFromFile(path string) (map[string]interface{}, error) {
	// Read the JSON file content
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse the JSON content into a map
	var jsonData map[string]interface{}
	err = json.Unmarshal(content, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return jsonData, nil
}

func WriteJsonToFile(jsonData map[string]interface{}, path string) error {

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
