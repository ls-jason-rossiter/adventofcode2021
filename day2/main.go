package main

import (
	"bufio"
	"fmt"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

type instruction struct {
	direction string
	distance  int
}

func main() {
	f := trash.MustLoadFile("input.txt")

	var instructions []instruction

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var ins instruction
		fmt.Sscanf(scanner.Text(), "%s %d", &ins.direction, &ins.distance)
		instructions = append(instructions, ins)
	}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []instruction) {
	x, y := 0, 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			x += instruction.distance
		case "up":
			y -= instruction.distance
		case "down":
			y += instruction.distance
		default:
			panic(":suffering:")
		}
	}

	fmt.Printf("part 1: %d * %d = %d\n", x, y, x*y)
}

func part2(instructions []instruction) {
	x, y, aim := 0, 0, 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			x += instruction.distance
			y += aim * instruction.distance
		case "up":
			aim -= instruction.distance
		case "down":
			aim += instruction.distance
		default:
			panic(":suffering:")
		}
	}

	fmt.Printf("part 2: %d * %d = %d\n", x, y, x*y)
}
