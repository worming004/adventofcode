package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Antenna struct {
	Position
	antennaType rune
}

type Antinode bool

type Position struct {
	x, y int
}

type step struct {
	sizeX, sizeY int
	a1, a2       Antenna
	anitnode     Position
}

func (s step) Print() string {
	sb := strings.Builder{}
	for y := 0; y < s.sizeY; y++ {
		for x := 0; x < s.sizeX; x++ {
			if x == s.a1.x && y == s.a1.y {
				sb.WriteRune(s.a1.antennaType)
			} else if x == s.a2.x && y == s.a2.y {
				sb.WriteRune(s.a2.antennaType)
			} else if x == s.anitnode.x && y == s.anitnode.y {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

type Gamemap struct {
	sizeX, sizeY int
	antennas     map[rune][]Antenna
	antinodes    map[Position]Antinode

	history []step
}

func Parse(input string) *Gamemap {
	result := new(Gamemap)
	result.antennas = make(map[rune][]Antenna)
	result.antinodes = make(map[Position]Antinode)

	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	result.sizeY = len(lines)
	for y, line := range lines {
		result.sizeX = len(line)
		for x, r := range line {
			if r == '.' {
				continue
			}

			result.antennas[r] = append(result.antennas[r], Antenna{Position{x, y}, r})
		}
	}
	return result
}

func (g *Gamemap) SeekAntinodes() {
	for _, antenna := range g.antennas {
		g.FindAntinodeByAntennaType(antenna[0], antenna[1:])
	}
}

func (g *Gamemap) CountAntinode() int {
	return len(g.antinodes)
}

func (g *Gamemap) FindAntinodeByAntennaType(a Antenna, as []Antenna) {
	if len(as) == 0 {
		return
	}
	for _, sub := range as {
		g.FindAntinodeWith2Antennas(a, sub)
	}
	g.FindAntinodeByAntennaType(as[0], as[1:])
}

func (g *Gamemap) FindAntinodeWith2Antennas(a1, a2 Antenna) {
	//antinode1
	{
		position := FindAntinodeWith2Antennas(a1, a2)
		g.addAntinodeIfNotExists(Antinode(true), position)
		g.history = append(g.history, step{g.sizeX, g.sizeY, a1, a2, position})
	}

	// antinode2
	{
		position := FindAntinodeWith2Antennas(a2, a1)
		g.addAntinodeIfNotExists(Antinode(true), position)
		g.history = append(g.history, step{g.sizeX, g.sizeY, a1, a2, position})
	}
}

func FindAntinodeWith2Antennas(a1, a2 Antenna) Position {
	x, y, diffx, diffy := 0, 0, 0, 0
	diffx = a2.x - a1.x
	diffy = a2.y - a1.y
	x = a1.x - diffx
	y = a1.y - diffy

	return Position{x, y}
}

func (g *Gamemap) addAntinodeIfNotExists(a Antinode, p Position) {
	if p.x < 0 || p.x >= g.sizeX || p.y < 0 || p.y >= g.sizeY {
		log.Printf("antinode out of bounds, %v", p)
		return
	}

	g.antinodes[p] = a
}

func main() {
	entries, err := os.ReadDir("steps")

	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		os.Remove(entry.Name())
	}

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	gamemap := Parse(string(input))
	gamemap.SeekAntinodes()

	for i, s := range gamemap.history {
		os.WriteFile(fmt.Sprintf("steps/step%d.txt", i), []byte(s.Print()), 0644)
		fmt.Println(s.Print())
	}

	fmt.Println(gamemap.CountAntinode())
}
