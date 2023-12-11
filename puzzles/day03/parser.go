package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type GridBlock struct {
	TopLeft     int
	TopRight    int
	BottomLeft  int
	BottomRight int
}

type Member struct {
	PartNumber int
	Row        int
	Start      int
	End        int
}

type Neighborhood struct {
	PartNumber int
	Block      GridBlock
}

func Build(input string, row int) map[int]Member {
	// build a map of the members of the neighborhood
	// each member is identified by the part number
	// and the row number
	members := map[int]Member{}
	numbers := ParseNumberSequences(input)
	for _, n := range numbers {
		// get the start and end positions of the number
		// in the input string
		start := strings.Index(input, strconv.Itoa(n))
		end := start + len(strconv.Itoa(n)) - 1
		member := Member{
			PartNumber: n,
			Row:        row,
			Start:      start,
			End:        end,
		}
		members[n] = member
	}
	return members
}

func ParseNumberSequences(input string) []int {
	// number sequences are separated by any symbol
	// use a regex to find all occurrences of a pattern
	pattern := regexp.MustCompile(`\d+`)
	found := pattern.FindAllString(input, -1)

	foundSequences := []int{}
	for _, f := range found {
		converted, err := strconv.Atoi(f)
		if err != nil {
			log.Warn("Could not parse number sequence", "line", input)
		}
		foundSequences = append(foundSequences, converted)
	}
	return foundSequences
}
