package parsers

import (
	"testing"
)

func TestParseBasic(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple - one",
			input: "one",
			want:  "1",
		},
		{
			name:  "simple - two",
			input: "two",
			want:  "2",
		},
		{
			name:  "simple - three",
			input: "three",
			want:  "3",
		},
		{
			name:  "simple - four",
			input: "four",
			want:  "4",
		},
		{
			name:  "simple - five",
			input: "five",
			want:  "5",
		},
		{
			name:  "simple - six",
			input: "six",
			want:  "6",
		},
		{
			name:  "simple - seven",
			input: "seven",
			want:  "7",
		},
		{
			name:  "simple - eight",
			input: "eight",
			want:  "8",
		},
		{
			name:  "simple - nine",
			input: "nine",
			want:  "9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDay01(tt.input, "lr"); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestParseCombined(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "combined - same digits",
			input: "oneone",
			want:  "11",
		},
		{
			name:  "combined - different digits",
			input: "onetwo",
			want:  "12",
		},
		{
			name:  "combined - mixed with other characters",
			input: "5nine@three",
			want:  "59@3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDay01(tt.input, "lr"); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseComplex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "complex - overlapping digits",
			input: "eightwothree",
			want:  "8wo3",
		},
		{
			name:  "complex - overlapping digits",
			input: "xtwone3four",
			want:  "x2ne34",
		},
		{
			name:  "complex - overlapping digits",
			input: "eightwo",
			want:  "8wo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDay01(tt.input, "lr"); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRightToLeft(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "complex - overlapping digits",
			input: "eightwothree",
			want:  "eigh23",
		},
		{
			name:  "complex - overlapping digits",
			input: "xtwone3four",
			want:  "xtw134",
		},
		{
			name:  "complex - overlapping digits",
			input: "eightwo",
			want:  "eigh2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDay01(tt.input, "rl"); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
