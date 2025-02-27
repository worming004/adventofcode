package main

import (
	"aventofcode2024/debug"
	"aventofcode2024/utils"
	"fmt"
	"strings"

	_ "embed"

	"github.com/bit101/go-ansi"
)

//go:embed input.txt
var input string

//go:embed try.txt
var try string

//go:embed try2.txt
var try2 string

const printAnsi = false
const printDebug = false

func main() {
	todo := input
	m, dirs := ParseAll(todo)
	splitted := utils.SplitLines(todo)
	dirLine := splitted[len(splitted)-1]
	fmt.Println(dirLine)
	m.Print(printAnsi)
	fmt.Println()

	for i, d := range dirs {
		m.RobotGoTo(d)
		if printDebug {
			fmt.Printf("Move: %s\n", string(dirLine[i]))
			m.Print(printAnsi)
			fmt.Println()
			fmt.Println()

		}
	}

	fmt.Printf("Sum of coordinates: %d\n", m.SumCoordinate())
}

type Position struct {
	x, y int
}

func (p Position) String() string {
	return fmt.Sprintf("(x:%d, y:%d)", p.x, p.y)
}

type CellContent string

const (
	Empty CellContent = "Empty"
	Box   CellContent = "Box"
	Wall  CellContent = "Wall"
)

type Cell struct {
	Position
	Content CellContent
}

type Map struct {
	Cells         [][]*Cell
	RobotPosition Position
}

func RuneToDirection(r rune) Direction {
	switch r {
	case '^':
		return North
	case 'v':
		return South
	case '>':
		return East
	case '<':
		return West
	}
	return nil
}

func ParseAll(i string) (*Map, []Direction) {

	splitted := utils.SplitLines(i)
	mapLine := strings.Join(splitted[:len(splitted)-2], "\n")
	m := ParseMap(mapLine)

	dirLine := splitted[len(splitted)-1]
	var directions []Direction
	for _, r := range dirLine {
		d := RuneToDirection(r)
		if d != nil {
			directions = append(directions, d)
		}
	}

	return m, directions
}

func ParseMap(i string) *Map {
	m := Map{}
	m.Cells = make([][]*Cell, len(utils.SplitLines(i)))
	lines := utils.SplitLines(i)
	for y, line := range lines {
		m.Cells[y] = make([]*Cell, len(line))
		for x, c := range line {
			switch c {
			case '.':
				m.Cells[y][x] = &Cell{Position{x, y}, Empty}
			case '#':
				m.Cells[y][x] = &Cell{Position{x, y}, Wall}
			case 'O':
				m.Cells[y][x] = &Cell{Position{x, y}, Box}
			case '@':
				m.Cells[y][x] = &Cell{Position{x, y}, Empty}
				m.RobotPosition = Position{x, y}
			}
		}
	}

	return &m
}

func (m *Map) RobotGoTo(d Direction) bool {
	posToTest := d(m.RobotPosition)
	if !m.IsInBoundary(posToTest) {
		panic("Out of boundary")
	}

	cellAtNextPos := m.GetCellAtPosition(posToTest)

	switch cellAtNextPos.Content {
	case Empty:
		m.RobotPosition = posToTest
		return true

	case Box:
		if m.BoxCellGoTo(cellAtNextPos, d) {
			m.RobotPosition = posToTest
			return true
		}
	}

	return false
}
func (m *Map) BoxCellGoTo(c *Cell, d Direction) bool {
	posToTest := d(c.Position)
	if !m.IsInBoundary(posToTest) {
		panic("Out of boundary")
	}

	cellAtNextPos := m.GetCellAtPosition(posToTest)
	switch cellAtNextPos.Content {
	case Empty:
		cellAtNextPos.Content = Box
		c.Content = Empty

		return true
	case Box:
		if m.BoxCellGoTo(cellAtNextPos, d) {
			cellAtNextPos.Content = Box
			c.Content = Empty

			return true
		}
	}

	return false
}

func (m *Map) GetCellAtPosition(p Position) *Cell {
	return m.Cells[p.y][p.x]
}

type Direction func(Position) Position

func North(p Position) Position {
	return Position{p.x, p.y - 1}
}
func South(p Position) Position {
	return Position{p.x, p.y + 1}
}
func East(p Position) Position {
	return Position{p.x + 1, p.y}
}

func West(p Position) Position {
	return Position{p.x - 1, p.y}
}

func (m *Map) IsInBoundary(p Position) bool {
	return true &&
		p.x >= 0 &&
		p.y >= 0 &&
		p.x < len(m.Cells[0]) &&
		p.y < len(m.Cells)
}

func (m *Map) Print(isansi bool) {
	for y, row := range m.Cells {
		debug.PrintLineReturn()
		for x, cell := range row {
			if cell == nil {
				fmt.Printf("position %d %d is nil", x, y)
			}
			if m.RobotPosition == cell.Position {
				if isansi {
					ansi.Printf(ansi.Green, "@")
				} else {
					fmt.Printf("@")
				}
				continue
			}

			switch cell.Content {
			case Empty:
				fmt.Printf(".")
			case Wall:
				fmt.Printf("#")
			case Box:
				if isansi {
					ansi.Printf(ansi.Yellow, "O")
				} else {
					fmt.Printf("O")
				}
			}
		}
	}
}

func (m *Map) SumCoordinate() int {
	total := 0
	for y, row := range m.Cells {
		for x, cell := range row {
			if cell.Content == Box {
				total += x + 100*y
			}
		}
	}

	return total
}
