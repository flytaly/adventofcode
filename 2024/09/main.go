package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func PartOne(disk string) {
	var sum, pos, idLeft int
	idRight := len(disk) / 2
	idPool := int(disk[idRight*2] - '0')

	for i := 0; idLeft < idRight; i++ {
		size := int(disk[i] - '0')
		start := pos

		// taken space
		if i%2 == 0 {
			// sum of arithmetic series
			end := pos + size - 1
			sum += (idLeft * (start + end) * size) / 2
			pos = end + 1
			idLeft++
			continue
		}
		// fill free space
		for j := start; j < start+size; j++ {
			if idPool == 0 {
				idRight--
				idPool = int(disk[idRight*2] - '0')
			}
			sum += pos * idRight
			pos, idPool = pos+1, idPool-1
		}
	}

	for _ = range idPool {
		sum += idRight * pos
		pos++
	}

	fmt.Println("Part  1:", sum)
}


func PartTwo(disk string) {

}

func main() {
	lines := []string{"2333133121414131402"}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines[0])
}
