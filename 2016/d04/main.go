package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Room struct {
	name     string
	sector   int
	checksum string
}

func parse(lines []string) []Room {
	rooms := make([]Room, len(lines))
	re := regexp.MustCompile(`(.+)-(\d+)\[(\w+)\]`)
	for i, line := range lines {
		match := re.FindAllStringSubmatch(line, -1)
		sector, _ := strconv.Atoi(match[0][2])
		rooms[i] = Room{
			name:     match[0][1],
			sector:   sector,
			checksum: match[0][3],
		}
	}
	return rooms
}

func isRealRoom(room Room) bool {
	freq := map[rune]int{}
	for _, l := range room.name {
		if l != '-' {
			freq[l] += 1
		}
	}
	letters := []rune{}
	for l := range freq {
		letters = append(letters, l)
	}
	slices.SortFunc(letters, func(a, b rune) int {
		diff := freq[b] - freq[a]
		if diff == 0 {
			return int(a) - int(b)
		}
		return diff
	})
	for _, l := range room.checksum {
		idx := slices.Index(letters, l)
		if idx == -1 || idx > 5 {
			return false
		}
	}
	return true
}

func PartOne(lines []string) (count int) {
	rooms := parse(lines)
	for _, room := range rooms {
		if isRealRoom(room) {
			count += room.sector
		}
	}
	return count
}

func (room *Room) decrypt() string {
	first := int('a')
	result := ""
	for _, l := range room.name {
		if l == '-' {
			result += " "
			continue
		}
		next := first + (int(l)+room.sector-first)%26
		result += string(rune(next))
	}
	return result
}

func PartTwo(lines []string) (count int) {
	rooms := parse(lines)
	for _, room := range rooms {
		if room.decrypt() == "northpole object storage" {
			return room.sector
		}
	}
	return -1
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
