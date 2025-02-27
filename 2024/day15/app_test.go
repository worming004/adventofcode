package main

import (
	_ "embed"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name      string
		args      args
		assertion func(*testing.T, *Map)
	}{
		{"Default parse", args{"####.\n#..#\n#.O#\n####"}, func(t *testing.T, m *Map) {
			if len(m.Cells) != 4 {
				t.Errorf("Expected 4 rows, got %d", len(m.Cells))
			}

			if m.Cells[2][2].Content != Box {
				t.Errorf("Expected box at 2,2, got %v", m.Cells[2][2].Content)
			}

			if m.Cells[0][0].Content != Wall {
				t.Errorf("Expected wall at 0,0, got %v", m.Cells[0][0].Content)
			}

			if m.Cells[1][1].Content != Empty {
				t.Errorf("Expected empty at 1,1, got %v", m.Cells[1][1].Content)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseMap(tt.args.i)
			tt.assertion(t, got)
		})
	}
}

func TestMap_BoxCellGoTo(t *testing.T) {
	type fields struct {
		input string
	}
	type args struct {
		startPos Position
		d        Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		assert func(*testing.T, *Map)
		want   bool
	}{
		{"Just Empty", fields{"#####\n#...#\n#.O.#\n#####"}, args{startPos: Position{2, 2}, d: North}, func(t *testing.T, m *Map) {
			if m.GetCellAtPosition(Position{x: 2, y: 1}).Content != Box {
				t.Errorf("Expected box at 2,1, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
			if m.GetCellAtPosition(Position{x: 2, y: 2}).Content != Empty {
				t.Errorf("Expected empty at 2,2, got %v", m.GetCellAtPosition(Position{x: 2, y: 2}).Content)
			}
		}, true},
		{"Push 3 boxes with empty space", fields{"#####\n#...#\n#.O.#\n#.O.#\n#.O.#\n#####"}, args{startPos: Position{2, 4}, d: North}, func(t *testing.T, m *Map) {
			if m.GetCellAtPosition(Position{x: 2, y: 1}).Content != Box {
				t.Errorf("Expected box at 2,1, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
			if m.GetCellAtPosition(Position{x: 2, y: 3}).Content != Box {
				t.Errorf("Expected empty at 2,3, got %v", m.GetCellAtPosition(Position{x: 2, y: 3}).Content)
			}
			if m.GetCellAtPosition(Position{x: 2, y: 4}).Content != Empty {
				t.Errorf("Expected empty at 2,4, got %v", m.GetCellAtPosition(Position{x: 2, y: 4}).Content)
			}
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ParseMap(tt.fields.input)
			got := m.BoxCellGoTo(m.GetCellAtPosition(tt.args.startPos), tt.args.d)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
			tt.assert(t, m)

			if t.Failed() {
				m.Print(false)
			}

		})
	}
}

func TestMap_RobotGoTo(t *testing.T) {
	type fields struct {
		input string
	}
	type args struct {
		d Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		assert func(*testing.T, *Map)
	}{
		{"Move robot on empty", fields{"#####\n#...#\n#.@.#\n#####"}, args{North}, func(t *testing.T, m *Map) {
			expectedPos := Position{x: 2, y: 1}
			if m.RobotPosition != expectedPos {
				t.Errorf("Expected robot at %v, got %v", expectedPos.String(), m.RobotPosition.String())
			}
		}},
		{"Move robot on wall", fields{"#####\n#.@.#\n#...#\n#####"}, args{North}, func(t *testing.T, m *Map) {
			expectedPos := Position{x: 2, y: 1}
			if m.RobotPosition != expectedPos {
				t.Errorf("Expected robot at %v, got %v", expectedPos.String(), m.RobotPosition.String())
			}
		}},
		{"Move robot on box then empty", fields{"#####\n#...#\n#.O.#\n#.@.#\n#####"}, args{North}, func(t *testing.T, m *Map) {
			expectedPos := Position{x: 2, y: 2}
			if m.RobotPosition != expectedPos {
				t.Errorf("Expected robot at %v, got %v", expectedPos.String(), m.RobotPosition.String())
			}

			if m.GetCellAtPosition(Position{x: 2, y: 1}).Content != Box {
				t.Errorf("Expected box at 2,1, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
			if m.GetCellAtPosition(Position{x: 2, y: 2}).Content != Empty {
				t.Errorf("Expected box at 2,2, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
		}},
		{"Move robot on box then wall", fields{"#####\n#.O.#\n#.O.#\n#.@.#\n#####"}, args{North}, func(t *testing.T, m *Map) {
			expectedPos := Position{x: 2, y: 3}
			if m.RobotPosition != expectedPos {
				t.Errorf("Expected robot at %v, got %v", expectedPos.String(), m.RobotPosition.String())
			}

			if m.GetCellAtPosition(Position{x: 2, y: 1}).Content != Box {
				t.Errorf("Expected box at 2,1, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
			if m.GetCellAtPosition(Position{x: 2, y: 2}).Content != Box {
				t.Errorf("Expected box at 2,2, got %v", m.GetCellAtPosition(Position{x: 2, y: 1}).Content)
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ParseMap(tt.fields.input)
			m.RobotGoTo(tt.args.d)
			tt.assert(t, m)
		})
	}
}

func TestMap_SumCoordinate(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Simple sum", fields{"#######\n#.....#\n#...O.#\n#######"}, 204},
		{"2 Sum", fields{"#######\n#....O#\n#...O.#\n#######"}, 309},
		{"From aoc small", fields{`########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`}, 2028},
		{"From aoc", fields{`##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`}, 10092},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ParseMap(tt.fields.input)
			if got := m.SumCoordinate(); got != tt.want {
				t.Errorf("Map.SumCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
