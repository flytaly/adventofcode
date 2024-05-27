package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[string, int]{
			{"1122", 3},
			{"91212129", 9},
		}
		tests.Run(t, P1)
	})
	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[string, int]{
			{"1212", 6},
			{"1221", 0},
		}
		tests.Run(t, P2)
	})
}
