package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

type Coords struct {
	x, y, z int
}

type Brik struct {
	idx    int
	coords []Coords
}

func (b *Brik) drop(n int) {
	b.coords[0].z -= n
	b.coords[1].z -= n
}

func readLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func parse(lines []string) [][]Coords {
	res := [][]Coords{}
	for i, line := range lines {
		split := strings.Split(line, "~")
		res = append(res, []Coords{})
		for _, v := range split {
			xyz := strings.Split(v, ",")
			nums := []int{}
			for _, c := range xyz {
				num, _ := strconv.Atoi(c)
				nums = append(nums, num)
			}
			res[i] = append(res[i], Coords{nums[0], nums[1], nums[2]})
		}
	}
	return res

}

func isIntersect(c1, c2 []Coords) bool {
	if c1[1].x < c2[0].x || c1[0].x > c2[1].x {
		return false
	}
	if c1[1].y < c2[0].y || c1[0].y > c2[1].y {
		return false
	}

	return true
}

func parseBriks(lines []string) []Brik {
	parsed := parse(lines)
	briks := []Brik{}
	for i, cc := range parsed {
		briks = append(briks, Brik{idx: i, coords: cc})
	}

	sort.SliceStable(briks, func(i, j int) bool {
		return briks[i].coords[0].z < briks[j].coords[0].z
	})

	for i := range briks {
		briks[i].idx = i
	}

	return briks
}

type Levels map[int]struct {
	bot []int
	top []int
}

func (ll *Levels) add(brik Brik) {
	b, t := brik.coords[0].z, brik.coords[1].z
	level := (*ll)[b]
	level.bot = append(level.bot, brik.idx)
	(*ll)[b] = level

	level = (*ll)[t]
	level.top = append(level.top, brik.idx)
	(*ll)[t] = level
}

func fallCount(levels Levels, briks []Brik, brik Brik) (count int) {
	for z := brik.coords[0].z - 1; z >= 1; z-- {
		for _, prevIdx := range levels[z].top {
			if isIntersect(brik.coords, briks[prevIdx].coords) {
				return count
			}
		}
		count++
	}
	return count
}

func fallSym(briks []Brik) Levels {
	levels := make(Levels) // level => briks
	for _, brik := range briks {
		fell := fallCount(levels, briks, brik)
		brik.drop(fell)
		levels.add(brik)
	}
	return levels
}

func PartOne(lines []string) {
	briks := parseBriks(lines)

	levels := fallSym(briks)
	support := map[int]sets.Set[int]{}
	supported := map[int]sets.Set[int]{}
	for _, v := range briks {
		support[v.idx] = sets.New[int]()
		supported[v.idx] = sets.New[int]()
	}

	for i, ll := range levels {
		for _, bellow := range ll.top {
			for _, above := range levels[i+1].bot {
				if isIntersect(briks[bellow].coords, briks[above].coords) {
					support[bellow].Insert(above)
					supported[above].Insert(bellow)
				}
			}
		}
	}

	// fmt.Println(support)
	// fmt.Println(supported)

	count := 0
	// check if every brik above has more than one support
	for _, brik := range briks {
		haveOtherSupport := 1
		for above := range support[brik.idx] {
			if supported[above].Len() < 2 {
				haveOtherSupport = 0
				break
			}
		}
		count += haveOtherSupport
	}

	fmt.Println("Part 1:", count)
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines) //497
}
