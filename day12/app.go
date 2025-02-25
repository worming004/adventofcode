package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed try.txt
var try string

func main() {
	m := CreateMap(input)
	fmt.Println("Map Created")
	m.FindAndSetRegions()
	fmt.Println("Region found")
	totalPrice := 0
	for _, r := range m.Regions {
		totalPrice += r.Price()
	}

	fmt.Println(totalPrice)
	fmt.Println(len(m.Regions))
}

type Position struct {
	x, y int
}
type Cell struct {
	Position
	Value string
}

func (c Cell) IsOnPosition(p Position) bool {
	return c.Position == p
}

type Region struct {
	Cells []*Cell
	Value string
}

func (r Region) Area() int {
	return len(r.Cells)
}
func (r Region) Contains(p Position) bool {
	for _, cell := range r.Cells {
		if cell.IsOnPosition(p) {
			return true
		}
	}
	return false
}
func (r Region) Perimeter() int {
	total := 0
	for _, pos := range r.Cells {
		if !r.Contains(pos.North()) {
			total++
		}
		if !r.Contains(pos.East()) {
			total++
		}
		if !r.Contains(pos.South()) {
			total++
		}
		if !r.Contains(pos.West()) {
			total++
		}
	}

	return total
}

func (r Region) Price() int {
	return r.Area() * r.Perimeter()
}

func (p Position) North() Position {
	return Position{p.x, p.y - 1}
}
func (p Position) East() Position {
	return Position{p.x + 1, p.y}
}
func (p Position) South() Position {
	return Position{p.x, p.y + 1}
}
func (p Position) West() Position {
	return Position{p.x - 1, p.y}
}

func (m *Map) FindAndSetRegions() {
	findUnknownPos := func(mm *Map, rs []*Region) (Position, bool) {
		for _, cellRow := range mm.Cells {
			var currentPos Position
			for _, cell := range cellRow {
				currentPos = cell.Position
				isInRegion := false
				for _, r := range rs {
					if r.Contains(currentPos) {
						isInRegion = true
					}
				}
				if !isInRegion {
					return currentPos, true
				}

			}
		}
		return Position{}, false
	}

	var AddNeighCellsToRegion func(m *Map, r *Region, currentCell *Cell)
	AddNeighCellsToRegion = func(m *Map, r *Region, currentCell *Cell) {
		neighb := []*Cell{}
		{
			var c *Cell
			var ok bool
			c, ok = m.GetCell(currentCell.North())
			if ok {
				neighb = append(neighb, c)
			}
			c, ok = m.GetCell(currentCell.East())
			if ok {
				neighb = append(neighb, c)
			}
			c, ok = m.GetCell(currentCell.South())
			if ok {
				neighb = append(neighb, c)
			}
			c, ok = m.GetCell(currentCell.West())
			if ok {
				neighb = append(neighb, c)
			}
		}

		for _, n := range neighb {
			if n.Value == currentCell.Value {
				if !r.Contains(n.Position) {
					r.Cells = append(r.Cells, n)
					AddNeighCellsToRegion(m, r, n)
				}
			}
		}
	}

	for unknownPos, ok := findUnknownPos(m, m.Regions); ok; unknownPos, ok = findUnknownPos(m, m.Regions) {
		cell, ok := m.GetCell(unknownPos)
		if ok {
			newRegion := Region{Cells: []*Cell{cell}, Value: cell.Value}
			m.Regions = append(m.Regions, &newRegion)
			AddNeighCellsToRegion(m, &newRegion, cell)
		}
	}
}

type Map struct {
	Regions []*Region
	Input   string
	Cells   [][]*Cell
}

func (m Map) GetCell(p Position) (*Cell, bool) {
	if p.x < 0 || p.y < 0 || p.x >= len(m.Cells[0]) || p.y >= len(m.Cells) {
		return nil, false
	}
	return m.Cells[p.y][p.x], true
}

func CreateMap(input string) Map {
	m := Map{Regions: []*Region{}, Input: input}
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	m.Cells = make([][]*Cell, len(lines))
	for y, line := range lines {
		m.Cells[y] = make([]*Cell, len(line))
		for x, v := range line {
			m.Cells[y][x] = &Cell{Position{x, y}, string(v)}
		}
	}

	return m
}

func (m Map) GetRegionsByValue(v string) []Region {

	var regions []Region
	for _, r := range m.Regions {
		if r.Value == v {
			regions = append(regions, *r)
		}
	}

	return regions

}
