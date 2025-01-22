package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	sm := NewState()

	result := 0

	for s.Scan() {
		singleChar := s.Text()

		rns := []rune(singleChar)
		sm.readRune(rns[0])
		if sm.IsEnd() {
			sm.logger.Printf("result found, ongoing first: %d, second: %d\n", sm.first, sm.second)
			a, b := sm.GetValues()
			result += a * b
		}
	}

	fmt.Println(result)
}

type state rune

const (
	nothing   state = 'n'
	m         state = 'm'
	u         state = 'u'
	l         state = 'l'
	enter     state = '('
	first     state = 'f'
	separator state = 's'
	second    state = 'o'
	end       state = ')'
)

type StateMachine struct {
	logger    *log.Logger
	state     state
	firstStr  string
	first     int
	secondStr string
	second    int
}

func NewState() *StateMachine {
	sm := &StateMachine{}
	sm.logger = log.New(os.Stdout, "state-machine: ", log.LstdFlags)
	sm.Reinit()
	return sm
}

func (sm *StateMachine) Reinit() {
	sm.state = nothing
	sm.first = 0
	sm.firstStr = ""
	sm.second = 0
	sm.secondStr = ""
}

func (sm *StateMachine) readRune(r rune) {
	sm.logger.Printf("state: %c, rune: %c\n", sm.state, r)
	if sm.state == nothing && r == 'm' {
		sm.state = m
	} else if sm.state == m && r == 'u' {
		sm.state = u
	} else if sm.state == u && r == 'l' {
		sm.state = l
	} else if sm.state == l && r == '(' {
		sm.state = enter
	} else if sm.state == enter && unicode.IsDigit(r) {
		sm.state = first
		sm.firstStr = string(r)
	} else if sm.state == first && unicode.IsDigit(r) {
		sm.firstStr = sm.firstStr + string(r)
	} else if sm.state == first && r == ',' {
		sm.state = separator
	} else if sm.state == separator && unicode.IsDigit(r) {
		sm.state = second
		sm.secondStr = string(r)
	} else if sm.state == second && unicode.IsDigit(r) {
		sm.secondStr = sm.secondStr + string(r)
	} else if sm.state == second && r == ')' {
		var err error
		sm.first, err = strconv.Atoi(sm.firstStr)
		if err != nil {
			panic(err)
		}
		sm.second, err = strconv.Atoi(sm.secondStr)
		if err != nil {
			panic(err)
		}
		sm.state = end
	} else {
		if sm.state != nothing {
			sm.Reinit()
			sm.readRune(r)
		}
	}
}

func (sm *StateMachine) IsEnd() bool {
	return sm.state == end
}

func (sm *StateMachine) GetValues() (int, int) {
	return sm.first, sm.second
}
