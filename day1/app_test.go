package main

import "testing"

var i = input{
	cola: []int{3, 4, 2, 8, 9},
	colb: []int{3, 4, 2, 5, 7},
}

func TestToPairs(t *testing.T) {
	p := i.toPairs()
	if len(p) != 5 {
		t.Errorf("Expected 3 pairs, got %d", len(p))
	}
	if p[0].a != 2 || p[0].b != 2 {
		t.Errorf("Expected 1, 1, got %d, %d", p[0].a, p[0].b)
	}
	if p[1].a != 3 || p[1].b != 3 {
		t.Errorf("Expected 2, 2, got %d, %d", p[1].a, p[1].b)
	}
	if p[2].a != 4 || p[2].b != 4 {
		t.Errorf("Expected 3, 3, got %d, %d", p[2].a, p[2].b)
	}
}

func TestDistance(t *testing.T) {
	p := i.toPairs()
	if p.distance() != 5 {
		t.Errorf("Expected 5, got %d", p.distance())
	}
}

var inputForSimilarity = input{
	cola: []int{3, 4, 2, 1, 3, 3},
	colb: []int{4, 3, 5, 3, 9, 3},
}

func TestSimilarity(t *testing.T) {
	if inputForSimilarity.similarity() != 31 {
		t.Errorf("Expected 0, got %d", inputForSimilarity.similarity())
	}
}
