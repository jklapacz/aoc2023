package day03

import (
	"testing"
)

func TestBuildNeighborhood(t *testing.T) {
	input := "..123..456.."

	neighborhood := Build(input, 1)
	t.Run("Build neighborhood", func(t *testing.T) {
		t.Logf("neighborhood: %v\n", neighborhood)
		member := neighborhood[123]
		if member.PartNumber != 122 {
			t.Errorf("expected %d, got %d\n", 123, member.PartNumber)
		}
		if member.Row != 1 {
			t.Errorf("expected %d, got %d\n", 1, member.Row)
		}
		if member.Start != 2 {
			t.Errorf("expected %d, got %d\n", 2, member.Start)
		}
	})
}

func TestParseNumberSequences(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			input: "1",
			want:  []int{1},
		},
		{
			input: "1...2",
			want:  []int{1, 2},
		},
		{
			input: ".664.598..",
			want:  []int{664, 598},
		},
		{
			input: "...*..................................%......#.....*....290.........................492.............656...@953.....................+830.....",
			want:  []int{290, 492, 656, 953, 830},
		},
	}
	for _, tt := range tests {
		t.Run("Parsing number sequences", func(t *testing.T) {
			t.Logf("---> %s\n", tt.input[0:1])
			output := ParseNumberSequences(tt.input)

			if len(output) != len(tt.want) {
				t.Errorf("expected %d, got %d\n", len(tt.want), len(output))
			}

			for i := range output {
				if output[i] != tt.want[i] {
					t.Errorf("expected %d, got %d\n", tt.want[i], output[i])
				}
			}
		})
	}
}
