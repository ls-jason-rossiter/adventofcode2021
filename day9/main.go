package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")
	defer f.Close()

	input := trash.ReaderToStrings(f)
	hm := &HeightMap{
		Map:     make([][]int, len(input)),
		Checked: make(map[Point]bool),
	}
	for i, s := range input {
		hm.Map[i] = make([]int, len(s))

		split := strings.Split(s, "")
		for j, v := range split {
			hm.Map[i][j] = trash.MustParseIntBase10(v)
		}
	}

	part1(hm, hm.LowestPoints())
	part2(hm, hm.LowestPoints())
}

func part1(hm *HeightMap, lowestPoints []Point) {
	riskSum := 0
	for _, point := range lowestPoints {
		riskSum += hm.Map[point.X][point.Y] + 1
	}
	fmt.Println("part 1:", riskSum)
}

func part2(hm *HeightMap, lowestPoints []Point) {
	var basinSizes []int
	for _, point := range lowestPoints {
		hm.Checked = make(map[Point]bool)
		basinSizes = append(basinSizes, hm.BasinSize(point))
	}

	sort.Ints(basinSizes)

	totalSize := 1
	for _, i := range basinSizes[len(basinSizes)-3:] {
		totalSize *= i
	}

	fmt.Println("part 2:", totalSize)
}

type Point struct {
	X, Y int
}

func (p Point) AdjacentPoints() []Point {
	return []Point{
		{p.X, p.Y - 1}, // above
		{p.X + 1, p.Y}, // right
		{p.X, p.Y + 1}, // below
		{p.X - 1, p.Y}, // left
	}
}

type HeightMap struct {
	Map     [][]int
	Checked map[Point]bool
}

// FilterPoints returns a list of Points that exist in the map, and are valid
// basin values (not 9).
func (hm *HeightMap) FilterPoints(points []Point) []Point {
	var validPoints []Point
	minX, minY := 0, 0
	maxX := len(hm.Map) - 1
	maxY := len(hm.Map[0]) - 1

	for _, point := range points {
		withinMapBounds := point.X >= minX && point.X <= maxX && point.Y >= minY && point.Y <= maxY
		if withinMapBounds {
			pVal := hm.Map[point.X][point.Y]
			if pVal < 9 {
				validPoints = append(validPoints, point)
			}
		}
	}
	return validPoints
}

// BasinSize returns the size of a basin from a given point.
func (hm *HeightMap) BasinSize(targetPoint Point) int {
	if _, ok := hm.Checked[targetPoint]; ok {
		return 0
	}
	hm.Checked[targetPoint] = true

	adjacentPoints := targetPoint.AdjacentPoints()
	points := hm.FilterPoints(adjacentPoints)
	if len(points) == 0 {
		return 0
	}

	size := 0
	for _, point := range points {
		pVal := hm.Map[point.X][point.Y]
		tVal := hm.Map[targetPoint.X][targetPoint.Y]
		if pVal > tVal {
			size += hm.BasinSize(point)
		}
	}
	return size + 1
}

func (hm HeightMap) LowestPoints() []Point {
	var lowestPoints []Point
	for row, heights := range hm.Map {
		for col, height := range heights {
			// Top row.
			if row == 0 {
				if col == 0 {
					// Top left.
					if check(height, hm.Map[row][col+1], hm.Map[row+1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if col == len(heights)-1 {
					// Top right.
					if check(height, hm.Map[row][col-1], hm.Map[row+1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if check(height, hm.Map[row][col-1], hm.Map[row][col+1], hm.Map[row+1][col]) {
					// Between
					lowestPoints = append(lowestPoints, Point{row, col})
					continue
				}
			} else if row == len(hm.Map)-1 {
				// Bottom row.
				if col == 0 {
					// Bottom left.
					if check(height, hm.Map[row][col+1], hm.Map[row-1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if col == len(heights)-1 {
					// Bottom right.
					if check(height, hm.Map[row][col-1], hm.Map[row-1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if check(height, hm.Map[row][col-1], hm.Map[row][col+1], hm.Map[row-1][col]) {
					// Between
					lowestPoints = append(lowestPoints, Point{row, col})
					continue
				}
			} else {
				// The rest.
				if col == 0 {
					// Left side.
					if check(height, hm.Map[row][col+1], hm.Map[row+1][col], hm.Map[row-1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if col == len(heights)-1 {
					// Right side.
					if check(height, hm.Map[row][col-1], hm.Map[row+1][col], hm.Map[row-1][col]) {
						lowestPoints = append(lowestPoints, Point{row, col})
						continue
					}
				} else if check(height, hm.Map[row][col-1], hm.Map[row][col+1], hm.Map[row+1][col], hm.Map[row-1][col]) {
					// Between
					lowestPoints = append(lowestPoints, Point{row, col})
					continue
				}
			}
		}
	}

	return lowestPoints
}

// Returns true if target is the lowest value.
func check(target int, vs ...int) bool {
	lowestValue := 99999
	for _, v := range vs {
		if v < lowestValue {
			lowestValue = v
		}
	}
	return target < lowestValue
}
