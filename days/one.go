package days

import (
	"aoc2023/utilities"
	"fmt"
	"strconv"
	"unicode"
)

func Run() {
	lines, err := utilities.ReadLines("data/input1a.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%d %s", len(lines), "lines loaded from file"))

	total := 0
	for _, line := range lines {
		number := ExtractNumbers(line)
		num, err := strconv.Atoi(number)

		if err != nil {
			panic(err)
		}

		total += num
	}

	fmt.Println(fmt.Sprintf("Total is %d", total))
}

func ExtractNumbers(line string) string {
	var first rune
	var last rune

	for _, r := range line {
		if unicode.IsNumber(r) {
			if first == 0 {
				first = r
				last = r
			} else {
				last = r
			}
		}
	}

	return fmt.Sprintf("%c%c", first, last)
}
