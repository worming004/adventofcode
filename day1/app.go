package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	a, b int
}

type pairs []pair

type input struct {
	cola, colb []int
}

func (i input) toPairs() pairs {
	sort.Sort(sort.IntSlice(i.cola))
	sort.Sort(sort.IntSlice(i.colb))
	result := pairs{}
	for idx := range i.cola {
		result = append(result, pair{i.cola[idx], i.colb[idx]})
	}
	return result
}

func (p pairs) distance() int {
	result := 0
	for _, p := range p {
		subr := p.a - p.b
		if subr < 0 {
			subr = -subr
		}
		result += subr
	}
	return result
}

func (p input) similarity() int {
	result := 0
	for _, a := range p.cola {
		for _, b := range p.colb {
			if a == b {
				result = result + a
			}
		}
	}
	return result
}

func main() {
	input := parse()
	pairs := input.toPairs()
	distance := pairs.distance()
	similarity := input.similarity()
	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("Similarity: %d\n", similarity)
}

func parse() input {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := input{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		splitted := strings.Split(t, " ")
		splitted = append(splitted[:1], splitted[3:]...)
		if len(splitted) != 2 {
			panic(fmt.Errorf("Invalid input, expected length of 2, got %d", len(splitted)))
		}

		a, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}
		result.cola = append(result.cola, a)
		result.colb = append(result.colb, b)
	}

	return result
}
