package main

import (
	"aoc/utils"
	"cmp"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type WireMap map[string]int

func (w WireMap) getPair(a, b string) (int, int, bool) {
	v1, ok1 := w[a]
	v2, ok2 := w[b]
	if ok1 && ok2 {
		return v1, v2, true
	}
	return 0, 0, false
}

type Gate [4]string

func trigger(a, b int, op string) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	default:
		panic("unknown op " + op)
	}
}

func PartOne(wires WireMap, gates []Gate) int {
	zMap := map[string]int{}
	for len(gates) > 0 {
		queue := []Gate{}
		for _, gate := range gates {
			if a, b, ok := wires.getPair(gate[0], gate[2]); ok {
				wires[gate[3]] = trigger(a, b, gate[1])
				if strings.HasPrefix(gate[3], "z") {
					zMap[gate[3]] = wires[gate[3]]
				}
				continue
			}
			queue = append(queue, gate)
		}
		gates = queue
	}

	keys := slices.Collect(maps.Keys(zMap))
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(b, a)
	})
	result := ""
	for _, key := range keys {
		result += strconv.Itoa(zMap[key])
	}
	num, _ := strconv.ParseInt(result, 2, 64)

	return int(num)
}

func parse(lines []string) (wires WireMap, gates []Gate) {
	div := slices.IndexFunc(lines, func(s string) bool {
		return s == ""
	})
	wires = WireMap{}
	for _, w := range lines[0:div] {
		split := strings.Split(w, ": ")
		n, _ := strconv.Atoi(split[1])
		wires[split[0]] = n
	}

	gates = []Gate{}
	re := regexp.MustCompile(`[\s->]+`)
	for _, g := range lines[div+1:] {
		split := re.Split(g, -1)
		gates = append(gates, [4]string(split))
	}
	return wires, gates
}

func main() {
	inputFile := "example.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	wires, gates := parse(utils.ReadLines(inputFile))

	ts := time.Now()
	fmt.Printf("Part 1: %d [%v] \n", PartOne(wires, gates), time.Since(ts))
}
