package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x, y int
}

type caseContent string

const (
	obstacle    caseContent = "obstacle"
	guard       caseContent = "guard"
	empty       caseContent = "empty"
	visited     caseContent = "visited"
	outOfBounds caseContent = "outOfBounds"
)

type guardMap struct {
	content  [][]caseContent
	guardian *guardian
}

type facing string

const (
	north facing = "north"
	est   facing = "est"
	south facing = "south"
	west  facing = "west"
)

type guardian struct {
	facing facing
	position
}

func (g *guardMap) GetCasecontent(p position) caseContent {
	if p.x < 0 || p.x >= len((*g).content) || p.y < 0 || p.y >= len((*g).content[0]) {
		return outOfBounds
	}
	return (*g).content[p.x][p.y]
}

func Parse(input string) *guardMap {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	result := guardMap{}
	result.guardian = &guardian{}
	result.content = make([][]caseContent, len(lines))
	result.guardian.facing = north
	for i, line := range lines {
		result.content[i] = make([]caseContent, len(line))
		for j, c := range line {
			switch c {
			case '.':
				result.content[i][j] = empty
			case '#':
				result.content[i][j] = obstacle
			case 'X':
				result.content[i][j] = visited
			case '^':
				result.content[i][j] = guard
				result.guardian = &guardian{facing: north, position: position{i, j}}
			default:
				panic("Invalid character")
			}
		}
	}

	return &result
}

func (g *guardMap) Tick() bool {
	guardPos := g.guardian.position
	direction := g.guardian.facing
	positionToTry := position{}

	switch direction {
	case north:
		positionToTry = position{guardPos.x - 1, guardPos.y}
	case est:
		positionToTry = position{guardPos.x, guardPos.y + 1}
	case south:
		positionToTry = position{guardPos.x + 1, guardPos.y}
	case west:
		positionToTry = position{guardPos.x, guardPos.y - 1}
	default:
		panic("Invalid direction")
	}

	contentToTry := g.GetCasecontent(positionToTry)
	switch contentToTry {
	case empty, visited:
		g.content[guardPos.x][guardPos.y] = visited
		g.guardian.position = positionToTry
	case obstacle:
		g.guardian.facing = Rotate(g.guardian.facing)
		return g.Tick()
	case outOfBounds:
		g.content[guardPos.x][guardPos.y] = visited
		return false
	}

	return true
}

func (g *guardMap) CountVisited() int {
	count := 0
	for _, line := range (*g).content {
		for _, c := range line {
			if c == visited {
				count++
			}
		}
	}
	return count
}

func Rotate(d facing) facing {
	switch d {
	case north:
		return est
	case est:
		return south
	case south:
		return west
	case west:
		return north
	default:
		panic("Invalid direction")
	}
}

func (g *guardMap) String() string {
	var result strings.Builder
	for _, line := range (*g).content {
		for _, c := range line {
			switch c {
			case empty:
				result.WriteRune('.')
			case obstacle:
				result.WriteRune('#')
			case guard:
				result.WriteRune('^')
			case visited:
				result.WriteRune('X')
			}
		}
		result.WriteRune('\n')
	}
	return result.String()
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	myMap := Parse(string(content))
	for myMap.Tick() {
		log.Println(myMap)
	}
	count := myMap.CountVisited()
	fmt.Println(count)
}
