package main

import "testing"

func Test_sumPartNumbers(t *testing.T) {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	want := 4361 // Expected sum of part numbers

	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "Test case from provided schematic",
			args: schematic,
			want: want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPartNumbers(tt.args); got != tt.want {
				t.Errorf("sumPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
