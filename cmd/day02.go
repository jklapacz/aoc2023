package cmd

import (
	"aoc/pkg/parsers"
	"aoc/puzzle_utils"
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
		config := parsers.Round{
			RedCubes:   redCubes,
			GreenCubes: greenCubes,
			BlueCubes:  blueCubes}
		lines := turnInputIntoSlice(input)
		games := []parsers.Day02{}
		for _, line := range lines {
			game := parsers.ParseDay02(line)
			games = append(games, game)
		}
		sum := 0
		if power {
			for _, game := range games {
				sum += game.GetPower()
			}
			fmt.Println("Sum of all powers:", sum)
		} else {
			for _, game := range games {
				if game.IsPossibleGame(config) {
					sum += game.Id
					log.Infof("Game %d: %v", game.Id, game.Rounds)
				}
			}
			fmt.Println("Sum of possible games:", sum)
		}
	},
}

func init() {
	day02Cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input for the problem")
	day02Cmd.Flags().IntVarP(&redCubes, "red", "r", 0, "Number of red cubes")
	day02Cmd.Flags().IntVarP(&greenCubes, "green", "g", 0, "Number of green cubes")
	day02Cmd.Flags().IntVarP(&blueCubes, "blue", "b", 0, "Number of blue cubes")
	day02Cmd.Flags().BoolVarP(&power, "power", "p", false, "Get the power of the game")
}
