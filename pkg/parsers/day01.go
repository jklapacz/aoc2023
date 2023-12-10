package parsers

import (
	"strings"

	"github.com/charmbracelet/log"
)

// map of digit names to digits
var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type ParseMode string

func findFirstDigit(input string) (string, int) {
	var digitIndex = map[string]int{
		"one":   strings.Index(input, "one"),
		"two":   strings.Index(input, "two"),
		"three": strings.Index(input, "three"),
		"four":  strings.Index(input, "four"),
		"five":  strings.Index(input, "five"),
		"six":   strings.Index(input, "six"),
		"seven": strings.Index(input, "seven"),
		"eight": strings.Index(input, "eight"),
		"nine":  strings.Index(input, "nine"),
	}
	smallestIndex := -1
	firstDigit := ""
	for key, value := range digitIndex {
		if smallestIndex == -1 && value != -1 {
			smallestIndex = value
			firstDigit = key
		}
		if value < smallestIndex && value != -1 {
			smallestIndex = value
			firstDigit = key
		}
	}
	return firstDigit, smallestIndex
}

func findLastDigit(input string) (string, int) {
	var digitIndex = map[string]int{
		"one":   strings.LastIndex(input, "one"),
		"two":   strings.LastIndex(input, "two"),
		"three": strings.LastIndex(input, "three"),
		"four":  strings.LastIndex(input, "four"),
		"five":  strings.LastIndex(input, "five"),
		"six":   strings.LastIndex(input, "six"),
		"seven": strings.LastIndex(input, "seven"),
		"eight": strings.LastIndex(input, "eight"),
		"nine":  strings.LastIndex(input, "nine"),
	}
	idx := -1
	digit := ""
	for key, value := range digitIndex {
		if idx == -1 && value != -1 {
			idx = value
			digit = key
		}
		if value >= idx && value != -1 {
			idx = value
			digit = key
		}
	}
	return digit, idx
}

func hasMoreDigits(input string) bool {
	for key := range digitMap {
		if strings.Contains(input, key) {
			return true
		}
	}
	return false
}

func ParseDay01(input string, mode string) string {
	if mode != "lr" && mode != "rl" {
		log.Fatal("mode must be lr or rl")
	}
	for hasMoreDigits(input) {
		digitToReplace := ""
		idx := -1
		if mode == "lr" {
			digitToReplace, idx = findFirstDigit(input)
		} else {
			digitToReplace, idx = findLastDigit(input)
		}
		input = input[:idx] + digitMap[digitToReplace] + input[idx+len(digitToReplace):]
	}
	return input
}
