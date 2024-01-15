package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stat struct {
	HP     int
	Damage int
	Armor  int
}

func (s Stat) hit(enemy Stat) int {
	return max(1, s.Damage-enemy.Armor)
}

func (s *Stat) addItems(items ...Item) {
	for _, item := range items {
		s.Damage += item.Damage
		s.Armor += item.Armor
	}
}

func (player Stat) isWin(boss Stat) bool {
	// ceil division
	playerHits := 1 + (boss.HP-1)/player.hit(boss)
	bossHits := 1 + (player.HP-1)/boss.hit(player)
	return playerHits <= bossHits
}

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

var (
	weapons = []Item{{8, 4, 0}, {10, 5, 0}, {25, 6, 0}, {40, 7, 0}, {74, 8, 0}}
	armor   = []Item{{0, 0, 0}, {13, 0, 1}, {31, 0, 2}, {53, 0, 3}, {75, 0, 4}, {102, 0, 5}}
	rings   = []Item{{0, 0, 0}, {25, 1, 0}, {50, 2, 0}, {100, 3, 0}, {20, 0, 1}, {40, 0, 2}, {80, 0, 3}}
)

func parser(lines []string) Stat {
	stats := Stat{}
	re := regexp.MustCompile(`(^.+): (\d+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		num, _ := strconv.Atoi(match[2])
		switch match[1] {
		case "Damage":
			stats.Damage = num
		case "Armor":
			stats.Armor = num
		case "Hit Points":
			stats.HP = num
		}
	}

	return stats
}

func ringCombinations() [][2]Item {
	result := make([][2]Item, 0)
	result = append(result, [2]Item{{0, 0, 0}, {0, 0, 0}})
	for i := 0; i < len(rings)-1; i++ {
		for j := i + 1; j < len(rings); j++ {
			result = append(result, [2]Item{rings[i], rings[j]})
		}
	}
	return result
}

func PartOne(lines []string) int {
	boss := parser(lines)
	ringComb := ringCombinations()
	gold := 0
	for _, weapon := range weapons[2:] {
		for _, arm := range armor {
			for _, ring := range ringComb {
				player := Stat{100, 0, 0}
				player.addItems(weapon, arm, ring[0], ring[1])
				if !player.isWin(boss) {
					continue
				}
				cost := weapon.Cost + arm.Cost + ring[0].Cost + ring[1].Cost
				if gold == 0 || cost < gold {
					gold = cost
				}
			}
		}
	}

	return gold
}

func PartTwo(lines []string) int {
	boss := parser(lines)
	ringComb := ringCombinations()
	gold := -1
	for _, weapon := range weapons {
		for _, arm := range armor {
			for _, ring := range ringComb {
				player := Stat{100, 0, 0}
				player.addItems(weapon, arm, ring[0], ring[1])
				if player.isWin(boss) {
					continue
				}
				cost := weapon.Cost + arm.Cost + ring[0].Cost + ring[1].Cost
				if cost > gold {
					gold = cost
				}
			}
		}
	}

	return gold
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
