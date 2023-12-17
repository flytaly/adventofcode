package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
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

func parse(lines []string) map[complex64]int {
	res := map[complex64]int{}
	for i, l := range lines {
		split := strings.Split(l, "")
		for j, v := range split {
			n, _ := strconv.Atoi(v)
			res[complex(float32(i), float32(j))] = n
		}
	}
	return res
}

type Block struct {
	cost int
	pos  complex64
	dir  complex64
	c    int // consecutive
}

func (b Block) String() string {
	return fmt.Sprintf("%v,%v,%d", b.pos, b.dir, b.c)
}

func findMin(grid map[complex64]int, end complex64) int {
	visited := map[string]bool{}
	queue := []Block{{0, 0, 1, 0}, {0, 0, 1i, 0}}

	for len(queue) > 0 {
		id := 0
		// TODO: use PriorityQueue
		for i, val := 1, queue[0].cost; i < len(queue); i++ {
			if queue[i].cost < val {
				val = queue[i].cost
				id = i
			}
		}

		b := queue[id]
		queue = append(queue[:id], queue[id+1:]...)
		if visited[b.String()] {
			continue
		}
		visited[b.String()] = true

		if b.pos == end {
			return b.cost
		}

		dir := b.dir
		if pos := b.pos + dir; b.c < 3 && grid[pos] != 0 {
			queue = append(queue, Block{b.cost + grid[pos], pos, dir, b.c + 1})
		}

		dir = dir * 1i // rotate 90deg
		if pos := b.pos + dir; grid[pos] != 0 {
			queue = append(queue, Block{b.cost + grid[pos], pos, dir, 1})
		}

		dir = dir * (-1) // rotate 180deg -> -90deg from initial
		if pos := b.pos + dir; grid[pos] != 0 {
			queue = append(queue, Block{b.cost + grid[pos], pos, dir, 1})
		}
	}

	return 0
}

func PartOne(lines []string) {
	grid := parse(lines)

	end := complex(float32(len(lines)-1), float32(len(lines[0])-1))

	fmt.Println("Part 1:", findMin(grid, end))
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
