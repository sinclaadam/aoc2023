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

	for i := 0; i < len(lines); i++ {
		number := ExtractNumbers(lines[i])
		num, err := strconv.Atoi(number)

		if err != nil {
			panic(err)
		}

		total += num
	}

	fmt.Println(fmt.Sprintf("Total is %d", total))
}

func ExtractNumbers(line string) string {
	var first uint8 = 0
	var last uint8 = 0

	for i := 0; i < len(line); i++ {
		c := line[i]
		if unicode.IsNumber(rune(c)) {
			if first == 0 {
				first = c - '0'
				last = c - '0'
			} else {
				last = c - '0'
			}
		}
	}

	return fmt.Sprintf("%d%d", first, last)
}
