package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "3,4,3,1,2"
	hatchery := parseInput([]string{input})
	expected := Hatchery{0, 1, 1, 2, 1, 0, 0, 0, 0}

	if !reflect.DeepEqual(expected, hatchery) {
		t.Errorf("Expected %v, got %v", expected, hatchery)
	}
}

func TestHatchery_LarvaCountAfter18(t *testing.T) {
	input := "3,4,3,1,2"
	hatchery := parseInput([]string{input})
	count := hatchery.LarvaCountAfter(18)

	if count != 26 {
		t.Errorf("Expected %v, got %v", 26, count)
	}
}

func TestHatchery_LarvaCountAfter80(t *testing.T) {
	input := "3,4,3,1,2"
	hatchery := parseInput([]string{input})
	count := hatchery.LarvaCountAfter(80)

	if count != 5934 {
		t.Errorf("Expected %v, got %v", 5934, count)
	}
}

func BenchmarkHatchery_Progress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hatchery := Hatchery{0, 1, 1, 2, 1, 0, 0, 0, 0}
		hatchery.LarvaCountAfter(18)
	}
}
