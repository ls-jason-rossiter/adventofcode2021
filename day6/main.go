package main

import (
	"fmt"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

// Starcraft Zerg theme:
// - fish = larvae
// - zero timer fish = queens

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	hatchery := parseInput(trash.ReaderToStrings(f))

	part1(hatchery)
	part2(hatchery)
}

func part1(hatchery Hatchery) {
	fmt.Printf("Part 1: %d larvae after 80 days\n", hatchery.LarvaCountAfter(80))
}

func part2(hatchery Hatchery) {
	fmt.Printf("Part 2: %d larvae after 256 days\n", hatchery.LarvaCountAfter(256))
}

// Hatchery is a slice of days (index) and larva counters (value).
// Using a slice rather than a fixed array as it makes it easier to manipulate.
type Hatchery []int

func (h Hatchery) LarvaCountAfter(days int) int {
	hatchery := h
	for i := 0; i < days; i++ {
		// Move all counters down as a day progresses.
		// Each larva with a zero timer becomes a queen and injects larva into
		// the hatchery at timer 8 (last position in slice).
		hatchery = append(hatchery[1:], hatchery[0])

		// The queens explode and turn back into larvae in timer 6 :sad:
		hatchery[6] += hatchery[8]
	}

	total := 0
	for _, i := range hatchery {
		total += i
	}

	return total
}

func parseInput(ss []string) Hatchery {
	split := strings.Split(ss[0], ",")
	hatchery := make([]int, 9)
	for _, s := range split {
		hatchery[trash.MustParseIntBase10(s)]++
	}
	return hatchery
}
