package main

import (
	"testing"
)

func TestFindRegion(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	m := CreateMap(input)

	m.FindAndSetRegions(input)
	regions := m.Regions
	if len(regions) != 11 {

		t.Errorf("Expected 11 regions, got %d", len(regions))
	}

	rRegion := m.GetRegionsByValue("R")
	if len(rRegion) != 1 {
		t.Errorf("Expected 1 regions with value R, got %d", len(rRegion))
	}
	if len(rRegion[0].Cells) != 12 {
		t.Errorf("Expected 12 Cells for region with value R, got %d", len(rRegion[0].Cells))
	}
}

func C(x, y int, v string) *Cell {
	return &Cell{Position{x, y}, v}
}

func TestPerimeter(t *testing.T) {
	type args struct {
		r Region
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{Region{[]*Cell{C(0, 0, "A")}, "A"}}, 4},
		{"AA-AA", args{Region{[]*Cell{C(0, 0, "A"), C(1, 0, "A"), C(0, 1, "A"), C(1, 1, "A")}, "A"}}, 8},
		{"AAA-AA", args{Region{[]*Cell{C(0, 0, "A"), C(1, 0, "A"), C(2, 0, "A"), C(0, 1, "A"), C(1, 1, "A")}, "A"}}, 10},
		{"AAA-AA-AAA", args{Region{[]*Cell{C(0, 0, "A"), C(1, 0, "A"), C(2, 0, "A"), C(0, 1, "A"), C(1, 1, "A"), C(0, 2, "A"), C(1, 2, "A"), C(2, 2, "A")}, "A"}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.r.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
