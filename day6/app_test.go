package main

import "testing"

func TestParse(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	myMap := Parse(input)

	if myMap.GetCasecontent(position{0, 0}) != empty {
		t.Errorf("Expected empty, got %s", myMap.GetCasecontent(position{0, 0}))
	}

	if myMap.GetCasecontent(position{3, 2}) != obstacle {
		t.Errorf("Expected obstacle, got %s", myMap.GetCasecontent(position{3, 2}))
	}

	if myMap.GetCasecontent(position{6, 4}) != guard {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{6, 4}))
	}
}

func TestTickEmpty(t *testing.T) {
	input := `...
.^.
...
`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != guard {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.facing != north {
		t.Errorf("Expected north, got %s", myMap.guardian.facing)
	}
	if myMap.guardian.position != (position{0, 1}) {
		t.Errorf("Expected {0, 1}, got %v", myMap.guardian.position)
	}
}

func TestTickObstacle(t *testing.T) {
	input := `.#.
.^.
...
`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != guard {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.facing != est {
		t.Errorf("Expected north, got %s", myMap.guardian.facing)
	}
	if myMap.guardian.position != (position{1, 2}) {
		t.Errorf("Expected {1, 2}, got %v", myMap.guardian.position)
	}
}

func TestTickObstacleTwice(t *testing.T) {
	input := `.#.
.^#
...
`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != guard {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.facing != south {
		t.Errorf("Expected north, got %s", myMap.guardian.facing)
	}
	if myMap.guardian.position != (position{2, 1}) {
		t.Errorf("Expected {2, 1}, got %v", myMap.guardian.position)
	}
}

func TestTickOutOfBounds(t *testing.T) {
	input := `...
.^.
...`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != guard {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	myMap.Tick()
	finalRes := myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	if myMap.GetCasecontent(position{0, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{0, 1}))
	}

	if finalRes {
		t.Errorf("Expected false, got true")
	}
}

func TestCount(t *testing.T) {
	input := `.XX
XX.
...`

	myMap := Parse(input)
	if myMap.CountVisited() != 4 {
		t.Errorf("Expected 4, got %d", myMap.CountVisited())
	}
}
