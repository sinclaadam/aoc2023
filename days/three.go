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

func RunThreePartTwo() {
	lines, err := utilities.ReadLines("data/input3.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(CalculateGearRatioSum(lines))
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

func CalculateGearRatioSum(lines []string) int {
	type coordinate struct {
		row int
		col int
	}

	type gear struct {
		partOne int
		partTwo int
	}

	sum := 0
	gears := make(map[coordinate]*gear)

	for i := 0; i < len(lines); i++ {
		number := ""
		gearCoordinate := coordinate{row: -1, col: -1}

		for j := 0; j < len(lines[i]); j++ {
			if unicode.IsNumber(rune(lines[i][j])) {
				number += string(lines[i][j])
				if gearCoordinate.row == -1 {
					gearCoordinate.row, gearCoordinate.col = isNextToGear(lines, i, j)
				}
			} else {
				if number != "" && gearCoordinate.row != -1 {
					num, err := strconv.Atoi(number)

					if err != nil {
						panic(err)
					}

					if entry, ok := gears[gearCoordinate]; ok {
						entry.partTwo = num
						gears[gearCoordinate] = entry
					} else {
						gears[gearCoordinate] = &gear{partOne: num}
					}
				}

				number = ""
				gearCoordinate = coordinate{row: -1, col: -1}
			}

			// edge case where we might be at the end of the line
			if j == len(lines[i])-1 && number != "" && gearCoordinate.row != -1 {
				num, err := strconv.Atoi(number)

				if err != nil {
					panic(err)
				}

				if entry, ok := gears[gearCoordinate]; ok {
					entry.partTwo = num
					gears[gearCoordinate] = entry
				} else {
					gears[gearCoordinate] = &gear{partOne: num}
				}
			}
		}
	}

	for _, v := range gears {
		if v.partOne != 0 && v.partTwo != 0 {
			sum += v.partOne * v.partTwo
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

func isNextToGear(lines []string, row int, col int) (r int, c int) {
	for i := max(0, row-1); i < min(row+2, len(lines)); i++ {
		for j := max(0, col-1); j < min(col+2, len(lines[i])); j++ {
			if rune(lines[i][j]) == '*' {
				return i, j
			}
		}
	}

	return -1, -1
}
