package main

import (
	"aoc/utils"
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

func getCombinations(list []string) iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		for i := len(list); i > 1; i-- {
			for _, indexes := range combin.Combinations(len(list), i) {
				elems := make([]string, 0, i)
				for _, idx := range indexes {
					elems = append(elems, list[idx])
				}
				if !yield(elems) {
					return
				}
			}
		}
	}
}

type LinkMap map[string]map[string]bool

func findLargest(linkMap LinkMap, node string, minLen int) (seq []string) {
	list := make([]string, 0, len(linkMap[node]))
	for link := range linkMap[node] {
		list = append(list, link)
	}
	for rest := range getCombinations(list) {
		if len(rest)+1 <= minLen {
			break
		}
		if areAllConnected(linkMap, rest) {
			seq = rest
			break
		}
	}

	return append([]string{node}, seq...)
}

func areAllConnected(linkMap LinkMap, links []string) bool {
	for i, link1 := range links[:len(links)-1] {
		for _, link2 := range links[i+1:] {
			if !linkMap[link1][link2] {
				return false
			}
		}
	}
	return true
}

func PartTwo(lines []string) string {
	linkMap := LinkMap{}
	for _, l := range lines {
		linkMap[strings.Split(l, "-")[0]] = map[string]bool{}
	}
	for _, l := range lines {
		ids := strings.Split(l, "-")
		linkMap[ids[0]][ids[1]] = true
		linkMap[ids[1]][ids[0]] = true
	}

	result := []string{}

	for node := range linkMap {
		seq := findLargest(linkMap, node, len(result))
		if len(seq) > len(result) {
			result = seq
		}
	}

	slices.Sort(result)
	return strings.Join(result, ",")
}

func main() {
	inputFile := "example.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := utils.ReadLines(inputFile)

	ts := time.Now()
	fmt.Printf("Part 1: %d [%v] \n", PartOne(lines), time.Since(ts))
	ts = time.Now()
	fmt.Printf("Part 2: %s [%v] \n", PartTwo(lines), time.Since(ts))
}
