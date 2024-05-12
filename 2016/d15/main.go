package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"time"
)

type Disk struct {
	start int
	total int
}

func (d Disk) getPosition(t int) int {
	return ((d.start + t) % d.total)
}

func parse(lines []string) []Disk {
	re := regexp.MustCompile(`Disc #(\d+) has (\d+) positions;.*at position (\d)+.`)
	disks := make([]Disk, len(lines))
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, -1)
		nums := utils.ToInts(m[0][1:])
		disks[nums[0]-1] = Disk{start: nums[2], total: nums[1]}
	}
	return disks
}

// Ideally, the Chinese remainder theorem should be used, but brute force is good enough here anyway.
// Also, in case of brute force, loops could be optimized by sorting the disks by their lengths
// and skipping any positions beyond the position of the disk with the smallest number of elements.
func calc(lines []string, isPartTwo bool) int {
	disks := parse(lines)
	if isPartTwo {
		disks = append(disks, Disk{start: 0, total: 11})
	}
	for t := 0; ; t++ {
		first := disks[0].getPosition(t + 1)
		allSame := true
		for j := 1; j < len(disks); j++ {
			if first != disks[j].getPosition(t+1+j) {
				allSame = false
				break
			}
		}
		if allSame {
			return t
		}
	}
}

func main() {
	lines := []string{
		"Disc #1 has 5 positions; at time=0, it is at position 4.",
		"Disc #2 has 2 positions; at time=0, it is at position 1.",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	ts := time.Now()
	fmt.Println("PartOne: ", calc(lines, false), time.Since(ts))
	ts = time.Now()
	fmt.Println("PartTwo: ", calc(lines, true), time.Since(ts))
}
