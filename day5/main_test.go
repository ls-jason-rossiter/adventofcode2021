package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	expected := Lines{
		{
			From: Point{X: 0, Y: 9},
			To:   Point{X: 5, Y: 9},
		},
		{
			From: Point{X: 8, Y: 0},
			To:   Point{X: 0, Y: 8},
		},
		{
			From: Point{X: 9, Y: 4},
			To:   Point{X: 3, Y: 4},
		},
		{
			From: Point{X: 2, Y: 2},
			To:   Point{X: 2, Y: 1},
		},
		{
			From: Point{X: 7, Y: 0},
			To:   Point{X: 7, Y: 4},
		},
		{
			From: Point{X: 6, Y: 4},
			To:   Point{X: 2, Y: 0},
		},
		{
			From: Point{X: 0, Y: 9},
			To:   Point{X: 2, Y: 9},
		},
		{
			From: Point{X: 3, Y: 4},
			To:   Point{X: 1, Y: 4},
		},
		{
			From: Point{X: 0, Y: 0},
			To:   Point{X: 8, Y: 8},
		},
		{
			From: Point{X: 5, Y: 5},
			To:   Point{X: 8, Y: 2},
		},
	}

	lines := parseInput(input)

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(lines))
	}

	if !reflect.DeepEqual(expected, lines) {
		t.Errorf("Expected \n%vgot \n%v", expected, lines)
	}
}

func TestLines_PointOverlapCount(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	lines := parseInput(input)

	pointOverlapCount := lines.PointOverlapCount(false)
	if pointOverlapCount != 5 {
		t.Errorf("Expected 5 overlaps, got %d", pointOverlapCount)
	}
}

func TestLines_PointOverlapCountWithDiagonals(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	lines := parseInput(input)

	pointOverlapCount := lines.PointOverlapCount(true)
	if pointOverlapCount != 12 {
		t.Errorf("Expected 12 overlaps, got %d", pointOverlapCount)
	}
}
