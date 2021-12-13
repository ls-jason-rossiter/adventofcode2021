package main

import (
	"fmt"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	paper, folds := parseInput(trash.ReaderToStrings(f))

	part1(paper, folds)
	part2(paper, folds)
}

func parseInput(ss []string) (Paper, []Fold) {
	var dots []Point
	var folds []Fold
	for _, s := range ss {
		if s == "" {
			continue
		}
		if strings.Contains(s, ",") {
			parts := strings.Split(s, ",")
			dots = append(dots, Point{
				X: trash.MustParseIntBase10(parts[0]),
				Y: trash.MustParseIntBase10(parts[1]),
			})
		}

		if strings.Contains(s, "fold") {
			parts := strings.Fields(s)
			fold := strings.Split(parts[2], "=")
			folds = append(folds, Fold{
				Direction: fold[0],
				Point:     trash.MustParseIntBase10(fold[1]),
			})
		}
	}

	return Paper{Dots: dots}, folds
}

func part1(paper Paper, folds []Fold) {
	paper = paper.Fold(folds[0])
	fmt.Println("part 1:", paper.DotCount())
}

func part2(paper Paper, folds []Fold) {
	foldedPaper := paper
	for _, fold := range folds {
		foldedPaper = foldedPaper.Fold(fold)
	}

	fmt.Println("part 2:")
	foldedPaper.Print()
}

type Paper struct {
	Dots []Point
}

func (p Paper) Print() {
	maxX, maxY := MaxValues(p.Dots)
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if DotsContains(p.Dots, Point{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p Paper) DotCount() int {
	return len(p.Dots)
}

func (p Paper) Fold(fold Fold) Paper {
	switch fold.Direction {
	case "x":
		return p.FoldX(fold.Point)
	case "y":
		return p.FoldY(fold.Point)
	default:
		panic("uhh")
	}
}

func (p Paper) FoldX(point int) Paper {
	// Find all dots where X > point
	var newDots []Point
	var movedDots []Point
	for _, dot := range p.Dots {
		if dot.X > point {
			movedDots = append(movedDots, dot)
		} else {
			newDots = append(newDots, dot)
		}
	}

	// Inverse the dot coordinates based on MaxY.
	for _, dot := range movedDots {
		dot.X = point*2 - dot.X
		if !DotsContains(newDots, dot) {
			newDots = append(newDots, dot)
		}
	}

	return Paper{Dots: newDots}
}

func (p Paper) FoldY(point int) Paper {
	// Find all dots where Y > point
	var newDots []Point
	var movedDots []Point
	for _, dot := range p.Dots {
		if dot.Y > point {
			movedDots = append(movedDots, dot)
		} else {
			newDots = append(newDots, dot)
		}
	}

	// Inverse the dot coordinates based on MaxY.
	for _, dot := range movedDots {
		dot.Y = point*2 - dot.Y
		if !DotsContains(newDots, dot) {
			newDots = append(newDots, dot)
		}
	}

	return Paper{Dots: newDots}
}

type Point struct {
	X, Y int
}

type Fold struct {
	Direction string
	Point     int
}

func MaxValues(points []Point) (int, int) {
	maxX, maxY := 0, 0
	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxX, maxY
}

func DotsContains(dots []Point, p Point) bool {
	for _, dot := range dots {
		if dot == p {
			return true
		}
	}
	return false
}
