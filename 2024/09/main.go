package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"slices"
	"time"
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

	fmt.Println("Part 1:", sum)
}

// returns the gap position and its size and removes the gap from the map
func minValidGap(gaps map[int][]int, minSize int, maxPos int) (gapPos int, gapSize int, ok bool) {
	gapPos = math.MaxInt

	for size, positions := range gaps {
		if size < minSize || len(positions) == 0 {
			continue
		}
		if p := positions[len(positions)-1]; gapPos > p && p < maxPos {
			gapPos = p
			gapSize = size
		}
	}

	if gapPos < math.MaxInt {
		gg := gaps[gapSize]
		gaps[gapSize] = gg[:len(gg)-1]
		return gapPos, gapSize, true
	}

	return 0, 0, false
}

func PartTwo(diskMap string) {
	disk := []int{}
	gaps := map[int][]int{} // size to block's position map
	//
	id := 0

	for i, value := range diskMap {
		num := int(value - '0')
		blockId := id
		id += 1 - i%2
		if i%2 != 0 {
			blockId = 0
			gaps[num] = append(gaps[num], len(disk))
		}
		disk = append(disk, slices.Repeat([]int{blockId}, num)...)
	}
	for i := range gaps {
		slices.Reverse(gaps[i])
	}

	var startPos int
	endPos := len(disk) - 1
	for i := len(diskMap) - 1; i > 2; i-- {
		size := int(diskMap[i] - '0')
		if i%2 != 0 {
			endPos = endPos - size
			continue
		}

		startPos = endPos - size + 1

		file := disk[startPos : endPos+1]
		gapPos, gapSize, ok := minValidGap(gaps, size, startPos)
		if !ok {
			endPos = startPos - 1
			continue
		}
		if gapLeft := gapSize - size; gapLeft > 0 {
			gaps[gapLeft] = append(gaps[gapLeft], gapPos+size)
			slices.SortFunc(gaps[gapLeft], func(a, b int) int {
				return b - a
			})
		}

		// insert the file into the gap.
		disk = slices.Replace(disk, gapPos, gapPos+size, file...)

		// clear the file
		for j := startPos; j <= endPos; j++ {
			disk[j] = 0
		}

		endPos = startPos - 1
	}

	sum := 0
	for i, value := range disk {
		sum += i * value
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	lines := []string{"2333133121414131402"}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines[0])
	ts := time.Now()
	PartTwo(lines[0])
	fmt.Println(time.Since(ts))
}
