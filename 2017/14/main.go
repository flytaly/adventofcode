package main

import (
	"aoc/2017/14/knothash"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func P1(input string) (count int) {
	for i := 0; i < 128; i++ {
		for _, ch := range knothash.Hash(input + "-" + strconv.Itoa(i)) {
			hex, _ := strconv.ParseUint(string(ch), 16, 64)
			count += bits.OnesCount64(hex)
		}
	}
	return count
}

type Cell struct {
	row, col int
}

func markRegionDFS(grid map[Cell]bool, visited map[Cell]bool, cell Cell) {
	visited[cell] = true
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		adj := Cell{cell.row + dir[0], cell.col + dir[1]}
		if grid[adj] && !visited[adj] {
			markRegionDFS(grid, visited, adj)
		}
	}
}

func P2(input string) (count int) {
	grid := map[Cell]bool{}
	for row := 0; row < 128; row++ {
		for j, ch := range knothash.Hash(input + "-" + strconv.Itoa(row)) {
			hex, _ := strconv.ParseUint(string(ch), 16, 64)
			for offset, ch := range fmt.Sprintf("%04b", hex) {
				grid[Cell{row, j*4 + offset}] = ch == '1'
			}
		}
	}

	visited := map[Cell]bool{}

	for cell, isUsed := range grid {
		if !isUsed || visited[cell] {
			continue
		}
		count++
		markRegionDFS(grid, visited, cell)
	}

	return count
}

func main() {
	input := "flqrgnkx"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	fmt.Println("Part 1 =>", P1(input))
	fmt.Println("Part 2 =>", P2(input))
}
