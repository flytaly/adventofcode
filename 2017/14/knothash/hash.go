// copy from day10

package knothash

import (
	"fmt"
	"strings"
)

func reverseCircular[T any](s []T, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		ii, jj := i%len(s), j%len(s)
		s[ii], s[jj] = s[jj], s[ii]
	}
}

func hashRound(positions []int, nums []int, pos, skip int) (int, int) {
	for _, section := range nums {
		reverseCircular(positions, pos, pos+section)
		pos = (pos + section + skip) % len(positions)
		skip = (skip + 1) % len(positions)
	}
	return pos, skip
}

func toDenseHash(sparse []int) [16]int {
	dense := [16]int{}
	for pos := 0; pos < 256; pos = pos + 16 {
		n := sparse[pos]
		for j := pos + 1; j < pos+16; j++ {
			n ^= sparse[j]
		}
		dense[pos/16] = n
	}
	return dense
}

func Hash(input string) string {
	positions := make([]int, 256)
	for i := 0; i < len(positions); i++ {
		positions[i] = i
	}

	chars := make([]int, len(input))
	for i, v := range input {
		chars[i] = int(v)
	}
	chars = append(chars, []int{17, 31, 73, 47, 23}...)

	var pos, skip int
	for i := 0; i < 64; i++ {
		pos, skip = hashRound(positions, chars, pos, skip)
	}

	output := strings.Builder{}
	for _, v := range toDenseHash(positions) {
		output.WriteString(fmt.Sprintf("%02x", v))
	}
	return output.String()
}
