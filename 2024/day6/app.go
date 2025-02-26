package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type position struct {
	x, y int
}

type caseContent string

const (
	obstacle     caseContent = "obstacle"
	empty        caseContent = "empty"
	visited      caseContent = "visited"
	visitedTwice caseContent = "visitedTwice"
	outOfBounds  caseContent = "outOfBounds"
)

type guardMap struct {
	content   [][]caseContent
	guardian  *guardian
	endResult parcourResult
	visited   int
}

type parcourResult string

const (
	unknown parcourResult = "unknown"
	success parcourResult = "success"
	twice   parcourResult = "twice"
)

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
				result.content[i][j] = visited
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
	case empty:
		g.guardian.position = positionToTry
		g.content[g.guardian.x][g.guardian.y] = visited
	case visited, visitedTwice:
		g.guardian.position = positionToTry
		g.content[g.guardian.x][g.guardian.y] = visitedTwice
		g.visited++
	case obstacle:
		g.guardian.facing = Rotate(g.guardian.facing)
		return g.Tick()
	case outOfBounds:
		g.endResult = success
		return false
	default:
		panic("Invalid case content")
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
func (g *guardMap) CountVisitedTwice() int {
	count := 0
	for _, line := range (*g).content {
		for _, c := range line {
			if c == visitedTwice {
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
	for i, line := range (*g).content {
		for j, c := range line {
			if i == g.guardian.x && j == g.guardian.y {
				result.WriteRune('^')
				continue
			}
			switch c {
			case empty:
				result.WriteRune('.')
			case obstacle:
				result.WriteRune('#')
			case visited:
				result.WriteRune('X')
			case visitedTwice:
				result.WriteRune('+')
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
	allEmptyPos := myMap.FindAllEmpty()
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(allEmptyPos))
	loopCounter := 0
	for _, emptyPos := range allEmptyPos {
		go func() {
			defer wg.Done()
			submap := myMap.Clone()
			submap.content[emptyPos.x][emptyPos.y] = obstacle
			for i := 0; i < 10000000; i++ {
				if !submap.Tick() {
					log.Println("break")
					return
				}
			}
			mu.Lock()
			defer mu.Unlock()
			loopCounter++
			log.Println("inc, current value: %d", loopCounter)
		}()
	}

	wg.Wait()

	fmt.Printf("loopCounter: %d", loopCounter)
	fmt.Printf("total solution tried: %d", len(allEmptyPos))
}

func (g *guardMap) Clone() *guardMap {
	content := make([][]caseContent, len(g.content))
	for i, line := range g.content {
		content[i] = make([]caseContent, len(line))
		copy(content[i], line)
	}
	return &guardMap{
		content: content,
		guardian: &guardian{
			facing:   g.guardian.facing,
			position: g.guardian.position,
		},
		endResult: unknown,
		visited:   0,
	}
}

func (g *guardMap) FindAllEmpty() []position {
	var result []position
	for i, line := range (*g).content {
		for j, c := range line {
			if c == empty {
				result = append(result, position{i, j})
			}
		}
	}

	return result
}
func foo() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	myMap := Parse(string(content))
	for myMap.Tick() {
		if os.Getenv("APP_PRINT") != "" {
			fmt.Println(myMap.String())
			time.Sleep(10 * time.Millisecond)
		}
	}

	fmt.Println(myMap)
	fmt.Printf("EndResult: %s\n", myMap.endResult)
	fmt.Printf("Position: %v\n", myMap.guardian.position)
	fmt.Printf("VisitedTwice: %d\n", myMap.CountVisitedTwice())
	// obs := myMap.ExportObstacleUsed()
	// for _, o := range obs {
	// 	fmt.Println(o)
	// }
}
