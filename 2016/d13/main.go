package main

import (
	"fmt"
	"os"
	"strconv"
)

type Coords struct {
	x int
	y int
}

type Cell struct {
	dist    int
	isWall  bool
	visited bool
}

func printGrid(grid map[Coords]Cell) {
	maxX, maxY := 0, 0
	for k := range grid {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := Coords{x, y}
			if x == 1 && y == 1 {
				fmt.Print("O")
			} else if grid[c].isWall {
				fmt.Print("#")
			} else if !grid[c].visited {
				fmt.Print("▓")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func isItAWall(x, y, n int) bool {
	v := x*x + 3*x + 2*x*y + y + y*y + n
	ones := 0
	for _, v := range strconv.FormatInt(int64(v), 2) {
		if v == '1' {
			ones++
		}
	}
	return ones%2 == 1
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func selectNextMin(grid map[Coords]Cell) Coords {
	minValue := MaxInt
	var minCoords Coords
	for coords, cell := range grid {
		if cell.isWall || cell.visited {
			continue
		}
		if cell.dist < minValue {
			minValue = cell.dist
			minCoords = coords
		}
	}
	return minCoords
}

// Djikstra-like
func PartOne(n int, xTarget, yTarget int) int {
	maze := map[Coords]Cell{{1, 1}: {0, false, false}}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for {
		current := selectNextMin(maze) // should be min heap instead
		if current.x == xTarget && current.y == yTarget {
			printGrid(maze)
			fmt.Println()
			return maze[current].dist
		}
		for _, dir := range dirs {
			coords := Coords{current.x + dir[0], current.y + dir[1]}
			if coords.x < 0 || coords.y < 0 || maze[coords].isWall || maze[coords].visited {
				continue
			}
			dist := maze[current].dist
			if nextCell, ok := maze[coords]; ok {
				maze[coords] = Cell{dist: min(dist+1, nextCell.dist), isWall: false}
				continue
			}
			maze[coords] = Cell{dist: dist + 1, isWall: isItAWall(coords.x, coords.y, n), visited: false}
		}
		cell := maze[current]
		maze[current] = Cell{dist: cell.dist, isWall: cell.isWall, visited: true}

	}
}

// bfs
func PartTwo(n int, steps int) int {
	visited := map[Coords]bool{{1, 1}: true}
	count := 1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	frontier := []Coords{{1, 1}}
	for i := 0; i < steps; i++ {
		nextFront := make([]Coords, 0)
		for _, current := range frontier {
			for _, dir := range dirs {
				coords := Coords{current.x + dir[0], current.y + dir[1]}
				if coords.x < 0 || coords.y < 0 || visited[coords] {
					continue
				}
				visited[coords] = true
				if !isItAWall(coords.x, coords.y, n) {
					count++
					nextFront = append(nextFront, coords)
				}
			}
		}
		frontier = nextFront
	}
	return count
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Test:")
		fmt.Println("PartOne: ", PartOne(10, 7, 4))
		fmt.Println("PartTwo: ", PartTwo(10, 6))
		return
	}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("PartOne: ", PartOne(input, 31, 39))
	fmt.Println("PartTwo: ", PartTwo(input, 50))
}
