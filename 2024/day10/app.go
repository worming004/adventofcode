package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	m := Parse(input)
	cursors := m.FindAllStartingCursor()
	ss := NewStateForPrint()
	total := 0
	totalScore := 0

	for _, c := range cursors {
		c.SubsribeToTravelVisitor(BuildVisitor(ss, m))
		heightFound := []Position{}
		counter := 0
		counterFunc := func(p Position) {
			alreadyFound := false
			for _, v := range heightFound {
				if v == p {
					alreadyFound = true
					break
				}
			}

			if !alreadyFound {
				heightFound = append(heightFound, p)
			}
			counter++
		}
		c.HeightFoundVisitor = counterFunc
		c.Travel(m)

		total += len(heightFound)
		totalScore += counter
	}

	fmt.Println(total)
	fmt.Println(totalScore)

}

type TrailMap struct {
	Places [][]int
}

func Parse(input string) TrailMap {
	var result TrailMap
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	result.Places = make([][]int, len(lines))
	for y, line := range lines {
		result.Places[y] = make([]int, len(line))

		for x, r := range line {
			result.Places[y][x] = rtoi(r)
		}
	}
	return result
}

func rtoi(r rune) int {
	s := string(r)
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

type Cursor struct {
	CurrentPos         Position
	CurrentVal         int
	TravelVisitor      func(Cursor)
	HeightFoundVisitor func(Position)
}

func (c *Cursor) SubsribeToTravelVisitor(f func(Cursor)) {
	currentF := c.TravelVisitor
	c.TravelVisitor = func(c Cursor) {
		if currentF != nil {
			currentF(c)
		}
		f(c)
	}
}
func (c *Cursor) SubsribeToHeightFoundVisitor(f func(Position)) {
	currentF := c.HeightFoundVisitor
	c.HeightFoundVisitor = func(p Position) {
		if currentF != nil {
			currentF(p)
		}
		f(p)
	}
}

// Go north. Return true if new position is valid. Return false if invalid or find an height
func (c Cursor) North(m TrailMap) (Cursor, bool) {
	if c.CurrentPos.y == 0 {
		return c, false
	}
	newPos := Position{c.CurrentPos.x, c.CurrentPos.y - 1}
	newPlace := m.GetValAt(newPos)
	return c.Decision(newPos, newPlace)
}

// Go east. Return true if new position is valid. Return false if invalid or find an height
func (c Cursor) East(m TrailMap) (Cursor, bool) {
	if c.CurrentPos.x == len(m.Places[0])-1 {
		return c, false
	}
	newPos := Position{c.CurrentPos.x + 1, c.CurrentPos.y}
	newPlace := m.GetValAt(newPos)
	return c.Decision(newPos, newPlace)
}

// Go South. Return true if new position is valid. Return false if invalid or find an height
func (c Cursor) South(m TrailMap) (Cursor, bool) {
	if c.CurrentPos.y == len(m.Places)-1 {
		return c, false
	}
	newPos := Position{c.CurrentPos.x, c.CurrentPos.y + 1}
	newPlace := m.GetValAt(newPos)
	return c.Decision(newPos, newPlace)
}

// Go West. Return true if new position is valid. Return false if invalid or find an height
func (c Cursor) West(m TrailMap) (Cursor, bool) {
	if c.CurrentPos.x == 0 {
		return c, false
	}
	newPos := Position{c.CurrentPos.x - 1, c.CurrentPos.y}
	newPlace := m.GetValAt(newPos)
	return c.Decision(newPos, newPlace)
}

func (c Cursor) Decision(newPos Position, newPlace int) (Cursor, bool) {
	if newPlace != c.CurrentVal+1 {
		return c, false
	}

	if newPlace == 9 {
		if c.HeightFoundVisitor != nil {
			c.HeightFoundVisitor(newPos)
		}
		newC := c
		newC.CurrentPos = newPos
		newC.CurrentVal = newPlace
		if newC.TravelVisitor != nil {
			newC.TravelVisitor(newC)
		}
		return c, false
	}

	c.CurrentPos = newPos
	c.CurrentVal = newPlace
	return c, true
}

// Travel returns the list of all the height found
func (c Cursor) Travel(m TrailMap) []Position {
	if c.TravelVisitor != nil {
		c.TravelVisitor(c)
	}
	var ok = true
	var newC Cursor
	var heightFound []Position
	newC, ok = c.North(m)
	if ok {
		newC.Travel(m)
	}
	newC, ok = c.East(m)
	if ok {
		newC.Travel(m)
	}
	newC, ok = c.South(m)
	if ok {
		newC.Travel(m)
	}
	newC, ok = c.West(m)
	if ok {
		newC.Travel(m)
	}

	return heightFound
}

func (m TrailMap) FindAllStartingCursor() []Cursor {
	var result []Cursor
	for y, line := range m.Places {
		for x, place := range line {
			if place == 0 {
				result = append(result, Cursor{CurrentPos: Position{x, y}, CurrentVal: 0})
			}
		}
	}
	return result
}

func (m TrailMap) GetValAt(p Position) int {
	return m.Places[p.y][p.x]
}

type Position struct {
	x, y int
}
