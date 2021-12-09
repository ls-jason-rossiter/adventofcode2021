package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	rows := parseInput(trash.ReaderToStrings(f))

	part1(rows)
	part2(rows)
}

// Returns a map of crab positions and their counts.
func parseInput(ss []string) []Row {
	var rows []Row
	for _, row := range ss {
		data := strings.SplitN(row, " | ", 2)
		rows = append(rows, Row{
			SignalPattern:   strings.Fields(data[0]),
			FourDigitOutput: strings.Fields(data[1]),
		})
	}
	return rows
}

type Row struct {
	SignalPattern   []string
	FourDigitOutput []string
}

func part1(rows []Row) {
	appearances := 0
	for _, row := range rows {
		for _, s := range row.FourDigitOutput {
			sLen := len(s)
			if sLen == 2 || sLen == 4 || sLen == 3 || sLen == 7 {
				appearances++
			}
		}
	}

	fmt.Println("part 1:", appearances)
}

func part2(rows []Row) {
	sumRows := 0
	for _, row := range rows {
		knownDigits := make([]string, 10)

		knownDigits[1] = findInRowFn(row, func(s string) bool {
			return len(s) == 2
		})
		knownDigits[4] = findInRowFn(row, func(s string) bool {
			return len(s) == 4
		})
		knownDigits[7] = findInRowFn(row, func(s string) bool {
			return len(s) == 3
		})
		knownDigits[8] = findInRowFn(row, func(s string) bool {
			return len(s) == 7
		})
		knownDigits[6] = findInRowFn(row, func(s string) bool {
			return len(s) == 6 && stringIntersection(s, knownDigits[7]) == 2
		})
		knownDigits[3] = findInRowFn(row, func(s string) bool {
			return len(s) == 5 && stringIntersection(s, knownDigits[1]) == 2
		})
		knownDigits[9] = findInRowFn(row, func(s string) bool {
			return len(s) == 6 && stringIntersection(s, knownDigits[4]) == 4
		})
		knownDigits[0] = findInRowFn(row, func(s string) bool {
			return len(s) == 6 && stringIntersection(s, knownDigits[4]) == 3 && stringIntersection(s, knownDigits[7]) == 3
		})
		knownDigits[2] = findInRowFn(row, func(s string) bool {
			return len(s) == 5 && stringIntersection(s, knownDigits[4]) == 2
		})
		knownDigits[5] = findInRowFn(row, func(s string) bool {
			return len(s) == 5 && stringIntersection(s, knownDigits[6]) == 5
		})

		var sortedFourDigitOutput []string
		for _, s := range row.FourDigitOutput {
			sortedFourDigitOutput = append(sortedFourDigitOutput, sortStringByCharacters(s))
		}

		fullCode := ""
		for _, s := range sortedFourDigitOutput {
			for number, code := range knownDigits {
				if s == code {
					fullCode += strconv.Itoa(number)
				}
			}
		}

		sumRows += trash.MustParseIntBase10(fullCode)
	}

	fmt.Println("part 2:", sumRows)
}

func findInRowFn(row Row, fn func(string) bool) string {
	for _, sp := range row.SignalPattern {
		if fn(sp) {
			return sortStringByCharacters(sp)
		}
	}
	panic("something's wrong!")
}

func stringIntersection(a, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				count++
			}
		}
	}
	return count
}

func sortStringByCharacters(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}
