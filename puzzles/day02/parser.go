package day02

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Round struct {
	RedCubes   int
	GreenCubes int
	BlueCubes  int
}

type Day02 struct {
	Id     int
	Rounds []Round
}

func (d Day02) HighWaterMark() Round {
	// returns the maximum number of cubes of each color from all rounds

	highWaterMarks := Round{0, 0, 0}
	for _, round := range d.Rounds {
		if round.RedCubes > highWaterMarks.RedCubes {
			highWaterMarks.RedCubes = round.RedCubes
		}
		if round.GreenCubes > highWaterMarks.GreenCubes {
			highWaterMarks.GreenCubes = round.GreenCubes
		}
		if round.BlueCubes > highWaterMarks.BlueCubes {
			highWaterMarks.BlueCubes = round.BlueCubes
		}
	}

	return highWaterMarks
}

func (d Day02) IsPossibleGame(config Round) bool {
	// returns true if the game is possible given the configuration
	// of cubes in the last round

	// get the high water mark
	highWaterMarks := d.HighWaterMark()

	// check if the configuration is possible
	if highWaterMarks.RedCubes > config.RedCubes {
		return false
	}
	if highWaterMarks.GreenCubes > config.GreenCubes {
		return false
	}
	if highWaterMarks.BlueCubes > config.BlueCubes {
		return false
	}

	return true
}

func (d Day02) GetPower() int {
	hwm := d.HighWaterMark()
	return hwm.RedCubes * hwm.GreenCubes * hwm.BlueCubes
}

func ParseDay02(input string) Day02 {
	output := Day02{}
	output.Id = getGame(input)
	output.Rounds = getRounds(input)

	return output
}

func getGame(input string) int {
	// check if the pattern 'Game <number>:' exists
	pattern := regexp.MustCompile(`Game \d+:`)
	prefix := pattern.FindString(input)

	id := strings.Replace(prefix, "Game ", "", 1)
	id = strings.Replace(id, ":", "", 1)
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Warn("Could not parse game ID", "line", input)
		return 0
	}
	return intId
}

func getRounds(input string) []Round {
	// first cleanup string by removing prefix
	pattern := regexp.MustCompile(`Game \d+: `)
	prefix := pattern.FindString(input)
	roundsString := strings.Replace(input, prefix, "", 1)

	roundList := []Round{}
	// split string into rounds based on semicolon
	rounds := strings.Split(roundsString, ";")
	for _, r := range rounds {
		round := Round{}

		round.RedCubes = getRedCubes(r)
		round.GreenCubes = getGreenCubes(r)
		round.BlueCubes = getBlueCubes(r)

		roundList = append(roundList, round)
	}
	return roundList
}

func getRedCubes(input string) int {
	return parseCubes(input, "red")
}
func getGreenCubes(input string) int {
	return parseCubes(input, "green")
}
func getBlueCubes(input string) int {
	return parseCubes(input, "blue")
}

func parseCubes(input, color string) int {
	// use a named capture group to get the number of "color" cubes
	pattern := regexp.MustCompile(`(?P<color>\d+) ` + color)
	matches := pattern.FindStringSubmatch(input)

	if len(matches) < 1 {
		return 0
	}

	redCount := matches[1]
	cubes, err := strconv.Atoi(redCount)
	if err != nil {
		log.Warn("Could not parse %s cubes", color)
		return 0
	}
	return cubes
}
