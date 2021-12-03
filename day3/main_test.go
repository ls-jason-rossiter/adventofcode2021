package main

import (
	"testing"
)

func TestGammaRate(t *testing.T) {
	nums := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	rate := gammaRate(nums)
	if rate != 22 {
		t.Errorf("Expected %d, got %v", 22, rate)
	}
}

func TestEpsilonRate(t *testing.T) {
	nums := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	rate := epsilonRate(nums)
	if rate != 9 {
		t.Errorf("Expected %d, got %v", 9, rate)
	}
}

func TestOxygenGeneratorRating(t *testing.T) {
	nums := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	rating := oxygenGeneratorRating(nums)
	if rating != 23 {
		t.Errorf("Expected %d, got %v", 23, rating)
	}
}

func TestCo2ScrubberRating(t *testing.T) {
	nums := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	rating := co2ScrubberRating(nums)
	if rating != 10 {
		t.Errorf("Expected %d, got %v", 10, rating)
	}
}
