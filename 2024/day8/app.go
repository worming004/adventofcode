package main

import (
	"errors"
	"fmt"
	"io/fs"
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
		//antinode1
		g.FindAntinodeWith2Antennas(a, sub)
		//antinode2
		g.FindAntinodeWith2Antennas(sub, a)
	}
	g.FindAntinodeByAntennaType(as[0], as[1:])
}

func (g *Gamemap) FindAntinodeWith2Antennas(a1, a2 Antenna) {
	g.antinodes[a1.Position] = Antinode(true)
	g.antinodes[a2.Position] = Antinode(true)
	position := FindAntinodeWith2Antennas(a1, a2)
	r := g.addAntinodeIfNotExists(Antinode(true), position)
	g.history = append(g.history, step{g.sizeX, g.sizeY, a1, a2, position})
	if r {
		g.FindAntinodeWith2Antennas(Antenna{position, a1.antennaType}, a1)
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

func (g *Gamemap) addAntinodeIfNotExists(a Antinode, p Position) bool {
	if p.x < 0 || p.x >= g.sizeX || p.y < 0 || p.y >= g.sizeY {
		log.Printf("antinode out of bounds, %v", p)
		return false
	}

	g.antinodes[p] = a
	return true
}

func main() {
	entries, err := os.ReadDir("steps")

	if errors.Is(err, fs.ErrNotExist) {
		log.Printf("steps dir does not exist")
		os.Mkdir("steps", 0744)
	} else if err != nil {
		panic(err)
	} else {
		for _, entry := range entries {
			os.Remove(entry.Name())
		}
	}

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	gamemap := Parse(string(input))
	gamemap.SeekAntinodes()

	for i, s := range gamemap.history {
		err := os.WriteFile(fmt.Sprintf("steps/step%d.txt", i), []byte(s.Print()), 0644)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(gamemap.CountAntinode())

}
