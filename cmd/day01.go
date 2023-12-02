package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var exampleInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func turnInputIntoSlice(input string) []string {
	// input is newline separated string
	return strings.Split(input, "\n")
}

func hasSpelledOutNumber(line string) bool {
	// returns true if the line contains a spelled out number
	return strings.Contains(line, "one") ||
		strings.Contains(line, "two") ||
		strings.Contains(line, "three") ||
		strings.Contains(line, "four") ||
		strings.Contains(line, "five") ||
		strings.Contains(line, "six") ||
		strings.Contains(line, "seven") ||
		strings.Contains(line, "eight") ||
		strings.Contains(line, "nine")
}

func replaceSpelledOutNumber(line string) string {
	// replaces first instance of a spelled out number with its value
	// strings.Replace()
	return line
}

func preprocessLine(line string) string {
	return line
}

func getFirstDigit(line string) int {
	location := strings.IndexFunc(line, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	if location == -1 {
		log.Error("No digit found", "line", line)
		return -1
	}
	digit := string(line[location])
	res, err := strconv.Atoi(digit)
	if err != nil {
		return -1
	}
	return res
}

func getLastDigit(line string) int {
	location := strings.LastIndexFunc(line, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	if location == -1 {
		log.Error("No digit found", "line", line)
		return -1
	}
	digit := string(line[location])
	res, err := strconv.Atoi(digit)
	if err != nil {
		return -1
	}
	return res
}

func getCalibrationNumber(line string) int {
	first := getFirstDigit(line)
	last := getLastDigit(line)

	combinedNumber := fmt.Sprint(first) + fmt.Sprint(last)

	res, err := strconv.Atoi(combinedNumber)
	if err != nil {
		return -1
	}
	log.Info(res)
	return res
}

func readFile(filename string) string {
	// opens a file and returns the contents as a string
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Error("Error reading file")
		return ""
	}
	return string(contents)
}

var inputFile string

var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Solve Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if inputFile == "" {
			log.Error("No input file specified")
			input = exampleInput
		} else {
			input = readFile(inputFile)
		}
		lines := turnInputIntoSlice(input)
		sum := 0
		for _, line := range lines {
			sum += getCalibrationNumber(line)
		}
		fmt.Println(sum)
	},
}

func init() {
	day01Cmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input for the problem")
}
