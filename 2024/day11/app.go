package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var stdoutLogger = log.New(os.Stdout, "", log.LstdFlags)
var nullLogger = log.New(io.Discard, "", log.LstdFlags)

func main() {
	loop, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	i := Parse(input)
	state := NewState(i, loop, WithLogger(*nullLogger))
	total := state.Blink()

	fmt.Println(total)
}

type Stone int

func NewStone(i int) Stone {
	return Stone(i)
}

func Blink(s Stone) []Stone {
	i := int(s)
	if i == 0 {
		return []Stone{1}
	}

	digitCount := digitCount(i)
	if digitCount%2 == 0 {
		v := split(i, digitCount)
		return []Stone{NewStone(v.a), NewStone(v.b)}
	}

	return []Stone{Stone(i * 2024)}
}

func digitCount(i int) int {
	if i == 0 {
		return 0
	}
	result := 0
	for {
		i /= 10
		result++
		if i == 0 {
			return result
		}
	}

}

func split(i, digitCount int) struct{ a, b int } {
	if digitCount%2 != 0 {
		panic("digitCount must be even")
	}
	middle := digitCount / 2

	dix := int(math.Pow(10, float64(middle)))
	b := i % dix
	a := (i - b) / dix
	return struct{ a, b int }{a, b}
}

type AppState struct {
	Stones   []Stone
	Logger   log.Logger
	Loop     int
	KnownRes map[StoneBlink]int
}

func NewState(i []int, loop int, options ...AppStateOptions) AppState {
	stones := make([]Stone, len(i))
	for j, v := range i {
		stones[j] = NewStone(v)
	}
	defaultLogger := log.New(os.Stdout, "", log.LstdFlags)
	a := AppState{Stones: stones, Logger: *defaultLogger, Loop: loop, KnownRes: map[StoneBlink]int{}}
	for _, o := range options {
		o(&a)
	}
	return a
}
func WithLogger(l log.Logger) AppStateOptions {
	return func(a *AppState) {
		a.Logger = l
	}
}

func Parse(input string) []int {
	splitted := strings.Split(input, " ")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}

	if strings.Contains(splitted[len(splitted)-1], "\n") {
		splitted[len(splitted)-1] = strings.Replace(splitted[len(splitted)-1], "\n", "", -1)
	}

	result := make([]int, len(splitted))
	for i, s := range splitted {
		result[i] = atoi(s)
	}

	return result
}
func atoi(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

type AppStateOptions func(*AppState)
type StoneBlink struct {
	Stone Stone
	Blink int
}

func (a *AppState) Blink() int {
	if a.Loop == 0 {
		return a.Length()
	}
	total := 0
	for _, s := range a.Stones {
		subTotal := 0
		if v, ok := a.KnownRes[StoneBlink{s, a.Loop}]; ok {
			subTotal = v
		} else {
			sub := AppState{Logger: a.Logger, KnownRes: a.KnownRes}
			sub.Loop = a.Loop - 1
			sub.Stones = Blink(s)
			subTotal = sub.Blink()
		}
		total += subTotal
		a.KnownRes[StoneBlink{s, a.Loop}] = subTotal
	}

	return total
}

func (a *AppState) Length() int {
	return len(a.Stones)
}
