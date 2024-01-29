package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"aaaaa-bbb-z-y-x-123[abxyz]",
				"a-b-c-d-e-f-g-h-987[abcde]",
				"not-a-real-room-404[oarel]",
				"totally-real-room-200[decoy]",
			}, 1514},
		}
		tests.Run(t, PartOne)
	})

	t.Run("decrypt", func(t *testing.T) {
		room := Room{name: "qzmt-zixmtkozy-ivhz", sector: 343}
		want := "very encrypted name"
		if got := room.decrypt(); got != want {
			t.Errorf("decrypt() = %v, want %v", got, want)
		}
	})
}
