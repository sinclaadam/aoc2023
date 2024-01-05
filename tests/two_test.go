package tests

import (
	"aoc2023/days"
	"fmt"
	"strconv"
	"testing"
)

func TestTwoPartOne(t *testing.T) {
	var tests = []struct {
		input  string
		answer bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for _, tt := range tests {
		testname := tt.input + "," + strconv.FormatBool(tt.answer)
		t.Run(testname, func(t *testing.T) {
			ans := days.IsPossibleGame(tt.input)
			if ans != tt.answer {
				t.Errorf("got %t, want %t", ans, tt.answer)
			}
		})
	}
}

func TestTwoPartTwo(t *testing.T) {
	var tests = []struct {
		input  string
		answer int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%d", tt.input, tt.answer)
		t.Run(testname, func(t *testing.T) {
			ans := days.PowerSet(tt.input)
			if ans != tt.answer {
				t.Errorf("got %d, want %d", ans, tt.answer)
			}
		})
	}
}
