package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parsed := parse(string(content))
	result := 0
	for _, l := range parsed {
		tester := newSecondOperationTester(l.result, l.values)
		if tester.test() {
			result += tester.result
		}
	}
	fmt.Println(result)
}

type line struct {
	result int
	values []int
}

func parse(input string) []line {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	result := make([]line, len(lines))
	for i, l := range lines {
		v := strings.Split(l, ":")[0]
		othersStr := strings.Split(strings.Split(l, ":")[1], " ")
		if othersStr[0] == "" {
			othersStr = othersStr[1:]
		}
		others := make([]int, len(othersStr))
		for i, v := range othersStr {
			others[i] = atoi(v)
		}
		result[i] = line{result: atoi(v), values: others}
	}

	return result
}

type operator func(int, int) int

func add(a, b int) int {
	return a + b
}
func mult(a, b int) int {
	return a * b
}
func concat(a, b int) int {
	bCopy := b
	digitCounter := 0
	for bCopy > 0 {
		bCopy = bCopy / 10
		digitCounter++
	}
	multiplier := 1
	for i := 0; i < digitCounter; i++ {
		multiplier *= 10
	}
	return (a * multiplier) + b
}

type operationTester struct {
	result    int
	values    []int
	operators []operator
}

func newDefaultOperationTester(expectedResult int, values []int) *operationTester {
	return &operationTester{
		expectedResult,
		values,
		[]operator{add, mult},
	}
}

func newSecondOperationTester(expectedResult int, values []int) *operationTester {
	return &operationTester{
		expectedResult,
		values,
		[]operator{add, mult, concat},
	}
}

func (ot *operationTester) test() bool {
	return subtest(ot.result, ot.values[0], ot.values[1:], ot.operators)
}

func subtest(expected, result int, values []int, operators []operator) bool {
	if len(values) == 0 {
		return expected == result
	}

	for _, operator := range operators {
		if subtest(expected, operator(result, values[0]), values[1:], operators) {
			return true
		}
	}
	return false
}

func atoi(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return r
}
