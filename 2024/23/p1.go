package main

import (
	"slices"
	"strings"
)

func findSets(links map[string][]string, node string) [][]string {
	rest := [][]string{}

	for i, id1 := range links[node] {
		for _, id2 := range links[node][i+1:] {
			rest = append(rest, []string{id1, id2})
		}
	}

	sets := [][]string{}
	for _, names := range rest {
		if slices.Contains(links[names[0]], names[1]) {
			sets = append(sets, []string{node, names[0], names[1]})
		}
	}

	return sets
}

func PartOne(lines []string) (result int) {
	net := map[string][]string{}
	for _, l := range lines {
		ids := strings.Split(l, "-")
		net[ids[0]] = append(net[ids[0]], ids[1])
		net[ids[1]] = append(net[ids[1]], ids[0])
	}

	sets := map[[3]string]struct{}{}
	for node := range net {
		for _, set := range findSets(net, node) {
			slices.Sort(set)
			sets[[3]string(set)] = struct{}{}
		}
	}

	var startWithT = func(s string) bool {
		return strings.HasPrefix(s, "t")
	}

	for set, _ := range sets {
		if slices.ContainsFunc(set[:], startWithT) {
			result++
		}
	}

	return result
}
