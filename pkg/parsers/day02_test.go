package parsers

import (
	"fmt"
	"testing"
)

func TestParsingGame(t *testing.T) {
	tests := []struct {
		input string
		want  Day02
	}{
		{
			input: "Game 1: 1 red",
			want: Day02{
				Id: 1,
				Rounds: []Round{
					{
						RedCubes:   1,
						GreenCubes: 0,
						BlueCubes:  0,
					},
				},
			},
		},
		{
			input: "Game 1: 2 green; 5 red; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: Day02{
				Id: 1,
				Rounds: []Round{
					{
						RedCubes:   0,
						GreenCubes: 2,
						BlueCubes:  0,
					},
					{
						RedCubes:   5,
						GreenCubes: 0,
						BlueCubes:  0,
					},
					{
						RedCubes:   1,
						GreenCubes: 3,
						BlueCubes:  4,
					},
					{
						RedCubes:   0,
						GreenCubes: 1,
						BlueCubes:  1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("Parsing round", func(t *testing.T) {
			output := ParseDay02(tt.input)

			fmt.Println(output)

			if output.Id != tt.want.Id {
				t.Errorf("ID mismatch, id=%d, want=%d", output.Id, tt.want.Id)
			}
			if len(output.Rounds) != len(tt.want.Rounds) {
				t.Errorf("Round count mismatch, count=%d, want=%d", len(output.Rounds), len(tt.want.Rounds))
			}
			for i, round := range output.Rounds {
				if round.RedCubes != tt.want.Rounds[i].RedCubes {
					t.Errorf("Red cube count mismatch, count=%d, want=%d", round.RedCubes, tt.want.Rounds[i].RedCubes)
				}
				if round.GreenCubes != tt.want.Rounds[i].GreenCubes {
					t.Errorf("Green cube count mismatch, count=%d, want=%d", round.GreenCubes, tt.want.Rounds[i].GreenCubes)
				}
				if round.BlueCubes != tt.want.Rounds[i].BlueCubes {
					t.Errorf("Blue cube count mismatch, count=%d, want=%d", round.BlueCubes, tt.want.Rounds[i].BlueCubes)
				}
			}
		})
	}
}

func TestParseDay02Basic(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "ID parsing",
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  2,
		},
		{
			name:  "ID parsing",
			input: "Game 69: 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  69,
		},
		{
			name:  "ID parsing",
			input: "Game 1: 2 green; 5 red; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ParseDay02(tt.input)

			fmt.Println(output)

			if output.Id != tt.want {
				t.Errorf("Parse() = %v, want %v", output, tt.want)
			}
		})
	}
}

func TestIfConfigurationIsPossible(t *testing.T) {
	tests := []struct {
		input  string
		config Round
		want   bool
	}{
		{
			input:  "Game 1: 1 red",
			config: Round{0, 0, 0},
			want:   false,
		},
		{
			input:  "Game 1: 1 red",
			config: Round{4, 0, 0},
			want:   true,
		},
		{
			input:  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			config: Round{12, 13, 14},
			want:   true,
		},
		{
			input:  "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			config: Round{12, 13, 14},
			want:   true,
		},
		{
			input:  "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			config: Round{12, 13, 14},
			want:   false,
		},
		{
			input:  "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			config: Round{12, 13, 14},
			want:   false,
		},
		{
			input:  "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			config: Round{12, 13, 14},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run("Checking if possible", func(t *testing.T) {
			output := ParseDay02(tt.input)
			if output.IsPossibleGame(tt.config) != tt.want {
				t.Logf("--> high water mark: %v\n", output.HighWaterMark())
				t.Errorf("IsPossibleGame() = %v, want %v", output.IsPossibleGame(tt.config), tt.want)
			}
		})
	}
}
