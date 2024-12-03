package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func solveExercise1Day03(input string) {
	reg := regexp.MustCompile(`mul\((\d*,\d*)\)`)

	result := reg.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, r := range result {
		numbers := StringToNumberOfIntegersBySymbol(r[1], ",")
		sum += numbers[0] * numbers[1]
	}

	fmt.Println(sum)
}

func solveExercise2Day03(input string) {
	currentIndex := 0
	enabled := true

	var sum int64
	for currentIndex < len(input) {
		if input[currentIndex] == 'm' {
			if input[currentIndex+1] == 'u' && input[currentIndex+2] == 'l' && input[currentIndex+3] == '(' {
				currentIndex += 4
				firstNumber := strings.Builder{}
				secondNumber := strings.Builder{}
				current := input[currentIndex]
				for current != ',' && IsDigit(current) {
					firstNumber.WriteString(string(current))
					currentIndex++
					current = input[currentIndex]
				}

				if input[currentIndex] == ',' {
					currentIndex++
					for input[currentIndex] != ')' {
						secondNumber.WriteString(string(input[currentIndex]))
						currentIndex++
					}
				} else {
					currentIndex++
					continue
				}

				if input[currentIndex] != ')' {
					continue
				}

				if enabled {
					first, _ := strconv.Atoi(firstNumber.String())
					second, _ := strconv.Atoi(secondNumber.String())

					fmt.Printf("%d,%d\n", first, second)
					sum += int64(first * second)
				}

			}
		} else if input[currentIndex] == 'd' {
			if input[currentIndex+1] == 'o' {
				if input[currentIndex+2] == '(' && input[currentIndex+3] == ')' {
					enabled = true
					currentIndex += 4
					continue
				}

				if input[currentIndex+2] == 'n' && input[currentIndex+3] == '\'' && input[currentIndex+4] == 't' {
					enabled = false
					currentIndex += 7
					continue
				}
			}
		}
		currentIndex++
	}

	fmt.Println(sum)
}
