package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	x, y, size, used, avail, usePerc int
}

func (n Node) key() string {
	return fmt.Sprintf("%d_%d", n.x, n.y)
}

func parse(lines []string) []Node {
	nodeList := []Node{}
	re := regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)\%`)
	for _, line := range lines[2:] {
		n := utils.ToInts(re.FindStringSubmatch(line))
		node := Node{n[1], n[2], n[3], n[4], n[5], n[6]}
		nodeList = append(nodeList, node)
	}
	return nodeList
}

func P1(lines []string) int {
	nodeList := parse(lines)

	count := 0
	for i := 0; i < len(nodeList); i++ {
		for j := 0; j < len(nodeList); j++ {
			a, b := nodeList[i], nodeList[j]
			if a.key() != b.key() && a.used > 0 && a.used <= b.avail {
				count++
			}
		}
	}

	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", P1(lines))
}
