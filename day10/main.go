package main

import (
	"fmt"
	"sort"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

var part1Scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var part2Scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	part2(part1(trash.ReaderToStrings(f)))
}

func part1(ss []string) []IncompleteLine {
	var incompleteLines []IncompleteLine
	var invalidChars []rune
	for _, line := range ss {
		isValid, badChar, stack := checkLine(line)
		if isValid {
			continue
		}

		isCorruptedLine := badChar != 0
		if isCorruptedLine {
			invalidChars = append(invalidChars, badChar)
		} else {
			incompleteLines = append(incompleteLines, IncompleteLine{
				line:  line,
				stack: stack,
			})
		}
	}

	fmt.Println("part 1:", calculateBadCharacterScore(invalidChars))

	return incompleteLines
}

func part2(incompleteLines []IncompleteLine) {
	openToCloseMap := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
		'<': '>',
	}

	var allScores []int
	for _, line := range incompleteLines {
		var closingChars []rune
		for line.stack.Len() > 0 {
			closingChars = append(closingChars, openToCloseMap[line.stack.Pop()])
		}

		score := 0
		for _, cc := range closingChars {
			score *= 5
			score += part2Scores[cc]
		}

		allScores = append(allScores, score)
	}

	sort.Ints(allScores)

	fmt.Println("part 2:", allScores[len(allScores)/2])
}

func calculateBadCharacterScore(rs []rune) int {
	var score int
	for _, r := range rs {
		score += part1Scores[r]
	}
	return score
}

func checkLine(line string) (bool, rune, *Stack) {
	stack := &Stack{}
	for _, c := range line {
		if isOpeningBracket(c) {
			stack.Push(c)
			continue
		}

		// Closing bracket but no opening brackets?
		if stack.Len() == 0 {
			return false, c, stack
		}

		switch {
		case c == '}':
			popped := stack.Pop()
			if popped != '{' {
				return false, c, stack
			}
		case c == ']':
			popped := stack.Pop()
			if popped != '[' {
				return false, c, stack
			}
		case c == ')':
			popped := stack.Pop()
			if popped != '(' {
				return false, c, stack
			}
		case c == '>':
			popped := stack.Pop()
			if popped != '<' {
				return false, c, stack
			}
		}
	}

	return stack.Len() == 0, 0, stack
}

type Stack struct {
	values []rune
}

func (b *Stack) Push(value rune) {
	b.values = append(b.values, value)
}

func (b *Stack) Pop() rune {
	if len(b.values) == 0 {
		return 0
	}

	popped := b.values[len(b.values)-1]
	b.values = b.values[:len(b.values)-1]
	return popped
}

func (b *Stack) Len() int {
	return len(b.values)
}

type IncompleteLine struct {
	line  string
	stack *Stack
}

var openingBrackets = map[rune]bool{
	'{': true,
	'(': true,
	'[': true,
	'<': true,
}

func isOpeningBracket(c rune) bool {
	_, ok := openingBrackets[c]
	return ok
}
