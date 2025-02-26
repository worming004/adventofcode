package main

import (
	"aventofcode2024/utils"
	_ "embed"
	"fmt"
	"regexp"

	"github.com/bit101/go-ansi"
)

//go:embed input.txt
var input string

func main() {
	boundX := 101
	boundY := 103
	robots := Parse(input, boundX, boundY)

	minim := 100000000
	var pattern []*Robot
	ticked := 0

	for i := 0; i < 100000; i++ {
		TickRobots(robots)
		sf := SafetyFactor(robots, boundX, boundY)
		if sf < minim {
			minim = sf
			pattern = Clone(robots)
			ticked = i
		}
	}

	PrintRobots(pattern, boundX, boundY)
	fmt.Printf("Second: %d", ticked+1)

}

func Clone(rs []*Robot) []*Robot {
	res := make([]*Robot, len(rs))
	for i, r := range rs {
		subr := *r
		res[i] = &subr
	}
	return res
}

type Position struct {
	x, y int
}
type Velocity struct {
	x, y int
}

type Robot struct {
	Pos Position
	Vel Velocity

	BoundX, BoundY int
}

func (r *Robot) TickNum(l int) {
	for i := 0; i < l; i++ {
		r.Tick()
	}
}

func TickRobots(rs []*Robot) {
	for i := range rs {
		rs[i].Tick()
	}
}
func (r *Robot) Tick() {
	r.Pos.x += r.Vel.x
	r.Pos.y += r.Vel.y

	if r.Pos.x < 0 {
		r.Pos.x += r.BoundX
	}
	if r.Pos.y < 0 {
		r.Pos.y += r.BoundY
	}
	if r.Pos.x >= r.BoundX {
		r.Pos.x -= r.BoundX
	}
	if r.Pos.y >= r.BoundY {
		r.Pos.y -= r.BoundY
	}
}

type Quadrant struct {
	TopLeft, TopRight, BottomLeft, BottomRight Position
}

func IsInQuad(q Quadrant, Robot *Robot) bool {
	return true &&
		Robot.Pos.x >= q.TopLeft.x &&
		Robot.Pos.x <= q.TopRight.x &&
		Robot.Pos.y >= q.TopLeft.y &&
		Robot.Pos.y <= q.BottomLeft.y
}

func PlusDivided(i int) int {
	interm := i / 2
	return i - interm
}

func SafetyFactor(rs []*Robot, BoundX, BoundY int) int {
	var safByQuad [4]int
	for i, q := range GetQuads(BoundX, BoundY) {
		for _, r := range rs {
			if IsInQuad(q, r) {
				safByQuad[i]++
			}
		}
	}

	return safByQuad[0] * safByQuad[1] * safByQuad[2] * safByQuad[3]
}

func Parse(s string, BoundX, BoundY int) []*Robot {
	lines := utils.SplitLines(s)
	result := make([]*Robot, len(lines))
	for i, line := range lines {
		result[i] = ParseLine(line, BoundX, BoundY)
	}

	return result
}
func ParseLine(line string, BoundX, BoundY int) *Robot {
	re := regexp.MustCompile(`-?\d+`)
	allStrs := re.FindAllString(line, -1)
	px := utils.Atoi(allStrs[0])
	py := utils.Atoi(allStrs[1])
	vx := utils.Atoi(allStrs[2])
	vy := utils.Atoi(allStrs[3])

	return &Robot{Position{px, py}, Velocity{vx, vy}, BoundX, BoundY}
}

func GetQuads(BoundX, BoundY int) [4]Quadrant {
	var quadrants [4]Quadrant
	quadrants[0] = Quadrant{Position{0, 0}, Position{BoundX/2 - 1, 0}, Position{0, BoundY/2 - 1}, Position{BoundX/2 - 1, BoundY/2 - 1}}
	quadrants[1] = Quadrant{Position{PlusDivided(BoundX), 0}, Position{BoundX - 1, 0}, Position{PlusDivided(BoundX), BoundY/2 - 1}, Position{BoundX - 1, BoundY/2 - 1}}
	quadrants[2] = Quadrant{Position{0, PlusDivided(BoundY)}, Position{BoundX/2 - 1, PlusDivided(BoundY)}, Position{0, BoundY - 1}, Position{BoundX/2 - 1, BoundY - 1}}
	quadrants[3] = Quadrant{Position{PlusDivided(BoundX), PlusDivided(BoundY)}, Position{BoundX - 1, PlusDivided(BoundY)}, Position{PlusDivided(BoundX), BoundY - 1}, Position{BoundX - 1, BoundY - 1}}
	return quadrants
}

func PrintRobots(rs []*Robot, bx, by int) {
	for y := 0; y < by; y++ {
		for x := 0; x < bx; x++ {
			count := 0
			for _, r := range rs {
				if r.Pos.x == x && r.Pos.y == y {
					count++
				}
			}
			if count > 0 {
				ansi.Printf(ansi.Red, "%d", count)
			} else {
				ansi.Printf(ansi.Green, "0")
			}
		}
		fmt.Printf("\n")
	}
}
