package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

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

func parse(input []string) (map[string][]string, []map[string]int) {
	re := regexp.MustCompile(`(\w+){(.+)}`)
	p1 := true
	workflows := map[string][]string{}
	ratings := []map[string]int{}

	for _, l := range input {
		if l == "" {
			p1 = false
			continue
		}
		if p1 {
			m := re.FindAllStringSubmatch(l, -1)[0]
			workflows[m[1]] = strings.Split(m[2], ",")
			continue
		}
		split := strings.Split(l[1:len(l)-1], ",")
		rating := map[string]int{}
		for _, v := range split {
			rt := strings.Split(v, "=")
			n, _ := strconv.Atoi(rt[1])
			rating[rt[0]] = n
		}
		ratings = append(ratings, rating)
	}

	return workflows, ratings
}

type Rule struct {
	ctg  string
	sign byte
	num  int
	dest string
}

func (r Rule) compare(a int) bool {
	if r.sign == '<' {
		return a < r.num
	}
	return a > r.num
}

func parseRule(r string) Rule {
	split := strings.Split(r, ":")
	id := ""
	if len(split) == 1 {
		return Rule{dest: split[0]}
	}
	cond, id := split[0], split[1]
	category, sign := string(cond[0]), cond[1]
	num, _ := strconv.Atoi(cond[2:])
	return Rule{ctg: category, sign: sign, num: num, dest: id}
}

func sumMap(s map[string]int) (res int) {
	for _, v := range s {
		res += v
	}
	return res
}

func PartOne(lines []string) {
	workflows, ratings := parse(lines)
	res := 0
	for _, rating := range ratings {
		for id := "in"; id != "A" && id != "R"; {
			for _, flow := range workflows[id] {
				rule := parseRule(flow)
				if rule.ctg == "" {
					id = rule.dest
				} else if rule.compare(rating[rule.ctg]) {
					id = rule.dest
					if id != "A" && id != "R" {
						break
					}
				}
				if id == "A" {
					res += sumMap(rating)
				}
				if id == "A" || id == "R" {
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", res)
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
