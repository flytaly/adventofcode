package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[string, string]{
			{"abc", "18f47a30"}}
		tests.Run(t, PartOne)
	})
}
