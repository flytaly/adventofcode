package utils

import "strconv"

// converts to slice of ints ignoring errors
func ToInts(s []string) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}
