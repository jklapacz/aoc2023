package puzzle_utils

import "strings"

func Lines(input string) []string {
	// input is newline separated string
	return strings.Split(input, "\n")
}
