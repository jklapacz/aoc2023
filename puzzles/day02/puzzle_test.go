package day02

import (
	"aoc/puzzle_utils"
	"testing"
)

func TestPuzzleDay02(t *testing.T) {
	// read in the input file
	input := puzzle_utils.ReadFile("../../inputs/day02-part1.txt")
	t.Run("Part 1", func(t *testing.T) {
		config := Round{
			RedCubes:   12,
			GreenCubes: 13,
			BlueCubes:  14,
		}
		sum := Part01(input, config)
		answer := 2101
		if sum != answer {
			t.Errorf("expected %d, got %d\n", answer, sum)
		}
	})
	t.Run("Part 2", func(t *testing.T) {
		sum := Part02(input)
		answer := 58269
		if sum != answer {
			t.Errorf("expected %d, got %d\n", answer, sum)
		}
	})
}
