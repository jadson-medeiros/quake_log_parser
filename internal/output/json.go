package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// WriteJsonToFile serializes the data to JSON and writes it to a file.
func WriteJsonToFile(data interface{}, fileName string) error {
	// Marshal the data into JSON format with indentation
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %v", err)
	}

	// Ensure the reports directory exists
	err = os.MkdirAll("./reports", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// Create the full file path
	filePath := filepath.Join("./reports", fileName+".json")

	// Write the JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	return nil
}
