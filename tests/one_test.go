package tests

import (
	"aoc2023/days"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	var tests = []struct {
		input  string
		answer string
	}{
		{"1abc2", "12"},
		{"pqr3stu8vwx", "38"},
		{"a1b2c3d4e5f", "15"},
		{"treb7uchet", "77"},
	}

	for _, tt := range tests {
		testname := tt.input + "," + tt.answer
		t.Run(testname, func(t *testing.T) {
			ans := days.ExtractNumbers(tt.input)
			if ans != tt.answer {
				t.Errorf("got %s, want %s", ans, tt.answer)
			}
		})
	}
}
