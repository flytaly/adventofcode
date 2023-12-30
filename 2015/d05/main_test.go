package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{"ugknbfddgicrmopn"}, 1},
			{[]string{"aaa"}, 1},
			{[]string{"jchzalrnumimnmhp"}, 0},
			{[]string{"haegwjzuvuyypxyu"}, 0},
			{[]string{"dvszwmarrgswjxmb"}, 0},
		}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{"qjhvhtzxzqqjkmpb"}, 1},
			{[]string{"xxyxx"}, 1},
			{[]string{"xxxx"}, 1},
			{[]string{"uurcxstgmygtbstg"}, 0},
			{[]string{"ieodomkazucvgmuy"}, 0},
		}
		tests.Run(t, PartTwo)
	})
}
