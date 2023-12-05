package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func readLines() []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func toInt(s []string) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

type ConvMap struct {
	dest   int
	source int
	size   int
}

func parse(lines []string) ([]int, [][]ConvMap) {
	re := regexp.MustCompile(`\d+`)
	maps := [][]ConvMap{}
	seeds := toInt(re.FindAllString(lines[0], -1))
	var m = []ConvMap{}
	for i := 3; i < len(lines); i++ {
		l := lines[i]
		if l == "" {
			continue
		}
		if l[len(l)-1] == ':' {
			maps = append(maps, m)
			m = []ConvMap{}
			continue
		}
		nums := toInt(re.FindAllString(lines[i], -1))
		m = append(m, ConvMap{dest: nums[0], source: nums[1], size: nums[2]})
	}
	maps = append(maps, m)
	return seeds, maps
}

func convert(maps []ConvMap, v int) int {
	for _, m := range maps {
		if v >= m.source && v < m.source+m.size {
			return m.dest + v - m.source
		}
	}
	return v
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func minSlice(ii ...int) int {
	res := MaxInt
	for _, i := range ii {
		if res > i {
			res = i
		}
	}
	return res
}

func PartOne(lines []string) int {
	values, convMaps := parse(lines)
	for _, maps := range convMaps {
		for i, v := range values {
			values[i] = convert(maps, v)
		}
	}
	return minSlice(values...)
}

type Span = [2]int

// sp should be subinterval of m
func mapSubrange(sp Span, m ConvMap) Span {
	spLen := sp[1] - sp[0]
	shift := sp[0] - m.source
	return Span{m.dest + shift, m.dest + shift + spLen}
}

func mapSpan(span Span, mp ConvMap) (mapped []Span, unmapped []Span) {
	m := Span{mp.source, mp.source + mp.size}

	if (span[0] >= m[1]) || (span[1] < m[0]) { // not overlaping
		return nil, []Span{span}
	}

	if span[0] < m[0] { // overlapping from the left
		unmapped = []Span{{span[0], m[0]}}
		if span[1] > m[1] {
			// Span{m[0], m[1]},
			mapped = []Span{mapSubrange(Span{m[0], m[1]}, mp)}
			unmapped = append(unmapped, Span{m[1], span[1]})
			return mapped, unmapped
		}
		mapped = []Span{mapSubrange(Span{m[0], span[1]}, mp)}
		return mapped, unmapped
	}
	// inside
	if span[1] <= m[1] {
		return []Span{mapSubrange(span, mp)}, nil
	}
	// overlapping from the inside to the right
	mapped = []Span{mapSubrange(Span{span[0], m[1]}, mp)}
	return mapped, []Span{{m[1], span[1]}}
}

func PartTwo(lines []string) int {
	seeds, convMaps := parse(lines)
	spans := []Span{}
	for i := 0; i < len(seeds); i += 2 {
		spans = append(spans, Span{seeds[i], seeds[i] + seeds[i+1]})
	}

	for _, maps := range convMaps {
		successful := []Span{}
		for _, m := range maps {
			skipped := []Span{}
			for _, sp := range spans {
				mapped, unmapped := mapSpan(sp, m)
				skipped = append(skipped, unmapped...)
				successful = append(successful, mapped...)
			}
			spans = skipped
		}
		spans = append(spans, successful...)
	}

	res := MaxInt
	for _, sp := range spans {
		if res > sp[0] {
			res = sp[0]
		}
	}
	return res
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))
}
