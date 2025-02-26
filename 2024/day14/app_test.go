package main

import (
	"reflect"
	"testing"
)

const BoundX = 101
const BoundY = 103

func TestRobot_Tick(t *testing.T) {
	type fields struct {
		Pos    Position
		Vel    Velocity
		BoundX int
		BoundY int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"From aoc", fields{Position{2, 4}, Velocity{2, -3}, 11, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Pos:    tt.fields.Pos,
				Vel:    tt.fields.Vel,
				BoundX: tt.fields.BoundX,
				BoundY: tt.fields.BoundY,
			}
			r.Tick()
			if r.Pos.x != 4 {
				t.Errorf("Expected x to be 4, got %d", r.Pos.x)
			}
			if r.Pos.y != 1 {
				t.Errorf("Expected y to be 1, got %d", r.Pos.y)
			}
			r.TickNum(4)
			if r.Pos.x != 1 {
				t.Errorf("Expected x to be 1, got %d", r.Pos.x)
			}
			if r.Pos.y != 3 {
				t.Errorf("Expected y to be 3, got %d", r.Pos.y)
			}
		})
	}
}

func TestGetQuads(t *testing.T) {
	type args struct {
		BoundX int
		BoundY int
	}
	tests := []struct {
		name string
		args args
		want [4]Quadrant
	}{
		{"From aoc", args{11, 7}, [4]Quadrant{
			{Position{0, 0}, Position{4, 0}, Position{0, 2}, Position{4, 2}},
			{Position{6, 0}, Position{10, 0}, Position{6, 2}, Position{10, 2}},
			{Position{0, 4}, Position{4, 4}, Position{0, 6}, Position{4, 6}},
			{Position{6, 4}, Position{10, 4}, Position{6, 6}, Position{10, 6}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuads(tt.args.BoundX, tt.args.BoundY); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuads() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestIsInQuad(t *testing.T) {
	type args struct {
		q     Quadrant
		Robot *Robot
	}
	defaultQuad := Quadrant{Position{1, 1}, Position{5, 1}, Position{1, 3}, Position{5, 3}}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"In-XEdge", args{defaultQuad, &Robot{Pos: Position{1, 2}}}, true},
		{"In+XEdge", args{defaultQuad, &Robot{Pos: Position{5, 2}}}, true},
		{"In-YEdge", args{defaultQuad, &Robot{Pos: Position{2, 1}}}, true},
		{"In+YEdge", args{defaultQuad, &Robot{Pos: Position{2, 3}}}, true},
		{"Out-X", args{defaultQuad, &Robot{Pos: Position{0, 2}}}, false},
		{"Out+X", args{defaultQuad, &Robot{Pos: Position{6, 2}}}, false},
		{"Out-Y", args{defaultQuad, &Robot{Pos: Position{2, 0}}}, false},
		{"Out+Y", args{defaultQuad, &Robot{Pos: Position{2, 4}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInQuad(tt.args.q, tt.args.Robot); got != tt.want {
				t.Errorf("IsInQuad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	type args struct {
		line   string
		BoundX int
		BoundY int
	}
	tests := []struct {
		name string
		args args
		want *Robot
	}{
		{"default", args{"p=0,4 v=3,-3", 7, 11}, &Robot{Position{0, 4}, Velocity{3, -3}, 7, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line, tt.args.BoundX, tt.args.BoundY); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullScenaerio(t *testing.T) {
	BoundX := 11
	BoundY := 7

	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	robots := Parse(input, BoundX, BoundY)

	for _, r := range robots {
		r.TickNum(100)
	}

	sf := SafetyFactor(robots, BoundX, BoundY)

	if sf != 12 {
		t.Errorf("Expected safety factor to be 12, got %d", sf)
	}

	if t.Failed() {
		PrintRobots(robots, BoundX, BoundY)
	}
}
