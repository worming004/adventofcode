package main

import (
	"fmt"

	"github.com/bit101/go-ansi"
)

func (ss *StateForPrint) Print() {
	for y, line := range ss.TrailMap.Places {
		for x, place := range line {
			currentPos := Position{x, y}
			if ss.CurrentPosition == currentPos {
				ansi.Printf(ansi.Cyan, "%d", place)
			} else if ss.IsVisited(currentPos) {
				ansi.Printf(ansi.Yellow, "%d", place)
			} else {
				ansi.Printf(ansi.White, "%d", place)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func (ss *StateForPrint) IsVisited(p Position) bool {
	for _, v := range ss.VisitedPos {
		if v == p {
			return true
		}
	}
	return false
}

type StateForPrint struct {
	VisitedPos      []Position
	HeightFound     []Position
	CurrentPosition Position
	TrailMap        TrailMap
}

func NewStateForPrint() *StateForPrint {
	return &StateForPrint{}
}

func BuildVisitor(ss *StateForPrint, m TrailMap) func(Cursor) {
	return func(c Cursor) {
		ss.TrailMap = m
		ss.CurrentPosition = c.CurrentPos
		ss.VisitedPos = append(ss.VisitedPos, c.CurrentPos)
		ss.VisitedPos = removeDuplicate(ss.VisitedPos)

		ss.Print()
	}
}

// remove duplicate from any slice
func removeDuplicate[T comparable](elements []T) []T {
	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements); j++ {
			if elements[i] == elements[j] {
				elements = append(elements[:j], elements[j+1:]...)
				j--
			}
		}
	}

	return elements
}
