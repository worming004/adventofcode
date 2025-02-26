package utils

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func SplitLines(i string) []string {
	lines := strings.Split(i, "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}

	return lines
}
