package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	cs := parseInput(trash.ReaderToStrings(f))

	cs.Print()

	part1(cs)
	part2(cs)
}

func parseInput(ss []string) *CaveSystem {
	cs := &CaveSystem{
		Caves: make(map[string]*Cave),
	}
	for _, s := range ss {
		parts := strings.Split(s, "-")
		cs.AddCave(&Cave{
			Name:  parts[0],
			IsBig: unicode.IsUpper(rune(parts[0][0])),
		})
		cs.AddCave(&Cave{
			Name:  parts[1],
			IsBig: unicode.IsUpper(rune(parts[1][0])),
		})
		cs.AddPath(parts[0], parts[1])
		cs.AddPath(parts[1], parts[0])
	}
	return cs
}

func part1(cs *CaveSystem) {
	fmt.Println("part 1:", cs.TraversePartOne("start", "end", map[string]bool{"start": true}))
}

func part2(cs *CaveSystem) {
	fmt.Println("part 2:", cs.TraversePartTwo("start", "end", map[string]int{"start": 1}, false))
}

type Cave struct {
	Name          string
	IsBig         bool
	AdjacentCaves []*Cave
}

type CaveSystem struct {
	Caves map[string]*Cave
}

func (c *CaveSystem) AddCave(cave *Cave) {
	if c.Caves[cave.Name] == nil {
		c.Caves[cave.Name] = cave
	}
}

func (c *CaveSystem) TraversePartOne(from, to string, visited map[string]bool) int {
	if from == to {
		return 1
	}

	count := 0
	fromCave := c.Cave(from)

	for _, cave := range fromCave.AdjacentCaves {
		if visited[cave.Name] && !cave.IsBig {
			continue
		}

		visited[cave.Name] = true

		count += c.TraversePartOne(cave.Name, to, visited)

		// Set "from" to unvisited so that we can backtrack.
		visited[cave.Name] = false
	}

	return count
}

func (c *CaveSystem) TraversePartTwo(from, to string, visited map[string]int, smallCaveVisitedTwice bool) int {
	if from == to {
		return 1
	}

	count := 0
	fromCave := c.Cave(from)

	for _, cave := range fromCave.AdjacentCaves {
		if !cave.IsBig && visited[cave.Name] > 0 {
			if smallCaveVisitedTwice {
				continue
			}
			smallCaveVisitedTwice = true
		}

		visited[cave.Name]++

		count += c.TraversePartTwo(cave.Name, to, visited, smallCaveVisitedTwice)

		// Set "from" to unvisited so that we can backtrack.
		// Also toggle the visitedTwice flag if necessary.
		visited[cave.Name]--
		if !cave.IsBig && visited[cave.Name] == 1 {
			smallCaveVisitedTwice = false
		}
	}

	return count
}

func (c *CaveSystem) Print() {
	for _, cave := range c.Caves {
		fmt.Printf("Cave %s (big: %v):", cave.Name, cave.IsBig)
		for _, c2 := range cave.AdjacentCaves {
			fmt.Printf(" %v", c2.Name)
		}
		fmt.Println()
	}
}

func (c *CaveSystem) AddPath(from, to string) {
	fromCave := c.Cave(from)
	toCave := c.Cave(to)

	if to == "start" {
		return
	}

	if fromCave != nil && toCave != nil {
		fromCave.AdjacentCaves = append(fromCave.AdjacentCaves, toCave)
	}
}

func (c *CaveSystem) Cave(name string) *Cave {
	return c.Caves[name]
}
