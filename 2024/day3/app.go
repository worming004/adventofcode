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

	d    state = 'd'
	o    state = 'o'
	n    state = 'n'
	tick state = '\''
	t    state = 't'
)

type StateMachine struct {
	logger    *log.Logger
	state     state
	firstStr  string
	first     int
	secondStr string
	second    int

	doState        bool
	evaluatingDo   bool
	evaluatingDont bool
}

func NewState() *StateMachine {
	sm := &StateMachine{}
	sm.logger = log.New(os.Stdout, "state-machine: ", log.LstdFlags)
	sm.Reinit()
	sm.doState = true
	return sm
}

func (sm *StateMachine) Reinit() {
	sm.state = nothing
	sm.first = 0
	sm.firstStr = ""
	sm.second = 0
	sm.secondStr = ""
	sm.evaluatingDo = false
	sm.evaluatingDont = false
}

func (sm *StateMachine) readRune(r rune) {
	sm.logger.Printf("state: %c, rune: %c\n", sm.state, r)
	// mul part
	if sm.state == nothing && r == 'm' {
		sm.state = m
	} else if sm.state == m && r == 'u' {
		sm.state = u
	} else if sm.state == u && r == 'l' {
		sm.state = l
	} else if sm.state == l && r == '(' && sm.evaluatingDo == false && sm.evaluatingDont == false {
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
		// do part
	} else if sm.state == nothing && r == 'd' {
		sm.state = d
	} else if sm.state == d && r == 'o' {
		sm.state = o
		sm.evaluatingDo = true
	} else if sm.state == o && r == '(' && sm.evaluatingDo == true {
		sm.state = enter
	} else if sm.state == enter && r == ')' && sm.evaluatingDo == true {
		sm.doState = true
		sm.logger.Printf("do state found\n")
		// don't part branching
	} else if sm.state == o && r == 'n' {
		sm.evaluatingDo = false
		sm.state = n
	} else if sm.state == n && r == '\'' {
		sm.state = tick
	} else if sm.state == tick && r == 't' {
		sm.evaluatingDont = true
		sm.state = t
	} else if sm.state == t && r == '(' && sm.evaluatingDont == true {
		sm.state = enter
	} else if sm.state == enter && r == ')' && sm.evaluatingDont == true {
		sm.doState = false
		sm.logger.Printf("don't state found\n")
	} else {
		if sm.state != nothing {
			sm.Reinit()
			sm.readRune(r)
		}
	}
}

func (sm *StateMachine) IsEnd() bool {
	return sm.state == end && sm.doState == true
}

func (sm *StateMachine) GetValues() (int, int) {
	return sm.first, sm.second
}
