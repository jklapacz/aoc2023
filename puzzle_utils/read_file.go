package puzzle_utils

import (
	"os"

	"github.com/charmbracelet/log"
)

func ReadFile(filename string) string {
	// opens a file and returns the contents as a string
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Error("Error reading file")
		return ""
	}
	return string(contents)
}
