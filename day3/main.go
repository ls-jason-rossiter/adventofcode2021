package main

import (
	"fmt"
	"strconv"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	ss := trash.ReaderToStrings(f)

	part1(ss)
	part2(ss)
}

func part1(ss []string) {
	fmt.Println("power consumption:", gammaRate(ss)*epsilonRate(ss))
}

func part2(ss []string) {
	fmt.Println("life support rating:", oxygenGeneratorRating(ss)*co2ScrubberRating(ss))
}

func gammaRate(ss []string) int {
	if len(ss) == 0 {
		return 0
	}

	rate := make([]rune, 0, len(ss[0]))
	for i := 0; i < len(ss[0]); i++ {
		rate = append(rate, mostCommonBit(ss, i))
	}

	return MustParseIntBase2(string(rate))
}

func epsilonRate(ss []string) int {
	if len(ss) == 0 {
		return 0
	}

	rate := make([]rune, 0, len(ss[0]))
	for i := 0; i < len(ss[0]); i++ {
		rate = append(rate, leastCommonBit(ss, i))
	}

	return MustParseIntBase2(string(rate))
}

func mostCommonBit(ss []string, position int) rune {
	zeroCount := 0
	oneCount := 0
	for _, s := range ss {
		if s[position] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}
	if oneCount >= zeroCount {
		return '1'
	}
	return '0'
}

func leastCommonBit(ss []string, position int) rune {
	zeroes := 0
	ones := 0
	for _, s := range ss {
		if s[position] == '1' {
			ones++
		} else {
			zeroes++
		}
	}
	if zeroes <= ones {
		return '0'
	}
	return '1'
}

func oxygenGeneratorRating(ss []string) int {
	if len(ss) == 0 {
		return 0
	}

	candidates := make([]string, len(ss))
	copy(candidates, ss[:])
	for i := 0; i < len(ss[0]); i++ {
		candidates = filterByBit(candidates, i, mostCommonBit(candidates, i))
		if len(candidates) == 1 {
			return MustParseIntBase2(candidates[0])
		}
	}
	return 0
}

func co2ScrubberRating(ss []string) int {
	if len(ss) == 0 {
		return 0
	}

	candidates := make([]string, len(ss))
	copy(candidates, ss[:])
	for i := 0; i < len(ss[0]); i++ {
		candidates = filterByBit(candidates, i, leastCommonBit(candidates, i))
		if len(candidates) == 1 {
			return MustParseIntBase2(candidates[0])
		}
	}
	return 0
}

// GitHub Copilot generated 99% of this function :explode:
func filterByBit(ss []string, position int, bit rune) []string {
	var filtered []string
	for _, s := range ss {
		if rune(s[position]) == bit {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func MustParseIntBase2(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
