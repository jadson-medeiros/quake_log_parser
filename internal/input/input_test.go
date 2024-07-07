package input

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile(t *testing.T, content string) string {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	return tmpfile.Name()
}

func TestGetAbsFilePath(t *testing.T) {
	expectedPath := "/path/to/file.log"
	os.Args = []string{"app", expectedPath}

	absPath, err := GetAbsFilePath()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if absPath != expectedPath {
		t.Errorf("Expected path '%s', but got '%s'", expectedPath, absPath)
	}
}

func TestGetLogFileContent(t *testing.T) {
	content := "Log file content"
	tmpFilePath := createTempFile(t, content)
	defer os.Remove(tmpFilePath)

	fileContent, err := GetLogFileContent(tmpFilePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fileContent != content {
		t.Errorf("Expected content '%s', but got '%s'", content, fileContent)
	}

	nonexistentFilePath := "nonexistent.log"
	_, err = GetLogFileContent(nonexistentFilePath)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestGetLogFileContent_UsingAssert(t *testing.T) {
	content := "Log file content"
	tmpFilePath := createTempFile(t, content)
	defer os.Remove(tmpFilePath)

	fileContent, err := GetLogFileContent(tmpFilePath)
	assert.Nil(t, err, "Unexpected error")
	assert.Equal(t, content, fileContent, "Expected content does not match actual content")

	nonexistentFilePath := "nonexistent.log"
	_, err = GetLogFileContent(nonexistentFilePath)
	assert.Error(t, err, "Expected error")
}
