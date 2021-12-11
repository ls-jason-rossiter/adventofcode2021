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

	ss := trash.ReaderToStrings(f)
	part1(parseInput(ss))
	part2(parseInput(ss))
}

func parseInput(ss []string) *Grid {
	g := &Grid{
		grid: [10][10]*Octopus{},
	}
	for i, s := range ss {
		octopuses := strings.Split(s, "")
		g.grid[i] = [10]*Octopus{}
		for j, octopus := range octopuses {
			g.grid[i][j] = &Octopus{
				Energy: trash.MustParseIntBase10(octopus),
			}
		}
	}
	return g
}

func part1(grid *Grid) {
	sumFlashes := 0
	for i := 0; i < 100; i++ {
		sumFlashes += grid.Step()
	}
	fmt.Println("part 1:", sumFlashes)
}

func part2(grid *Grid) {
	flashes := 0
	steps := 0
	for flashes != 100 {
		flashes = grid.Step()
		steps++
	}
	fmt.Println("part 2:", steps)
}

type Grid struct {
	grid [10][10]*Octopus
}

func (g *Grid) String() string {
	var grid strings.Builder
	for _, oct := range g.grid {
		var row strings.Builder
		for _, octopus := range oct {
			energy := strconv.Itoa(octopus.Energy)
			if octopus.Energy > 9 {
				energy = "X"
			}
			fmt.Fprintf(&row, "%s", energy)
		}
		fmt.Fprintf(&grid, "%s\n", row.String())
	}
	return grid.String()
}

func (g *Grid) Step() int {
	flashCount := 0

	// Increase each octopus' energy level by one.
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g.grid[i][j].Energy++
		}
	}

	// When an octopus has more than 9 energy it flashes.
	var octopusesToFlash []Point
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.grid[i][j].CanFlash() {
				octopusesToFlash = append(octopusesToFlash, Point{i, j})
			}
		}
	}
	for _, p := range octopusesToFlash {
		flashCount += g.HandleFlash(p)
	}

	// Reset flashed octopuses back to energy zero.
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.grid[i][j].Flashed {
				g.grid[i][j].Energy = 0
				g.grid[i][j].Flashed = false
			}
		}
	}

	return flashCount
}

func (g *Grid) HandleFlash(p Point) int {
	octopus := g.grid[p.X][p.Y]
	if !octopus.CanFlash() {
		return 0
	}

	flashCount := 1
	octopus.Flashed = true

	// Each flash increases its neighbour's energy by 1.
	for _, n := range g.FilterPoints(p.Neighbours()) {
		g.grid[n.X][n.Y].Energy++

		if g.grid[n.X][n.Y].CanFlash() {
			flashCount += g.HandleFlash(n)
		}
	}

	return flashCount
}

func (g *Grid) FilterPoints(ps []Point) []Point {
	minX, minY := 0, 0
	maxX, maxY := 9, 9

	var validPoints []Point
	for _, point := range ps {
		withinMapBounds := point.X >= minX && point.X <= maxX && point.Y >= minY && point.Y <= maxY
		if withinMapBounds {
			validPoints = append(validPoints, point)
		}
	}
	return validPoints
}

type Octopus struct {
	Energy  int
	Flashed bool
}

func (o *Octopus) CanFlash() bool {
	return !o.Flashed && o.Energy > 9
}

type Point struct {
	X int
	Y int
}

func (p Point) Neighbours() []Point {
	return []Point{
		{p.X - 1, p.Y - 1}, // top left
		{p.X, p.Y - 1},     // above
		{p.X + 1, p.Y - 1}, // top right
		{p.X - 1, p.Y},     // left
		{p.X + 1, p.Y},     // right
		{p.X - 1, p.Y + 1}, // bottom left
		{p.X, p.Y + 1},     // below
		{p.X + 1, p.Y + 1}, // bottom right
	}
}
