package day02

import (
	"aoc/pkg/parsers"
	"aoc/puzzle_utils"
)

func Part01(input string, config Round) int {
	lines := puzzle_utils.Lines(input)
	sum := 0
	for _, line := range lines {
		game := parsers.ParseDay02(line)
		if game.IsPossibleGame(config) {
			sum += game.Id
		}
	}
	return sum
}

func Part02(input string) int {
	lines := puzzle_utils.Lines(input)
	sum := 0
	for _, line := range lines {
		game := parsers.ParseDay02(line)
		sum += game.GetPower()
	}
	return sum
}
