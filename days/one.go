package days

import (
	"aoc2023/utilities"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const maxWindow = 5

func Run() {
	lines, err := utilities.ReadLines("data/input1a.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%d %s", len(lines), "lines loaded from file"))

	total := 0
	for _, line := range lines {
		number := ExtractNumbersPartTwo(line)
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

func ExtractNumbersPartTwo(line string) string {
	numbers := map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

	var first rune
	var last rune

	for i := 0; i < len(line); i++ {
		code := rune(line[i])

		if unicode.IsNumber(code) {
			if first == 0 {
				first = code
				last = code
			} else {
				last = code
			}
		} else {
			window := min(len(line)-i, maxWindow)
			slice := line[i : i+window]

			for key, value := range numbers {
				if strings.HasPrefix(slice, key) {
					if first == 0 {
						first = value
						last = value
					} else {
						last = value
					}
				}
			}
		}
	}

	return fmt.Sprintf("%c%c", first, last)
}
