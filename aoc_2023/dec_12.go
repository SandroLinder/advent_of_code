package main

import (
	"strings"
)

func isValidAssignment(assigned string, numbers []int) bool {
	i := 0

	for _, number := range numbers {
		for i < len(assigned) && assigned[i] != '#' {
			i++
		}

		numberOfSprings := 0

		for i < len(assigned) && assigned[i] == '#' {
			numberOfSprings++
			i++
		}

		if numberOfSprings != number {
			return false
		}
	}

	return true
}

func solveProblem12Part1() {
	input := ReadFileIntoArray("resources/dec_12/example.txt")
	//input := ReadFileIntoArray("resources/dec_12/input.txt")

	for _, line := range input {
		split := strings.Split(line, " ")
		StringToNumberOfIntegersBySymbol(split[1], ",")

	}
}
