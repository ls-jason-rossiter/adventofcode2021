package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	lines := parseInput(trash.ReaderToStrings(f))

	part1(lines)
	part2(lines)
}

func part1(lines Lines) {
	fmt.Println("part 1:", lines.PointOverlapCount(false))
}

func part2(lines Lines) {
	fmt.Println("part 2:", lines.PointOverlapCount(true))
}

type Point struct {
	X, Y int
}

type Line struct {
	From, To Point
}

type Lines []Line

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	}
	return 0
}

func (l Lines) PointOverlapCount(includeDiagonals bool) int {
	points := make(map[Point]int)

	for _, line := range l {
		deltaX := line.To.X - line.From.X
		deltaY := line.To.Y - line.From.Y

		isSameRow := deltaX == 0
		isSameCol := deltaY == 0
		isDiagonal := math.Abs(float64(deltaX)) == math.Abs(float64(deltaY))

		if isDiagonal && !includeDiagonals {
			continue
		}

		var direction Point
		if !isSameRow && !isSameCol && !isDiagonal {
			continue
		}

		direction = Point{Sgn(deltaX), Sgn(deltaY)}

		currentPoint := line.From
		points[currentPoint]++

		for currentPoint != line.To {
			nextPoint := Point{currentPoint.X + direction.X, currentPoint.Y + direction.Y}
			points[nextPoint]++
			currentPoint = nextPoint
		}
	}

	overlapCount := 0
	for _, v := range points {
		if v > 1 {
			overlapCount++
		}
	}
	return overlapCount
}

func parseInput(ss []string) Lines {
	var lines Lines

	for _, s := range ss {
		fields := strings.Fields(s)
		from := strings.Split(fields[0], ",")
		to := strings.Split(fields[2], ",")

		lines = append(lines, Line{
			From: Point{mustParseIntBase10(from[0]), mustParseIntBase10(from[1])},
			To:   Point{mustParseIntBase10(to[0]), mustParseIntBase10(to[1])},
		})
	}

	return lines
}

func mustParseIntBase10(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
