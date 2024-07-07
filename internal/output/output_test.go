package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteJsonToFile(t *testing.T) {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	fileName := "testfile"

	err := WriteJsonToFile(data, fileName)
	if err != nil {
		t.Fatalf("Unexpected WriteJsonToFile error: %v", err)
	}

	filePath := filepath.Join("./reports", fileName+".json")
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read the written file: %v", err)
	}

	var resultData map[string]interface{}
	err = json.Unmarshal(fileContent, &resultData)
	if err != nil {
		t.Fatalf("Failed to unmarshal the JSON content: %v", err)
	}

	if len(resultData) != len(data) {
		t.Errorf("Unexpected number of elements in the result data. Got %d, want %d", len(resultData), len(data))
	}

	for key, value := range data {
		if resultValue, ok := resultData[key]; ok {
			if resultValue != value {
				t.Errorf("Unexpected value for key '%s'. Got %v, want %v", key, resultValue, value)
			}
		} else {
			t.Errorf("Key '%s' not found in the result data", key)
		}
	}

	// Clean up the written file and directory
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to clean up the written file: %v", err)
	}

	err = os.RemoveAll("./reports")
	if err != nil {
		t.Fatalf("Failed to clean up the reports directory: %v", err)
	}
}
