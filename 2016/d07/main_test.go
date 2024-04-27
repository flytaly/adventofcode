package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[string, bool]{
			{"abba[mnop]qrst", true},
			{"abcd[bddb]xyyx", false},
			{"aaaa[qwer]tyui", false},
			{"ioxxoj[asdfgh]zxcvbn", true},
		}
		tests.Run(t, supportTLS)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[string, bool]{
			{"aba[bab]xyz", true},
			{"xyx[xyx]xyx", false},
			{"aaa[kek]eke", true},
			{"zazbz[bzb]cdb", true},
		}
		tests.Run(t, supportSSL)
	})
}
