package p2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var Cards = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func CardStrenght(card string) int {
	for index, v := range Cards {
		if v == card {
			return index
		}
	}
	return 0
}

func HandStrenght(hand string) int {
	hands := map[string]int{
		"11111": 0, // High card
		"2111":  1, // One pair
		"221":   2,
		"311":   3,
		"32":    4,
		"41":    5,
		"5":     6,
	}

	m := map[string]int{}
	for i := 0; i < len(hand); i++ {
		m[string(hand[i])]++
	}
	joker := 0
	if len(m) > 1 {
		joker = m["J"]
		delete(m, "J")
	}
	repeats := []string{}
	for _, v := range m {
		repeats = append(repeats, fmt.Sprint(v))
	}
	sort.Slice(repeats, func(i, j int) bool {
		return repeats[i] > repeats[j]
	})
	strongest, _ := strconv.Atoi(repeats[0])
	repeats[0] = fmt.Sprint(strongest + joker)
	handType := strings.Join(repeats, "")

	return hands[handType]
}

func IsStronger(a, b string) bool {
	valA := HandStrenght(a)
	valB := HandStrenght(b)
	if valA == valB {
		for i := 0; i < 5; i++ {
			c1S := CardStrenght(string(a[i]))
			c2S := CardStrenght(string(b[i]))
			if c1S == c2S {
				continue
			}
			return c1S > c2S
		}

	}
	return valA > valB
}
