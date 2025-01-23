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
	t := newTab(string(content))
	r := t.findOccurences()
	fmt.Printf("Result: %d", r)
}

type tab [][]rune
type position struct {
	x, y int
}

func (t tab) findOccurences() int {
	var result int
	length := len(t)
	for x := range t {
		if x == 0 {
			continue
		}
		if x == length-1 {
			continue
		}
		for y := range t[x] {
			if y == 0 {
				continue
			}
			if y == length-1 {
				continue
			}
			if t[y][x] == 'A' {
				//diag1
				d1 := ""
				d1 += string(t[y-1][x-1])
				d1 += string(t[y+1][x+1])
				if d1 == "MS" || d1 == "SM" {
					d2 := ""
					//diag2
					d2 += string(t[y-1][x+1])
					d2 += string(t[y+1][x-1])

					if d2 == "MS" || d2 == "SM" {
						result++
					}
				}
			}
		}
	}

	return result

}

func newTab(content string) tab {
	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	var result tab
	for i := range lines {
		result = append(result, []rune(lines[i]))
	}
	return result
}
