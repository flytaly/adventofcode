package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			}, 605},
		}
		tests.Run(t, PartOne)
	})
	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			}, 982},
		}
		tests.Run(t, PartTwo)
	})
}
