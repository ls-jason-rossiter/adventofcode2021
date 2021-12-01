package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/ls-jason-rossiter/adventofcode2021/trash"
)

func main() {
	f := trash.MustLoadFile("input.txt")

	depthIncreaseCount := -1
	currentDepth := 0

	var nums []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		depth, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}

		if depth > currentDepth {
			depthIncreaseCount++
		}

		currentDepth = depth

		nums = append(nums, depth)
	}

	fmt.Printf("part 1: depth increases %v times\n", depthIncreaseCount)

	depthWindowIncreaseCount := -1
	currentDepthWindow := 0

	for i := range nums {
		if i+2 > len(nums)-1 {
			break
		}

		depthWindow := nums[i] + nums[i+1] + nums[i+2]

		if depthWindow > currentDepthWindow {
			depthWindowIncreaseCount++
		}

		currentDepthWindow = depthWindow
	}

	fmt.Printf("part 2: depth window increases %v times\n", depthWindowIncreaseCount)
}
