package main

import (
	"testing"
)

func TestConstantFuelCost(t *testing.T) {
	input := []string{"16,1,2,0,4,2,7,1,2,14"}
	crabs := parseInput(input)

	fuelCost := constantFuelCost(2, crabs)
	if fuelCost != 37 {
		t.Errorf("Expected fuel cost to be 37, got %d", fuelCost)
	}
}

func TestParseInput2(t *testing.T) {
	input := []string{"16,1,2,0,4,2,7,1,2,14"}
	crabs := parseInput(input)

	fuelCost := scalingFuelCost(5, crabs)
	if fuelCost != 168 {
		t.Errorf("Expected fuel cost to be 168, got %d", fuelCost)
	}
}
