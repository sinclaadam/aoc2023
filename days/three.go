package days

import (
	"aoc2023/utilities"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

var symbols = []rune{'@', '#', '$', '%', '&', '*', '/', '+', '=', '-'} // potentially just not a number and not a '.'

func RunThreePartOne() {
	lines, err := utilities.ReadLines("data/input3.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(CalculatePartNumberSum(lines))
}

func CalculatePartNumberSum(lines []string) int {
	sum := 0

	for i := 0; i < len(lines); i++ {
		number := ""
		nextToPart := false

		for j := 0; j < len(lines[i]); j++ {
			if unicode.IsNumber(rune(lines[i][j])) {
				number += string(lines[i][j])
				nextToPart = nextToPart || isNextToPart(lines, i, j)
			} else {
				if number != "" && nextToPart {
					num, err := strconv.Atoi(number)

					if err != nil {
						panic(err)
					}

					sum += num
				}

				number = ""
				nextToPart = false
			}

			// edge case where we might be at the end of the line
			if j == len(lines[i])-1 && number != "" && nextToPart {
				num, err := strconv.Atoi(number)

				if err != nil {
					panic(err)
				}

				sum += num
			}
		}
	}

	return sum
}

func isNextToPart(lines []string, row int, col int) bool {
	for i := max(0, row-1); i < min(row+2, len(lines)); i++ {
		for j := max(0, col-1); j < min(col+2, len(lines[i])); j++ {
			if slices.Contains(symbols, rune(lines[i][j])) {
				return true
			}
		}
	}

	return false
}
