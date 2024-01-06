package main

import (
	"aoc/utils"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func recursiveSum(node interface{}, skipRed bool) (count int) {
	switch result := node.(type) {
	case map[string]interface{}: // obj
		for _, v := range result {
			if skipRed && v == "red" {
				return 0
			}
			count += recursiveSum(v, skipRed)
		}
	case []interface{}: // array
		for _, v := range result {
			count += recursiveSum(v, skipRed)
		}
	default: // value
		switch reflect.TypeOf(result).Kind() {
		case reflect.Int:
			count += result.(int)
		case reflect.Float64:
			count += int(result.(float64))
		}
	}

	return count
}

func PartOne(lines string) int {
	var data interface{}
	_ = json.Unmarshal([]byte(lines), &data)
	return recursiveSum(data, false)
}

func PartTwo(lines string) int {
	var data interface{}
	_ = json.Unmarshal([]byte(lines), &data)
	return recursiveSum(data, true)
}

func main() {
	lines := []string{`[1,{"c":"red","b":2},3]`}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines[0]))
	fmt.Println("PartTwo: ", PartTwo(lines[0]))
}
