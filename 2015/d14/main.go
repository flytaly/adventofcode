package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func parser(lines []string) []Reindeer {
	re := regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	res := []Reindeer{}
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, 1)
		nums := utils.ToInts(m[0][1:])
		res = append(res, Reindeer{speed: nums[0], time: nums[1], rest: nums[2]})

	}
	return res
}

type Reindeer struct {
	speed, time, rest int
}

func (r Reindeer) travelled(seconds int) int {
	burstTime := r.time + r.rest
	bursts := seconds / burstTime
	timeLeft := seconds - bursts*burstTime

	return (r.speed*r.time)*bursts + min(timeLeft, r.time)*r.speed
}

func PartOne(lines []string, totalTime int) (distance int) {
	reindeers := parser(lines)

	for _, r := range reindeers {
		distance = max(distance, r.travelled(totalTime))
	}

	return distance
}

func PartTwo(lines []string, totalTime int) (bestScore int) {
	reindeers := parser(lines)
	points := make([]int, len(reindeers))

	for i := 1; i <= totalTime; i++ {
		dists := utils.Map2(reindeers, func(r Reindeer) int {
			return r.travelled(i)
		})
		bestDist := slices.Max(dists)
		for index, d := range dists {
			if d == bestDist {
				points[index]++
			}
		}
	}

	for _, v := range points {
		if v > bestScore {
			bestScore = v
		}
	}

	return bestScore
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	if len(lines) == 0 {
		return
	}
	fmt.Println("PartOne: ", PartOne(lines, 2503))
	fmt.Println("PartTwo: ", PartTwo(lines, 2503))
}
