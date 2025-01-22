package main

import (
	"aventofcode2024/parser"
	"strconv"
	"strings"
)

func main() {
	reports := parse()
	var safeReports int
	for _, r := range reports {
		if r.isSafe() {
			safeReports++
		}
	}
	println(safeReports)
}

type level int

type report []level

func (r report) isSafe() bool {
	if !(r.isAlwaysIncreasing() || r.isAlwaysDecreasing()) {
		return false
	}

	if !r.isDifferSafe() {
		return false
	}

	return true
}

func (r report) isAlwaysIncreasing() bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] > r[i+1] {
			return false
		}
	}
	return true
}
func (r report) isAlwaysDecreasing() bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] < r[i+1] {
			return false
		}
	}
	return true
}

func (r report) isDifferSafe() bool {
	for i := 0; i < len(r)-1; i++ {
		diff := r[i+1] - r[i]
		if diff < -3 || diff > 3 {
			return false
		}

		if diff == 0 {
			return false
		}
	}
	return true
}

func parse() []report {
	lines := parser.Parse("input.txt")

	var result []report
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		splittedLvl := sliceAtoLvl(splitted)

		result = append(result, splittedLvl)
	}

	return result

}

func sliceAtoLvl(s []string) []level {
	var result []level = make([]level, 0, len(s))
	for _, s := range s {
		i, _ := strconv.Atoi(s)
		result = append(result, level(i))
	}
	return result
}
