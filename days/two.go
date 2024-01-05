package days

import (
	"aoc2023/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

var (
	redReg   = regexp.MustCompile("[0-9]+ red")
	greenReg = regexp.MustCompile("[0-9]+ green")
	blueReg  = regexp.MustCompile("[0-9]+ blue")
)

func RunTwoPartOne() {
	lines, err := utilities.ReadLines("data/input2.txt")

	if err != nil {
		panic(err)
	}

	total := 0

	for i, line := range lines {
		if IsPossibleGame(line) {
			total += i + 1
		}
	}

	fmt.Println(total)
}

func RunTwoPartTwo() {
	lines, err := utilities.ReadLines("data/input2.txt")

	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range lines {
		total += PowerSet(line)
	}

	fmt.Println(total)
}

func IsPossibleGame(line string) bool {
	reds := redReg.FindAllString(line, -1)
	blues := blueReg.FindAllString(line, -1)
	greens := greenReg.FindAllString(line, -1)

	for _, match := range reds {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		if number > maxRed {
			return false
		}
	}

	for _, match := range blues {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		if number > maxBlue {
			return false
		}
	}

	for _, match := range greens {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		if number > maxGreen {
			return false
		}
	}

	return true
}

func PowerSet(line string) int {
	reds := redReg.FindAllString(line, -1)
	blues := blueReg.FindAllString(line, -1)
	greens := greenReg.FindAllString(line, -1)

	redHighest := 0
	blueHighest := 0
	greenHighest := 0

	for _, match := range reds {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		redHighest = max(redHighest, number)
	}

	for _, match := range blues {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		blueHighest = max(blueHighest, number)
	}

	for _, match := range greens {
		number, _ := strconv.Atoi(strings.Split(match, " ")[0])

		greenHighest = max(greenHighest, number)
	}

	return redHighest * blueHighest * greenHighest
}
