package cmd

import (
	"aoc/puzzle_utils"
	"aoc/puzzles/day02"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var redCubes int
var greenCubes int
var blueCubes int
var power bool

var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Solve Day 2",
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if inputFile == "" {
			log.Fatal("No input file specified")
		} else {
			input = puzzle_utils.ReadFile(inputFile)
		}
		if power {
			sum := day02.Part02(input)
			fmt.Printf("Sum of all powers: %d\n", sum)
			return
		}
		config := day02.Round{
			RedCubes:   redCubes,
			GreenCubes: greenCubes,
			BlueCubes:  blueCubes,
		}
		sum := day02.Part01(input, config)
		fmt.Printf("Sum of possible games: %d\n", sum)
	},
}

func init() {
	day02Cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input for the problem")
	day02Cmd.Flags().IntVarP(&redCubes, "red", "r", 0, "Number of red cubes")
	day02Cmd.Flags().IntVarP(&greenCubes, "green", "g", 0, "Number of green cubes")
	day02Cmd.Flags().IntVarP(&blueCubes, "blue", "b", 0, "Number of blue cubes")
	day02Cmd.Flags().BoolVarP(&power, "power", "p", false, "Get the power of the game")
}
