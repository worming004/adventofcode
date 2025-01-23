package main

import (
	"testing"
)

var testRules = rules{
	rule{97, 13},
	rule{97, 61},
	rule{97, 47},
	rule{75, 29},
	rule{61, 13},
	rule{75, 53},
	rule{29, 13},
	rule{97, 29},
	rule{53, 29},
	rule{61, 53},
	rule{97, 53},
	rule{61, 29},
	rule{47, 13},
	rule{75, 47},
	rule{97, 75},
	rule{47, 61},
	rule{75, 61},
	rule{47, 29},
	rule{75, 13},
	rule{53, 13},
}

func Test_rules_IsValid(t *testing.T) {
	type args struct {
		values []uint
	}
	tests := []struct {
		name  string
		r     rules
		args  args
		want  bool
		value int
	}{
		{"test1", testRules, args{[]uint{75, 47, 61, 53, 29}}, true, 61},
		{"test1", testRules, args{[]uint{97, 61, 53, 29, 13}}, true, 53},
		{"test1", testRules, args{[]uint{75, 29, 13}}, true, 29},
		{"test1", testRules, args{[]uint{75, 97, 47, 61, 53}}, false, 0},
		{"test1", testRules, args{[]uint{61, 13, 29}}, false, 0},
		{"test1", testRules, args{[]uint{97, 13, 75, 29, 47}}, false, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, _ := tt.r.IsValid(tt.args.values)
			if got := isValid; got != tt.want {
				t.Errorf("rules.IsValid() = %v, want %v", got, tt.want)
				return
			}
			if isValid {
				middle := getMiddle(tt.args.values)
				if middle != uint(tt.value) {
					t.Errorf("getMiddle() = %v, want %v", middle, tt.value)
				}
			}
		})
	}
}

func TestParse(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	ParsedRules, ParsedValues := Parse(input)
	if ParsedRules[0] != (rule{47, 53}) {
		t.Errorf("ParsedRules[0] = %v, want %v", ParsedRules[0], pair{47, 53})
	}
	if ParsedRules[2] != (rule{97, 61}) {
		t.Errorf("ParsedRules[2] = %v, want %v", ParsedRules[2], pair{97, 61})
	}

	if !sliceEquals(ParsedValues[0], []uint{75, 47, 61, 53, 29}) {
		t.Errorf("ParsedValues[0] = %v, want %v", ParsedValues[0], []uint{75, 47, 61, 53, 29})
	}
	if !sliceEquals(ParsedValues[2], []uint{75, 29, 13}) {
		t.Errorf("ParsedValues[2] = %v, want %v", ParsedValues[2], []uint{75, 29, 13})
	}
}

func sliceEquals(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMakeValid(t *testing.T) {
	tests := []struct {
		name     string
		r        rules
		input    []uint
		expected []uint
	}{
		{"test1", testRules, []uint{75, 97, 47, 61, 53}, []uint{97, 75, 47, 61, 53}},
		{"test2", testRules, []uint{61, 13, 29}, []uint{61, 29, 13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.r.MakeValid(tt.input)
			if !sliceEquals(res, tt.expected) {
				t.Errorf("makeValid() = %v, want %v", res, tt.expected)
			}
		})
	}
}
