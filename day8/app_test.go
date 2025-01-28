package main

import "testing"

var input string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestParse(t *testing.T) {

	gamemap := Parse(input)

	if gamemap.sizeX != 12 {
		t.Errorf("sizeX = %d; want 12", gamemap.sizeX)
	}
	if gamemap.sizeY != 12 {
		t.Errorf("sizeY = %d; want 12", gamemap.sizeY)
	}

	if gamemap.antennas['0'][0].antennaType != '0' {
		t.Errorf("antennaType = %c; want 0", gamemap.antennas['0'][0].antennaType)
	}

	if gamemap.antennas['0'][0].x != 8 {
		t.Errorf("x = %d; want 8", gamemap.antennas['0'][0].x)
	}
	if gamemap.antennas['0'][0].y != 1 {
		t.Errorf("y = %d; want 1", gamemap.antennas['0'][0].y)
	}

	if len(gamemap.antennas['0']) != 4 {
		t.Errorf("len(antennas['0']) = %d; want 4", len(gamemap.antennas['0']))
	}
	if len(gamemap.antennas) != 2 {
		t.Errorf("len(antennas) = %d; want 2", len(gamemap.antennas))
	}
}

var expected string = `.........
#..#..#..
.........
..CAD....
#.B.B.#..
..DAC....
.........
#..#..#..
.........`

func TestSeekAntinodes(t *testing.T) {
	gamemap := Parse(`........
........
........
..CAD...
..B.B...
..DAC...
........
........
........`,
	)

	gamemap.SeekAntinodes()

	if len(gamemap.antinodes) != 16 {
		t.Errorf("len(antinodes) = %d; want 16", len(gamemap.antinodes))
	}

	expectedPositions := []Position{
		//A
		{3, 1},
		{3, 7},
		//B
		{6, 4},
		{0, 4},
		//C
		{0, 1},
		{6, 7},
		//D
		{6, 1},
		{0, 7},
	}

	for _, expectedPosition := range expectedPositions {
		if a, ok := gamemap.antinodes[expectedPosition]; !ok && !bool(a) {
			t.Errorf("antinode %v not found", expectedPosition)
		}
	}

	count := gamemap.CountAntinode()
	if count != 16 {
		t.Errorf("count = %d; want 16", count)
	}
}

func TestSeekAntinodesOOB(t *testing.T) {
	gamemap := Parse(`...
.AA
...`)
	gamemap.SeekAntinodes()
	if len(gamemap.antinodes) != 3 {
		t.Errorf("len(antinodes) = %d; want 3", len(gamemap.antinodes))
	}
}
func TestSeekAntinodesResonance(t *testing.T) {
	gamemap := Parse(`............
.A..........
...A........
............
............
............
............`)
	gamemap.SeekAntinodes()
	if len(gamemap.antinodes) != 6 {
		t.Errorf("len(antinodes) = %d; want 4", len(gamemap.antinodes))
	}

	expectedPositions := []Position{
		{1, 1},
		{3, 2},
		{5, 3},
		{7, 4},
		{9, 5},
		{11, 6},
	}

	for _, expectedPosition := range expectedPositions {
		if a, ok := gamemap.antinodes[expectedPosition]; !ok && !bool(a) {
			t.Errorf("antinode %v not found", expectedPosition)
		}
	}
}
