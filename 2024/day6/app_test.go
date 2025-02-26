package main

import (
	"testing"
)

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

	if myMap.GetCasecontent(position{6, 4}) != visited {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{6, 4}))
	}
	expectedGuardPosition(t, myMap, position{6, 4})
}

func TestTickEmpty(t *testing.T) {
	input := `...
.^.
...
`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}

	expectedGuardPosition(t, myMap, position{1, 1})
	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	expectedGuardPosition(t, myMap, position{0, 1})
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
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	expectedGuardPosition(t, myMap, position{1, 1})

	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.facing != est {
		t.Errorf("Expected north, got %s", myMap.guardian.facing)
	}

	expectedGuardPosition(t, myMap, position{1, 2})
}

func TestTickObstacleTwice(t *testing.T) {
	input := `.#.
.^#
...
`

	myMap := Parse(input)
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected guard, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.position != (position{1, 1}) {
		t.Errorf("Expected {1, 1} for guard position, got %v", myMap.guardian.position)
	}

	myMap.Tick()
	if myMap.GetCasecontent(position{1, 1}) != visited {
		t.Errorf("Expected visited, got %s", myMap.GetCasecontent(position{1, 1}))
	}
	if myMap.guardian.facing != south {
		t.Errorf("Expected north, got %s", myMap.guardian.facing)
	}
	expectedGuardPosition(t, myMap, position{2, 1})
}

func TestTickOutOfBounds(t *testing.T) {
	input := `...
.^.
...`

	myMap := Parse(input)
	expectedGuardPosition(t, myMap, position{1, 1})

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

func TestVisitedTwice(t *testing.T) {
	input := `......
.#....
.....#
......
......
.^..#.
`

	myMap := Parse(input)
	for myMap.Tick() {
		t.Log(myMap.String())
	}

	if myMap.GetCasecontent(position{4, 1}) != visitedTwice {
		t.Errorf("Expected visitedTwice, got %s", myMap.GetCasecontent(position{4, 1}))
		t.Log(myMap.String())
	}

}

func TestVisitederTwice_Short(t *testing.T) {
	input := `.X.
...
.^.
`
	myMap := Parse(input)
	for myMap.Tick() {
		t.Log(myMap.String())
	}

	if myMap.GetCasecontent(position{0, 1}) != visitedTwice {
		t.Errorf("Expected visitedTwice, got %s", myMap.GetCasecontent(position{0, 1}))
		t.Log(myMap.String())
	}

	visitedTwiceCount := myMap.CountVisitedTwice()
	if visitedTwiceCount != 1 {
		t.Errorf("Expected 1, got %d", visitedTwiceCount)
	}
}

func expectedGuardPosition(t *testing.T, g *guardMap, p position) {
	if g.guardian.position != p {
		t.Errorf("Expected %v, got %v", p, g.guardian.position)
	}
}
