package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		ops := []string{
			"rect 3x2",
			"rotate column x=1 by 1",
			"rotate row y=0 by 4",
			"rotate column x=1 by 1",
		}
		grid := lightGrid(ops, 7, 3)
		fmt.Println(grid)
		res := grid.count()
		if res != 6 {
			t.Errorf("PartOne() = %v, want %v", res, 6)
		}
	})
}
