package d23

import (
	"fmt"
)

func copyMap(m map[complex128]bool) map[complex128]bool {
	n := make(map[complex128]bool)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func parse(lines []string) (map[complex128]string, complex128, complex128) {
	grid := make(map[complex128]string)
	var startPos, endPos complex128
	for i, l := range lines {
		for j, v := range l {
			c := complex(float64(i), float64(j))
			grid[c] = string(v)
			if i == 0 && v == '.' {
				startPos = c
			}
			if i == len(lines)-1 && v == '.' {
				endPos = c
			}
		}
	}
	return grid, startPos, endPos
}

type Path struct {
	pos      complex128
	visited  map[complex128]bool
	finished bool
}

func (p Path) clone() Path {
	return Path{pos: p.pos, visited: copyMap(p.visited), finished: p.finished}
}

func (p *Path) addPos(newPos complex128) {
	p.visited[newPos] = true
	p.pos = newPos
}

func longestPath(paths []Path) (res int, index int) {
	mx, id := 0, 0
	for i, p := range paths {
		if len(p.visited) > mx {
			mx = len(p.visited)
			id = i
		}
	}
	return mx, id
}

func PartOne(lines []string) {
	grid, pos, _ := parse(lines)

	paths := []Path{{
		pos:      pos,
		finished: false,
		visited:  map[complex128]bool{pos: true},
	}}

	for {
		moves := 0
		newPaths := []Path{}
		for pathIdx, path := range paths {
			if path.finished {
				continue
			}
			pos, visited := path.pos, path.visited
			visitQ := []complex128{}
			for _, dir := range []complex128{-1, -1i, 1, 1i} {
				npos := pos + dir
				cell := grid[npos]
				if cell == "" || cell == "#" || visited[npos] {
					continue
				}
				if dir == -1 && cell == "v" {
					continue
				}
				if dir == 1 && cell == "^" {
					continue
				}
				if dir == -1i && cell == ">" {
					continue
				}
				if dir == 1i && cell == "<" {
					continue
				}

				if int(real(npos)) == len(lines)-1 {
					path.finished = true
					paths[pathIdx] = path
					break
				}
				visitQ = append(visitQ, npos)
			}

			moves += len(visitQ)

			if len(visitQ) == 0 {
				continue
			}

			// Add new paths first, and then update the current path
			// so it won't be cloned into new paths.
			for _, pos := range visitQ[1:] {
				newPath := path.clone()
				newPath.addPos(pos)
				newPaths = append(newPaths, newPath)
			}
			path.addPos(visitQ[0])
			paths[pathIdx] = path
		}
		if moves == 0 {
			break
		}
		paths = append(paths, newPaths...)
	}

	result, id := longestPath(paths)
	PrintPathComplex(lines, paths[id].visited)

	fmt.Println("Part 1:", result)
}
