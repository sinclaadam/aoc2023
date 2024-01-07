package days

import (
	"aoc2023/utilities"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunFourPartOne() {
	lines, err := utilities.ReadLines("data/input4.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(CalculateWinningPoints(lines))
}

func RunFourPartTwo() {
	lines, err := utilities.ReadLines("data/input4.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(CalculateTotalScratchyCards(lines))
}

func CalculateWinningPoints(lines []string) int {
	totalPoints := 0

	for _, line := range lines {
		winningNumbers := getWinningNumbers(line)
		scratchyNumbers := getScratchyNumbers(line)
		linepoints := 0

		for _, num := range scratchyNumbers {
			if slices.Contains(winningNumbers, num) {
				linepoints = incrementPoints(linepoints)
			}
		}

		totalPoints += linepoints
	}

	return totalPoints
}

func CalculateTotalScratchyCards(lines []string) int {
	copies := make(map[int]int)
	for i, _ := range lines {
		copies[i] = 1
	}

	for i, line := range lines {
		winningNumbers := getWinningNumbers(line)
		scratchyNumbers := getScratchyNumbers(line)

		matchingNumbers := 0

		for _, num := range scratchyNumbers {
			if slices.Contains(winningNumbers, num) {
				matchingNumbers++
			}
		}

		for j := i + 1; matchingNumbers >= 1; j++ {
			copies[j] = copies[j] + copies[i]
			matchingNumbers--
		}
	}

	sum := 0
	for _, v := range copies {
		sum += v
	}

	return sum
}

func getWinningNumbers(line string) []int {
	var numbers []int

	winningNumbers := line[strings.Index(line, ":")+1 : strings.Index(line, "|")]
	winningNumbers = strings.Trim(winningNumbers, " ")

	for _, number := range strings.Split(winningNumbers, " ") {
		num, err := strconv.Atoi(number)

		if err == nil { // sometimes we get spaces but they will fail and we can ignore those.
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func getScratchyNumbers(line string) []int {
	var numbers []int

	scratchyNumbers := line[strings.Index(line, "|")+1:]
	scratchyNumbers = strings.Trim(scratchyNumbers, " ")

	for _, number := range strings.Split(scratchyNumbers, " ") {
		num, err := strconv.Atoi(number)

		if err == nil { // sometimes we get spaces but they will fail and we can ignore those.
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func incrementPoints(value int) int {
	switch value {
	case 0:
		return 1
	default:
		return value * 2
	}
}
