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

	res := values[0]
	for _, i := range values {
		if res > i {
			res = i
		}
	}
	return res
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
}
