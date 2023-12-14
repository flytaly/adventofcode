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

func PartOne(lines []string) {
	result := 0
	for i := 0; i < len(lines[0]); i++ {
		load := len(lines)
		for j, line := range lines {
			if line[i] == 'O' {
				result += load
				load -= 1
			}
			if line[i] == '#' {
				load = len(lines) - j - 1
			}

		}
	}

	fmt.Println("Part 1:", result)
}

func rollNorth(grid *[][]rune) {
	for i, g := 0, *grid; i < len(g[0]); i++ {
		for j, empty := 0, 0; j < len(g); j++ {
			cell := g[j][i]
			if cell == 'O' {
				g[empty][i], g[j][i] = cell, g[empty][i]
				empty++
			}
			if cell == '#' {
				empty = j + 1
			}
		}
	}
}

func rollSouth(grid *[][]rune) {
	for i, g := 0, *grid; i < len(g[0]); i++ {
		for j, empty := len(g)-1, len(g)-1; j >= 0; j-- {
			cell := g[j][i]
			if cell == 'O' {
				g[empty][i], g[j][i] = cell, g[empty][i]
				empty--
			}
			if cell == '#' {
				empty = j - 1
			}
		}
	}
}

func rollWest(grid *[][]rune) {
	for j, g := 0, *grid; j < len(g); j++ {
		for i, empty := 0, 0; i < len(g[0]); i++ {
			cell := (g)[j][i]
			if cell == 'O' {
				g[j][empty], g[j][i] = cell, g[j][empty]
				empty++
			}
			if cell == '#' {
				empty = i + 1
			}
		}
	}
}

func rollEast(grid *[][]rune) {
	last := len((*grid)[0]) - 1
	for j, g := 0, *grid; j < len(g); j++ {
		for i, empty := last, last; i >= 0; i-- {
			cell := (g)[j][i]
			if cell == 'O' {
				g[j][empty], g[j][i] = cell, g[j][empty]
				empty--
			}
			if cell == '#' {
				empty = i - 1
			}
		}
	}

}

func countLoad(lines []string) (result int) {
	for i := 0; i < len(lines[0]); i++ {
		for j, line := range lines {
			if line[i] == 'O' {
				result += len(lines) - j
			}
		}
	}

	return result
}

func rollCycle(rocks *[][]rune) {
	rollNorth(rocks)
	rollWest(rocks)
	rollSouth(rocks)
	rollEast(rocks)
}

func gridId(grid [][]rune) string {
	var sb strings.Builder
	for i, row := range grid {
		for _, v := range row {
			sb.WriteRune(v)
		}
		if i != len(grid)-1 {
			sb.WriteString("|")
		}
	}
	return sb.String()
}

const CYCLES_NUM = 1000000000

func PartTwo(lines []string) {
	rocks := [][]rune{}

	rolls := map[string]int{}

	for _, l := range lines {
		rocks = append(rocks, []rune(l))
	}

	start, period := 0, CYCLES_NUM+1
	rolls[gridId(rocks)] = 0
	for i := 1; i <= CYCLES_NUM; i++ {
		rollCycle(&rocks)
		id := gridId(rocks)
		if prevIndex, ok := rolls[id]; ok {
			start, period = i, i-prevIndex
			break
		}
		rolls[id] = i
	}
	// fmt.Println("Period", period, "start", start)

	offset := (CYCLES_NUM - start) % period
	for id, idx := range rolls {
		if idx != start-period+offset {
			continue
		}
		grid := strings.Split(id, "|")
		fmt.Println("Part 2:", countLoad(grid))
	}
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
	PartTwo(lines)
}
