package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func readLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

type Module struct {
	tp   string
	dest []string
	on   bool
	memo map[string]string // conjunction module memory from type -> pulse type
}

func (m Module) memoOnlyHigh() bool {
	for _, v := range m.memo {
		if v != "h" {
			return false
		}
	}
	return true
}

type Pulse struct {
	from string
	tp   string // l or h
	to   string
}

func parse(input []string) map[string]Module {
	m := make(map[string]Module)
	for _, v := range input {
		s := strings.Split(v, " -> ")
		id, dest := s[0], strings.Split(s[1], ", ")
		if id == "broadcaster" {
			m[id] = Module{tp: "broadcaster", dest: dest}
			continue
		}
		m[id[1:]] = Module{tp: id[:1], dest: dest, on: false, memo: make(map[string]string)}
	}
	return m
}

func PartOne(lines []string) {
	modules := parse(lines)

	sent := [2]int{0, 0}
	count := func(pulseType string) {
		if pulseType == "h" {
			sent[1]++
			return
		}
		sent[0]++
	}

	for name, m := range modules {
		for _, dest := range m.dest {
			if modules[dest].tp == "&" {
				modules[dest].memo[name] = "l"
			}
		}
	}

	for i := 0; i < 1000; i++ {
		Q := []Pulse{{tp: "l", to: "broadcaster"}}
		count("l")
		for len(Q) > 0 {
			p := Q[0]
			Q = Q[1:]
			m := modules[p.to]
			pulseType := p.tp
			destType, dest := m.tp, m.dest

			if destType == "broadcaster" {
				for _, v := range dest {
					Q = append(Q, Pulse{tp: pulseType, to: v, from: "broadcaster"})
					count(pulseType)
				}
			}

			if destType == "%" { // flipflop
				if p.tp == "h" {
					continue
				}
				nextType := "h"
				if m.on {
					nextType = "l"
				}
				m.on = !m.on
				modules[p.to] = m
				for _, v := range dest {
					Q = append(Q, Pulse{tp: nextType, to: v, from: p.to})
					count(nextType)
				}
			}

			if destType == "&" { // conjunction
				m.memo[p.from] = pulseType
				nextType := "h"
				if m.memoOnlyHigh() {
					nextType = "l"
				}
				for _, v := range dest {
					Q = append(Q, Pulse{tp: nextType, to: v, from: p.to})
					count(nextType)
				}
			}

		}
	}

	fmt.Println("Part 1:", sent, "=>", sent[0]*sent[1])
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
