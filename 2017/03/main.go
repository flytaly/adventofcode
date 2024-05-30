package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func manhattan(a, b []int) int {
	return abs(a[0], b[0]) + abs(a[1], b[1])
}

func P1(input int) (steps int) {
	sideSize := int(math.Ceil(math.Sqrt(float64(input))))
	sideSize += (1 - sideSize%2) // make sure it's odd
	br := sideSize * sideSize    // next bottom-right corner value
	n := (sideSize - 1) / 2
	coords := []int{n, n}          // br coords {1,1}, {3,3}, {5,5}...
	for i := br; i != input; i-- { // move backward from bottom-right
		if i > br-(sideSize-1) { //  left
			coords[0]--
		} else if i > br-(sideSize-1)*2 { //  up
			coords[1]--
		} else if i > br-(sideSize-1)*3 { // right
			coords[0]++
		} else if i > br-(sideSize-1)*4 { // down
			coords[1]++
		}
	}
	return manhattan(coords, []int{0, 0})
}

type Dir = int

const (
	UP Dir = iota
	LEFT
	DOWN
	RIGHT
)

type C struct {
	x, y int
}

func (cur *C) move(dir Dir) {
	switch dir {
	case UP: // up
		cur.y -= 1
	case LEFT: // left
		cur.x -= 1
	case DOWN: // down
		cur.y += 1
	case RIGHT: // right
		cur.x += 1
	}

}

func neighborSum(spiral map[C]int, square C) (sum int) {
	for _, v := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		neighbor := C{square.x + v[0], square.y + v[1]}
		sum += spiral[neighbor]
	}
	return sum
}

func P2(input int) (res int) {
	spiral := map[C]int{C{0, 0}: 1}

	cur := C{0, 0}
	for side := 1; ; side++ {
		cur.x += 1 // go to the next layer
		spiral[cur] = neighborSum(spiral, cur)
		for _, dir := range []Dir{UP, LEFT, DOWN, RIGHT} {
			for i := 0; i < side; i++ {
				cur.move(dir)
				if spiral[cur] = neighborSum(spiral, cur); spiral[cur] > input {
					return spiral[cur]
				}
			}
			if dir == UP {
				side++
			}
		}
	}
}

func main() {
	input := 23
	if len(os.Args) > 1 {
		input, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println("input:", input)
	fmt.Println("Part 1 => ", P1(input))
	fmt.Println("Part 2 => ", P2(input))
}
