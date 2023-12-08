package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
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

func parseDirs(l string) (dir []int) {
	for _, d := range l {
		if d == 'L' {
			dir = append(dir, 0)
			continue
		}
		dir = append(dir, 1)
	}
	return dir
}

func parseSteps(lines []string) map[string][]string {
	steps := make(map[string][]string)
	re := regexp.MustCompile(`[A-Z|0-9]+`)
	for _, l := range lines {
		r := re.FindAllString(l, -1)
		steps[r[0]] = r[1:]
	}
	return steps
}

func PartOne(lines []string) int {
	dir := parseDirs(lines[0])
	steps := parseSteps(lines[2:])

	count := 0
	for step := "AAA"; step != "ZZZ"; count++ {
		d := dir[count%len(dir)]
		step = steps[step][d]
	}
	return count
}

func findLoop(start string, dirs []int, steps map[string][]string) (count int) {
	zz := []string{}
	var include = func(step string) bool {
		for _, s := range zz {
			if s == step {
				return true
			}
		}
		return false
	}
	step := start
	for i := 0; ; i++ {
		count++
		d := dirs[i%len(dirs)]
		step = steps[step][d]
		if strings.HasSuffix(step, "Z") {
			if len(zz) == 0 {
				count = 0
			}
			if include(step) {
				break
			}
			zz = append(zz, step)
			continue
		}
	}
	return count
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums ...int) int {
	a, b, rest := nums[0], nums[1], nums[2:]
	result := a * b / gcd(a, b)

	for _, c := range rest {
		result = lcm(result, c)
	}

	return result
}

func PartTwo(lines []string) int {
	dirs := parseDirs(lines[0])
	steps := parseSteps(lines[2:])

	starts := []string{}
	for key := range steps {
		if strings.HasSuffix(key, "A") {
			starts = append(starts, key)
		}
	}

	loops := []int{}
	for _, start := range starts {
		loops = append(loops, findLoop(start, dirs, steps))
	}
	return lcm(loops...)
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))
}
