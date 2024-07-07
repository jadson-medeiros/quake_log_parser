package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/jadson-medeiros/quake_log_parser/internal/input"
	"github.com/jadson-medeiros/quake_log_parser/internal/output"
	"github.com/jadson-medeiros/quake_log_parser/internal/parser"
)

func main() {
	start := time.Now()

	// Get the absolute path of the log file
	path, err := input.GetAbsFilePath()
	if err != nil {
		fmt.Println("Error getting file path:", err)
		return
	}

	fmt.Println("Reading log file at", path, "...")

	// Read the content of the log file
	content, err := input.GetLogFileContent(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Parsing log data...")

	// Parse the log data using the parser
	matches := parser.Parse(content)

	fmt.Println("Writing report to JSON file...")

	// Write the report to a JSON file
	filename := filepath.Base(path)
	err = output.WriteReportToFile(matches, filename)
	if err != nil {
		log.Fatalf("Error writing report to file: %v", err)
	}

	// Calculate and display the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Parsing Quake log completed in %s\n", elapsed)
}
