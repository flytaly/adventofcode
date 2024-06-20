package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Coords struct {
	x, y, z int
}

type Particle struct {
	index   int
	p, v, a Coords
}

func (p *Particle) tick() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func parse(input []string) []Particle {
	particles := make([]Particle, len(input))
	var parseCoords = func(s string, c *Coords) {
		_, err := fmt.Sscanf(s, "<%d,%d,%d>", &c.x, &c.y, &c.z)
		if err != nil {
			panic(err)
		}
	}
	for i, line := range input {
		split := strings.Split(line, ", ")
		particles[i] = Particle{}
		particles[i].index = i
		parseCoords(split[0][2:], &particles[i].p)
		parseCoords(split[1][2:], &particles[i].v)
		parseCoords(split[2][2:], &particles[i].a)
	}
	return particles
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(p Coords) int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func filter[T any](s []T, f func(T) bool) []T {
	out := []T{}
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// sort by acceleration and speed
func P1(input []string) int {
	pp := parse(input)

	sort.Slice(pp, func(i, j int) bool {
		return manhattan(pp[i].a) < manhattan(pp[j].a)
	})

	pp = filter(pp, func(p Particle) bool {
		return manhattan(p.a) == manhattan(pp[0].a)
	})

	sort.Slice(pp, func(i, j int) bool {
		return manhattan(pp[i].v) < manhattan(pp[j].v)
	})

	pp = filter(pp, func(p Particle) bool {
		return manhattan(p.v) == manhattan(pp[0].v)
	})
	return pp[0].index
}

func P2(input []string) int {
	particles := parse(input)
	prevLen, collisionTick := len(particles), 0
	// Stop if there are no changes in 100 ticks.
	// Probably, there should be a better stopping condition,
	// but it works anyway.
	for tick := 0; tick-collisionTick < 100; tick++ {
		places := map[Coords]int{}
		for p := range particles {
			particles[p].tick()
			places[particles[p].p] += 1
		}
		particles = filter(particles, func(p Particle) bool {
			return places[p.p] == 1
		})
		if prevLen != len(particles) {
			collisionTick = tick
			prevLen = len(particles)
		}
	}

	return len(particles)
}

func main() {
	lines := []string{
		"p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>",
		"p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>",
		"p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>",
		"p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
	fmt.Println("Part 2 =>", P2(lines))
}
