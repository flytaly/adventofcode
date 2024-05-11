package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func getSequences(str string) (firstTriplet string, fives []string) {
	firstTriplet = ""
	fives = []string{}

	prev, size := rune(0), 0
	for i, ch := range str {
		if ch != prev {
			prev, size = ch, 1
			continue
		}
		size += 1
		if size == 3 && firstTriplet == "" {
			firstTriplet = str[i-2 : i+1]
		}
		if size == 5 {
			fives = append(fives, str[i-4:i+1])
		}
	}
	return firstTriplet, fives
}

func getHash(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func getRepeatedHash(input string, repeat int) string {
	hash := input
	for i := 0; i < repeat; i++ {
		hash = getHash(hash)
	}
	return hash
}

func max[T int](a []T) T {
	b := a[0]
	for _, v := range a {
		if v > b {
			b = v
		}
	}

	return b
}

func getKey(input string, repeat int) int {
	keys := []int{}
	triplets := map[string][]int{}

	for i := 0; len(keys) < 64; i++ {
		h := getRepeatedHash(input+fmt.Sprint(i), repeat)
		triplet, fives := getSequences(h)
		if triplet == "" {
			continue
		}
		boundary := i - 1000
		for _, five := range fives {
			tri := five[0:3]
			for _, index := range triplets[tri] {
				if index < boundary {
					continue
				}
				keys = append(keys, index)
				if len(keys) == 64 {
					break
				}
			}
		}
		triplets[triplet] = append(triplets[triplet], i)
	}

	return max(keys)
}

func main() {
	input := "abc"
	if len(os.Args) >= 2 {
		input = os.Args[1]
	}
	fmt.Printf("%q\n", input)
	ts := time.Now()
	fmt.Println("PartOne: ", getKey(input, 1), time.Since(ts))
	ts = time.Now()
	fmt.Println("PartTwo: ", getKey(input, 2017), time.Since(ts))
}
