package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	if !filepath.IsAbs(inputFile) {
		inputFile = filepath.Join(path.Dir(filename), inputFile)
	}
	f, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}
