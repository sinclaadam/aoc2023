package tests

import (
	"aoc2023/days"
	"testing"
)

func TestThreePartOne(t *testing.T) {
	input := []string{"467..114..", "...*......", "..35..633.", "......#...", "617*......", ".....+.58.", "..592.....", "......755.", "...$.*....", ".664.598.."}
	want := 4361
	ans := days.CalculatePartNumberSum(input)

	if want != ans {
		t.Fatalf(`sum incorrect, want %d, got %d`, want, ans)
	}
}
