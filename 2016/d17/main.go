package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var room = []string{
	"######",
	"#    #",
	"#    #",
	"#    #",
	"#   XV",
	"####VV",
}

func getCell(pos []int) byte {
	return room[pos[0]][pos[1]]
}

var openChars = []byte{'b', 'c', 'd', 'e', 'f'}

func isOpen(ch byte) bool {
	return slices.Contains(openChars, ch)
}

type Position struct {
	pos  []int
	code []byte
}

func findPath(passcode []byte, isPartTwo bool) string {
	dirs := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
	dirChars := []byte{'U', 'D', 'L', 'R'}

	queue := []Position{{pos: []int{1, 1}, code: passcode}}
	maxLen := 0
	for {
		if len(queue) == 0 {
			if isPartTwo {
				return strconv.Itoa(maxLen)
			}
			fmt.Println("no	more moves")
			return ""
		}
		nextQueue := []Position{}
		minPath := ""
		for _, step := range queue {
			// fmt.Println(step.pos, string(step.code))
			hash := fmt.Sprintf("%x", md5.Sum(step.code))
			pos := step.pos
			for i, dir := range dirs {
				if !isOpen(hash[i]) {
					continue
				}
				nextPos := []int{pos[0] + dir[0], pos[1] + dir[1]}
				switch getCell(nextPos) {
				case 'X':
					if isPartTwo {
						if length := len(step.code) - len(passcode) + 1; length > maxLen {
							maxLen = length
						}
						continue
					}
				case 'V':
					path := string(step.code[len(passcode):])
					if minPath == "" || len(path) <= len(minPath) {
						minPath = path
					}
					continue
				case '#':
					continue
				}
				nextQueue = append(nextQueue, Position{
					pos:  nextPos,
					code: append(slices.Clone(step.code), dirChars[i]),
				})
			}
		}
		if minPath != "" {
			return minPath
		}
		queue = nextQueue
	}
}

func main() {
	input := "ihgpwlah"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	fmt.Println("PartOne: ", findPath([]byte(input), false))
	fmt.Println("PartTwo: ", findPath([]byte(input), true))
}
