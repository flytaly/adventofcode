package main

import (
	"fmt"
	"os"
	"strconv"
)

func PartOne(row, col int) int {
	val := 20151125
	for r, c := 1, 1; c != col || r != row; {
		c += 1
		r -= 1
		if r <= 0 {
			r = c
			c = 1
		}
		val = (val * 252533) % 33554393
	}
	return val
}

func main() {
	row, col := 2, 5
	if len(os.Args) > 2 {
		row, _ = strconv.Atoi(os.Args[1])
		col, _ = strconv.Atoi(os.Args[2])
	}
	fmt.Printf("row %d col %d\n", row, col)
	fmt.Println("PartOne: ", PartOne(row, col))
}
