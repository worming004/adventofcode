package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	content = content[:len(content)-1]

	result := findOccurrence(string(content))
	fmt.Printf("Result: %d", result)

}

func findOccurrence(input string) int {
	tab := toTab(input)
	var result int
	for _, line := range tab {
		result += strings.Count(line, "XMAS")
	}
	return result
}

func toTab(input string) []string {
	lines := strings.Split(input, "\n")

	length := len(lines[0])

	var result []string
	//horizontal
	for _, line := range lines {
		result = append(result, line)
	}
	//vertical
	for i := 0; i < length; i++ {
		var line string
		for _, l := range lines {
			line += string(l[i])
		}
		result = append(result, line)
	}

	// 1st diagonal
	{
		for _, r := range diag1(lines) {
			result = append(result, r)
		}
	}
	// 2nd diagonal
	{
		for _, r := range diag2(lines) {
			result = append(result, r)
		}
	}
	copyOfResult := make([]string, len(result))
	copy(copyOfResult, result)
	for _, res := range copyOfResult {
		var reversed string
		for _, rn := range res {
			reversed = string(rn) + reversed
		}
		result = append(result, reversed)
	}

	return result
}

type cursor struct {
	firstNext  bool
	x, y       int
	maxX, maxY int
	strategy   func(cursor) (int, int)
}

func topRightStrategy(c cursor) (int, int) {
	return c.x + 1, c.y - 1
}

func bottomRightStrategy(c cursor) (int, int) {
	return c.x + 1, c.y + 1
}

func (c *cursor) next() bool {
	if c.firstNext {
		c.firstNext = false
		return true
	}
	c.x, c.y = c.strategy(*c)
	if c.x > c.maxX || c.y > c.maxY || c.x < 0 || c.y < 0 {
		return false
	}
	return true
}

type position struct {
	x, y int
}

func diag1(lines []string) []string {
	var result []string
	length := len(lines[0])
	startingpositions := []position{}

	for i := 0; i < length; i++ {
		startingpositions = append(startingpositions, position{0, i})
	}
	for i := 1; i < length; i++ {
		startingpositions = append(startingpositions, position{i, length - 1})
	}
	for _, pos := range startingpositions {
		c := cursor{true, pos.x, pos.y, length - 1, length - 1, topRightStrategy}
		result = append(result, c.getResults(lines))
	}
	return result
}
func diag2(lines []string) []string {
	var result []string
	length := len(lines[0])
	startingpositions := []position{}

	for i := 0; i < length; i++ {
		startingpositions = append(startingpositions, position{0, i})
	}
	for i := 1; i < length; i++ {
		startingpositions = append(startingpositions, position{i, 0})
	}
	for _, pos := range startingpositions {
		c := cursor{true, pos.x, pos.y, length - 1, length - 1, bottomRightStrategy}
		result = append(result, c.getResults(lines))
	}
	return result
}

func (c cursor) getResults(lines []string) string {
	var line string
	for c.next() {
		line += string(lines[c.y][c.x])
	}
	return line
}
