package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func (r rules) IsValid(values []uint) (bool, rule) {
	for _, rule := range r {
		if !rule.IsValid(values) {
			return false, rule
		}
	}
	return true, rule{}
}

func (r rules) MakeValid(values []uint) []uint {
	result := make([]uint, len(values))
	copy(result, values)
	for {
		ok, rule := r.IsValid(result)
		if ok {
			break
		}
		a := rule.a
		pos := 0
		for i, v := range result {
			if v == a {
				pos = i
				break
			}
		}

		result = append(result[:pos], result[pos+1:]...)
		result = append([]uint{a}, result...)
	}
	return result
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
		if ok, _ := rules.IsValid(v); ok {
			result += getMiddle(v)
		}
	}

	fmt.Printf("First half result: %d\n", result)

	var result2 uint
	wg := sync.WaitGroup{}
	wg.Add(len(values))
	for _, v := range values {
		mu := sync.Mutex{}
		go func(sv []uint) {
			defer wg.Done()
			ok, _ := rules.IsValid(v)
			if ok {
				return
			}

			corrected := rules.MakeValid(v)
			mu.Lock()
			defer mu.Unlock()
			result2 += getMiddle(corrected)
		}(v)

	}
	wg.Wait()

	fmt.Printf("Second half result: %d\n", result2)
}
