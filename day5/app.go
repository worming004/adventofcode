package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rules []rule
type rule pair
type pair struct {
	a, b uint
}

func (r rule) IsValid(values []uint) bool {
	var pairsToVerify = []pair{}
	for i := range values {
		for j := i; j < len(values); j++ {
			pairsToVerify = append(pairsToVerify, pair{values[i], values[j]})
		}
	}

	for _, p := range pairsToVerify {
		if p.a == r.b && p.b == r.a {
			return false
		}
	}
	return true
}

func (r rules) IsValid(values []uint) bool {
	for _, rule := range r {
		if !rule.IsValid(values) {
			return false
		}
	}
	return true
}

func getMiddle(values []uint) uint {
	length := len(values)

	return values[(length-1)/2]
}

func Parse(input string) (rules, [][]uint) {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var step string = "first"

	var rRules rules
	var rTab [][]uint
	for i := 0; i < len(lines); i++ {
		if step == "first" {

			if lines[i] == "" {
				step = "second"
				continue
			}
			vals := strings.Split(lines[i], "|")
			rRules = append(rRules, rule{atoi(vals[0]), atoi(vals[1])})
		}

		if step == "second" {
			if lines[i] == "" {
				continue
			}

			vals := strings.Split(lines[i], ",")
			tab := []uint{}
			for _, v := range vals {
				tab = append(tab, atoi(v))
			}
			rTab = append(rTab, tab)
		}

	}
	return rRules, rTab
}

func atoi(s string) uint {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return uint(r)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rules, values := Parse(string(content))
	var result uint
	for _, v := range values {
		if rules.IsValid(v) {
			result += getMiddle(v)
		}
	}

	fmt.Println(result)
}
