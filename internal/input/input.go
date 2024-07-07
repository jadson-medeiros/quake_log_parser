package input

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// GetAbsFilePath returns the absolute path of the log file.
// If a command line argument is provided, it will be used as the path.
// Otherwise, a default path will be used.
func GetAbsFilePath() (string, error) {
	var path string

	if len(os.Args) > 1 {
		path = strings.Join(os.Args[1:], " ")
	} else {
		path = "assets/logs/qgames.log"
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

// GetLogFileContent returns the content of the log file as a string.
func GetLogFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
