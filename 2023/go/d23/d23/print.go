package d23

import (
	"fmt"
	"strings"
)

func PrintPathComplex(grid []string, visited map[complex128]bool) {
	count := 0
	sb := strings.Builder{}
	for i := range grid {
		for j := 0; j < len(grid[i]); j++ {
			if visited[complex(float64(i), float64(j))] {
				sb.WriteRune('O')
				count++
				continue
			}
			sb.WriteByte(grid[i][j])
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
