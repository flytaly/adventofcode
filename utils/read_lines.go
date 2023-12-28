package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadLines(inputFile string) []string {
	if !filepath.IsAbs(inputFile) {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		inputFile = filepath.Join(path, inputFile)
	}
	f, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}
