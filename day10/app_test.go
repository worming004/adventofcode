package main

import "testing"

func TestTravel(t *testing.T) {
	tests := []struct {
		name             string
		mapStr           string
		want             int
		cursorsFoundWant int
		wantByCursor     []int
	}{
		{"simple",
			`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			36, 9,
			[]int{5, 6, 5, 3, 1, 3, 5, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Parse(tt.mapStr)
			cursors := m.FindAllStartingCursor()

			if len(cursors) != tt.cursorsFoundWant {
				t.Errorf("Cursors found: %d, expected %d", len(cursors), tt.cursorsFoundWant)
			}

			var result int = 0
			for i, c := range cursors {
				heightFound := []Position{}
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
				}

				c.SubsribeToHeightFoundVisitor(counterFunc)
				c.Travel(m)

				if len(heightFound) != tt.wantByCursor[i] {
					t.Errorf("Cursor %d: %d, expected %d", i, len(heightFound), tt.wantByCursor[i])
				}

				result += len(heightFound)
			}

			if result != tt.want {
				t.Errorf("Travel() = %d, want %d", result, tt.want)
			}

		})
	}
}
