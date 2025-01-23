package main

import (
	"strings"
	"testing"
)

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

func TestInput(t *testing.T) {
	if findOccurrence(input) != 18 {
		t.Error("Expected 18, got ", findOccurrence(input))
	}
}

func TestDiag1(t *testing.T) {
	input := `abc
def
ghi`

	lines := strings.Split(input, "\n")
	dg := diag1(lines)
	if len(dg) != 5 {
		t.Error("Expected 5, got ", len(dg))
	}

	if dg[0] != "a" {
		t.Error("Expected a, got ", dg[0])
	}
	if dg[1] != "db" {
		t.Error("Expected db, got ", dg[0])
	}
	if dg[2] != "gec" {
		t.Error("Expected gec, got ", dg[0])
	}
	if dg[3] != "hf" {
		t.Error("Expected hf, got ", dg[0])
	}
	if dg[4] != "i" {
		t.Error("Expected i, got ", dg[0])
	}
}

func TestDiag2(t *testing.T) {
	input := `abc
def
ghi`

	lines := strings.Split(input, "\n")
	dg := diag2(lines)
	if len(dg) != 5 {
		t.Error("Expected 5, got ", len(dg))
	}

	if dg[0] != "aei" {
		t.Error("Expected aei, got ", dg[0])
	}
	if dg[1] != "dh" {
		t.Error("Expected dh, got ", dg[0])
	}
	if dg[2] != "g" {
		t.Error("Expected g, got ", dg[0])
	}
	if dg[3] != "bf" {
		t.Error("Expected be, got ", dg[0])
	}
	if dg[4] != "c" {
		t.Error("Expected c, got ", dg[0])
	}
}

//
// . . 2
// . 1 .
// 0 . .
//
// 0: i=2;j=0
// 1: i=1;j=1
// 2: i=0;j=2
//
//
// 0 . .
// . 1 .
// . . 2
//
// 0: i=0;j=0
// 1: i=1;j=1
// 2: i=2;j=2
//
