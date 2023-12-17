package main

import (
	"container/heap"
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

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Block

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Block))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
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

func (b Block) id() string {
	return fmt.Sprintf("%v,%v,%d", b.pos, b.dir, b.c)
}

func findMin(grid map[complex64]int, end complex64, ultra bool) int {
	visited := map[string]bool{}
	queue := []Block{{0, 0, 1, 0}, {0, 0, 1i, 0}}

	pq := make(PriorityQueue, len(queue))
	for i, priority := range queue {
		pq[i] = &priority
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		b := heap.Pop(&pq).(*Block)
		if visited[b.id()] {
			continue
		}
		visited[b.id()] = true

		canStop, maxC := true, 3
		if ultra {
			maxC = 10
			canStop = b.c >= 4
		}

		if b.pos == end && canStop {
			return b.cost
		}

		dir := b.dir
		if pos := b.pos + dir; grid[pos] != 0 && b.c < maxC {
			heap.Push(&pq, &Block{b.cost + grid[pos], pos, dir, b.c + 1})
		}

		dir = dir * 1i // rotate 90deg
		if pos := b.pos + dir; grid[pos] != 0 && canStop {
			heap.Push(&pq, &Block{b.cost + grid[pos], pos, dir, 1})
		}

		dir = dir * (-1) // rotate 180deg -> -90deg from initial
		if pos := b.pos + dir; grid[pos] != 0 && canStop {
			heap.Push(&pq, &Block{b.cost + grid[pos], pos, dir, 1})
		}
	}

	return 0
}

func PartOne(lines []string) {
	grid := parse(lines)

	end := complex(float32(len(lines)-1), float32(len(lines[0])-1))

	fmt.Println("Part 1:", findMin(grid, end, false))
}

func PartTwo(lines []string) {
	grid := parse(lines)

	end := complex(float32(len(lines)-1), float32(len(lines[0])-1))

	fmt.Println("Part 2:", findMin(grid, end, true))
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
	PartTwo(lines) //1048 low
}
