package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	bingo, numbersToMark := parseInput(trash.ReaderToStrings(f))
	winningBoards := bingo.Play(numbersToMark)
	if winningBoards == nil || len(winningBoards) == 0 {
		panic("no winning boards")
	}

	part1(winningBoards)
	part2(winningBoards)
}

func part1(winningBoards []*BingoBoard) {
	winningBoard := winningBoards[0]
	fmt.Println("part 1:", winningBoard.Score())
}

func part2(winningBoards []*BingoBoard) {
	winningBoard := winningBoards[len(winningBoards)-1]
	fmt.Println("part 2:", winningBoard.Score())
}

// Returns a BingoBoards and list of numbers to mark.
func parseInput(ss []string) (BingoBoards, []int) {
	if len(ss) == 0 {
		return nil, nil
	}

	var numbersToMark []int
	for _, s := range strings.Split(ss[0], ",") {
		num := mustParseIntBase10(s)
		numbersToMark = append(numbersToMark, num)
	}

	var boards []*BingoBoard

	board := &BingoBoard{}
	var boardRow int
	for _, s := range ss[2:] {
		if s == "" {
			boards = append(boards, board)
			board = &BingoBoard{}
			boardRow = 0
			continue
		}

		board.Board[boardRow] = [5]*BingoNumber{}

		for i, n := range strings.Fields(s) {
			num := mustParseIntBase10(n)
			board.Board[boardRow][i] = &BingoNumber{Number: num}
		}
		boardRow++
	}

	// Append final board.
	boards = append(boards, board)

	return boards, numbersToMark
}

func mustParseIntBase10(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

type BingoBoards []*BingoBoard

// Play returns a list of boards in winning order.
func (bg BingoBoards) Play(numbersToMark []int) []*BingoBoard {
	var winningBoards []*BingoBoard
	for _, num := range numbersToMark {
		for _, board := range bg {
			if board.MarkNumber(num) {
				winningBoards = append(winningBoards, board)
			}
		}
	}
	return winningBoards
}

// BingoNumber is a number on a bingo board.
type BingoNumber struct {
	Number int
	Marked bool
}

// BingoBoard holds the state of a bingo board.
type BingoBoard struct {
	Finished      bool
	WinningNumber int
	Board         [5][5]*BingoNumber
}

func (bb *BingoBoard) Score() int {
	sumOfUnmarkedNumbers := 0
	for _, n := range bb.UnmarkedNumbers() {
		sumOfUnmarkedNumbers += n
	}
	return sumOfUnmarkedNumbers * bb.WinningNumber
}

// MarkNumber returns true if the given marked number triggers a win.
func (bb *BingoBoard) MarkNumber(num int) bool {
	if bb.Finished {
		return false
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if bb.Board[i][j].Number == num {
				bb.Board[i][j].Marked = true
			}
		}
	}

	if bb.HasWon() {
		bb.Finished = true
		bb.WinningNumber = num
		return true
	}
	return false
}

func (bb *BingoBoard) UnmarkedNumbers() []int {
	var nums []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !bb.Board[i][j].Marked {
				nums = append(nums, bb.Board[i][j].Number)
			}
		}
	}
	return nums
}

func (bb *BingoBoard) HasWon() bool {
	return bb.CheckRows() || bb.CheckCols()
}

func (bb *BingoBoard) CheckRows() bool {
	for i := 0; i < 5; i++ {
		markedCount := 0
		for _, bn := range bb.Board[i] {
			if bn.Marked {
				markedCount++
			}
		}
		if markedCount == 5 {
			return true
		}
	}
	return false
}

func (bb *BingoBoard) CheckCols() bool {
	for i := 0; i < 5; i++ {
		markedCount := 0
		for _, bn := range bb.Board {
			if bn[i].Marked {
				markedCount++
			}
		}
		if markedCount == 5 {
			return true
		}
	}
	return false
}
