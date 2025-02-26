package main

import (
	"aventofcode2024/debug"
	"aventofcode2024/utils"

	"github.com/bit101/go-ansi"
)

type Position struct {
	x, y int
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

func main() {
	input := "#####\n#..@#\n#.O.#\n#####"
	m := Parse(input)

	m.Print()
}

func Parse(i string) *Map {
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

func (m *Map) Print() {
	for _, row := range m.Cells {
		debug.PrintLineReturn()
		for _, cell := range row {
			if m.RobotPosition == cell.Position {
				ansi.Printf(ansi.Green, "@")
			} else {
				switch cell.Content {
				case Empty:
					print(".")
				case Wall:
					print("#")
				case Box:
					ansi.Printf(ansi.Yellow, "O")
				}
			}
		}
	}
}
