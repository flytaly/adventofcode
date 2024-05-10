package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

const FLOORS_NUM = 4

type Floors struct {
	data    [][]string
	itemMap map[string]int
	items   *[]string // maintain stable item's order for all instances
	current int       // current floor
	lowest  int       // lowest floor
}

func (f *Floors) add(floor int, item string) {
	f.data[floor] = append(f.data[floor], item)
	// f.data[floor][item] = true
	f.itemMap[item] = floor
}

func (f *Floors) move(to int, items ...string) {
	for _, item := range items {
		f.add(to, item)
		f.itemMap[item] = to
	}

	// remove
	floor := []string{}
	for _, name := range f.data[f.current] {
		if !slices.Contains(items, name) {
			floor = append(floor, name)
		}
	}
	f.data[f.current] = floor

	f.current = to
}

func (f Floors) getCurrentItems() []string {
	return f.data[f.current]
}

func isValidFloor(floorItems map[string]bool) bool {
	chips := []string{}
	hasGenerators := false

	for item, ok := range floorItems {
		if !ok {
			continue
		}
		itemId, itemType := item[0:1], item[1:2]
		if itemType == "G" {
			hasGenerators = true
		}
		chips = append(chips, itemId)
	}

	if !hasGenerators {
		return true
	}

	for _, chipId := range chips {
		if !floorItems[chipId+"G"] {
			// chip without generator
			return false
		}
	}

	return true
}

func isValidRemove(hasItems []string, items ...string) bool {
	potential := map[string]bool{}
	for _, item := range hasItems {
		potential[item] = true
	}
	for _, item := range items {
		potential[item] = false
	}
	return isValidFloor(potential)
}

func isValidAddition(hasItems []string, items ...string) bool {
	potential := map[string]bool{}
	for _, item := range hasItems {
		potential[item] = true
	}
	for _, item := range items {
		potential[item] = true
	}
	return isValidFloor(potential)
}

func (f Floors) isValidMove(to int, items ...string) bool {
	if !isValidRemove(f.data[f.current], items...) {
		return false
	}
	if !isValidAddition(f.data[to], items...) {
		return false
	}
	return true
}

func (f Floors) clone() Floors {
	newFloors := make([][]string, len(f.data))
	for i, floor := range f.data {
		newFloors[i] = slices.Clone(floor)
	}
	newItemMap := make(map[string]int)
	for item, floor := range f.itemMap {
		newItemMap[item] = floor
	}
	return Floors{
		data:    newFloors,
		itemMap: newItemMap,
		current: f.current,
		items:   f.items,
		lowest:  f.lowest,
	}
}

func (f Floors) serialize() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("E:%d", f.current))
	for _, item := range *(f.items) {
		sb.WriteString(fmt.Sprintf(";%s:%d", item, f.itemMap[item]))
	}
	return sb.String()
}

func createFloors() Floors {
	floors := make([][]string, FLOORS_NUM)
	for i := 0; i < FLOORS_NUM; i++ {
		floors[i] = make([]string, 0)
	}
	f := Floors{
		data:    floors,
		items:   &[]string{},
		itemMap: map[string]int{},
	}
	return f
}

func parse(input []string) Floors {
	generatorsRe := regexp.MustCompile(`a (\w+) generator`)
	chipsRe := regexp.MustCompile(`a (\w+)-compatible microchip`)

	floors := createFloors()

	for i, line := range input {
		for _, match := range generatorsRe.FindAllStringSubmatch(line, -1) {
			name := strings.ToUpper(match[1][0:1]) + "G"
			floors.add(i, name)
			*floors.items = append(*floors.items, name)
		}
		for _, match := range chipsRe.FindAllStringSubmatch(line, -1) {
			name := strings.ToUpper(match[1][0:1]) + "M"
			floors.add(i, name)
			*floors.items = append(*floors.items, name)
		}
	}

	return floors
}

func Count(lines []string, partTwo bool) int {
	initial := parse(lines)
	if partTwo {
		for _, name := range []string{"EG", "EM", "DG", "DM"} {
			initial.add(0, name)
			*initial.items = append(*initial.items, name)
		}
	}
	memo := map[string]struct{}{}

	// moves one or two items
	moveItems := func(floors Floors, to int, indices ...int) (*Floors, bool) {
		items := []string{floors.getCurrentItems()[indices[0]]}
		if len(indices) > 1 {
			items = append(items, floors.getCurrentItems()[indices[1]])
		}
		if !floors.isValidMove(to, items...) {
			return nil, false
		}

		f := floors.clone()

		f.move(to, items...)

		stateHash := f.serialize()
		if _, has := memo[stateHash]; has {
			return nil, false
		}
		memo[stateHash] = struct{}{}
		if len(f.data[FLOORS_NUM-1]) == len(*f.items) {
			return &f, true
		}

		if len(f.data[f.lowest]) == 0 {
			f.lowest = f.current
		}
		return &f, false
	}

	frontier := []Floors{initial}
	for step := 1; step < 1000; step++ {
		nextFrontier := []Floors{}
		for _, floors := range frontier {
			floorItems := floors.getCurrentItems()
			for i := 0; i < len(floorItems); i++ {
				for _, dir := range []int{1, -1} {
					nextFloor := floors.current + dir
					if nextFloor < floors.lowest || nextFloor >= len(floors.data) {
						continue
					}
					moves := [][]int{}
					for j := i + 1; j < len(floorItems); j++ {
						moves = append(moves, []int{i, j})
					}
					moves = append(moves, []int{i})
					for _, ids := range moves {
						newFloors, isEnd := moveItems(floors, nextFloor, ids...)
						if isEnd {
							fmt.Printf("FOUND! STEP %d. Frontier size %d\n", step, len(frontier))
							return step
						}
						if newFloors != nil {
							nextFrontier = append(nextFrontier, *newFloors)
						}
					}
				}
			}
		}
		frontier = nextFrontier
	}

	return -1
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	ts := time.Now()
	fmt.Printf("PartOne: %d [%s]\n", Count(lines, false), time.Since(ts).Truncate(time.Millisecond))
	ts = time.Now()
	fmt.Printf("PartTwo: %d [%s]\n", Count(lines, true), time.Since(ts).Truncate(time.Millisecond))
}
