package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	crabs := parseInput(trash.ReaderToStrings(f))

	part1(crabs)
	part2(crabs)
}

// Returns a map of crab positions and their counts.
func parseInput(ss []string) map[int]int {
	positions := strings.Split(ss[0], ",")
	crabPositionCounts := make(map[int]int)
	for _, s := range positions {
		v := trash.MustParseIntBase10(s)
		crabPositionCounts[v]++
	}
	return crabPositionCounts
}

func part1(crabs map[int]int) {
	var fuelCosts []int
	for crabPos := range crabs {
		fuelCosts = append(fuelCosts, constantFuelCost(crabPos, crabs))
	}

	sort.Ints(fuelCosts)

	fmt.Println("part 1:", fuelCosts[0])
}

func part2(crabs map[int]int) {
	maxCrabPos := 0
	for crabPos := range crabs {
		if crabPos > maxCrabPos {
			maxCrabPos = crabPos
		}
	}

	var fuelCosts []int
	for i := 0; i < maxCrabPos; i++ {
		fuelCosts = append(fuelCosts, scalingFuelCost(i, crabs))
	}

	sort.Ints(fuelCosts)

	fmt.Println("part 2:", fuelCosts[0])
}

func constantFuelCost(position int, crabs map[int]int) int {
	cost := 0
	for crabPos, count := range crabs {
		cost += int(math.Abs(float64(crabPos-position))) * count
	}
	return cost
}

func scalingFuelCost(position int, crabs map[int]int) int {
	cost := 0
	for crabPos, count := range crabs {
		stepCount := int(math.Abs(float64(crabPos - position)))
		cost += (stepCount * (stepCount + 1) / 2) * count
	}
	return cost
}
