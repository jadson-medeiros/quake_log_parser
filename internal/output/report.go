package output

import (
	"fmt"

	"github.com/jadson-medeiros/quake_log_parser/internal/pkg/data"
)

// WriteReportToFile writes the match data to a JSON file.
func WriteReportToFile(matches map[string]*data.MatchData, filename string) error {
	err := writeGroupedInformation(matches, filename)
	if err != nil {
		return fmt.Errorf("failed to write report to file: %v", err)
	}

	fmt.Println("Report written to", filename+".json")
	return nil
}

// writeGroupedInformation writes the match data to a JSON file.
func writeGroupedInformation(matches map[string]*data.MatchData, filename string) error {
	return WriteJsonToFile(matches, filename)
}
