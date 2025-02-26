package main

import "testing"

var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_tab_findOccurences(t *testing.T) {
	tests := []struct {
		name string
		tr   string
		want int
	}{
		{"test1", input, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := newTab(tt.tr)
			if got := ta.findOccurences(); got != tt.want {
				t.Errorf("tab.findOccurences() = %v, want %v", got, tt.want)
			}
		})
	}
}
