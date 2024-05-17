package main

import (
	"fmt"
	"os"
	"time"
)

func countTiles(input string, rows int) (count int) {
	row := make([]byte, len(input)+2)
	for i, v := range input {
		if v == '^' {
			row[i+1] = 1
			continue
		}
		count++
	}

	temp := make([]byte, len(row))
	for r := 0; r < rows-1; r++ {
		for i := 1; i < len(row)-1; i++ {
			temp[i] = row[i-1] ^ row[i+1]
			if temp[i] == 0 {
				count++
			}
		}
		row, temp = temp, row
	}

	return count
}

func main() {
	input, rows := ".^^.^.^^^^", 10
	if len(os.Args) > 1 {
		input = os.Args[1]
		rows = 40
	}
	fmt.Println("PartOne: ", countTiles(input, rows))
	ts := time.Now()
	fmt.Println("PartTwo: ", countTiles(input, 400000), time.Since(ts))
}
