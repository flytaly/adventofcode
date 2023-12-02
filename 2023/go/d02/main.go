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

type Cubes = map[string]int

func parseLine(line string) (int, []Cubes) {
	var rCube = regexp.MustCompile(`(\d+)\s*(red|green|blue)`)
	var rNum = regexp.MustCompile(`\d+`)
	split := strings.Split(line, ":")
	gameNum, _ := strconv.Atoi(rNum.FindString(split[0]))
	setstr := strings.Split(split[1], ";")

	var results = []Cubes{}
	for _, set := range setstr {
		s := Cubes{}
		for _, cube := range strings.Split(set, ",") {
			match := rCube.FindStringSubmatch(cube)
			n, _ := strconv.Atoi(match[1])
			s[match[2]] = n
		}
		results = append(results, s)
	}

	return gameNum, results
}

func PartOne(lines []string) int {
	result := 0
	limits := Cubes{"red": 12, "green": 13, "blue": 14}

	var checkIfPossible = func(game []Cubes) bool {
		for _, set := range game {
			for cube, used := range set {
				if limits[cube] < used {
					return false
				}
			}
		}
		return true
	}

	for _, l := range lines {
		gameNum, sets := parseLine(l)
		if checkIfPossible(sets) {
			result += gameNum
			continue
		}
		fmt.Println(gameNum, "is not possible")
	}
	return result
}

func PartTwo(lines []string) int {
	result := 0

	var findMin = func(game []Cubes) Cubes {
		minVals := Cubes{}
		for _, set := range game {
			for cube, used := range set {
				if minVals[cube] < used {
					minVals[cube] = used
				}
			}
		}
		return minVals
	}

	for _, l := range lines {
		gameNum, sets := parseLine(l)
		mins := findMin(sets)
		pow := mins["red"] * mins["green"] * mins["blue"]
		fmt.Printf("%d %v => %d\n", gameNum, mins, pow)
		result += pow
	}

	return result
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))
}
