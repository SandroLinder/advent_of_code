package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveProblem01() {
	var input = ReadFileIntoArray("resources/dec_01/ex_01.txt")

	solveEx01(input)
	solveEx02(input)
}

func solveEx02(input []string) {
	sum := 0
	for _, line := range input {
		var digits []int

		for i := 0; i < len(line); i++ {
			if isADigit(line[i]) {
				number, _ := strconv.Atoi(string(line[i]))
				digits = append(digits, number)
			}

			number, ok := test(line[i:])

			if ok {
				digits = append(digits, number)
			}
		}

		number := strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1])
		numberAsInt, _ := strconv.Atoi(number)

		sum += numberAsInt
	}

	fmt.Println(sum)
}

func test(str string) (int, bool) {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	for k, _ := range numbers {
		if strings.HasPrefix(str, k) {
			return numbers[k], true
		}
	}

	return -1, false
}

func solveEx01(input []string) {
	sum := 0

	for _, row := range input {
		start := 0
		end := len(row) - 1

		for start < end {
			if isADigit(row[start]) {
				break
			}

			start++
		}
		firstDigit := row[start]

		for end >= 0 {
			if isADigit(row[end]) {
				break
			}

			end--
		}

		secondDigit := row[end]

		number := string(firstDigit) + string(secondDigit)
		numberAsInt, _ := strconv.Atoi(number)

		sum += numberAsInt
	}

	fmt.Println(sum)
}

func isADigit(u uint8) bool {
	return u >= 48 && u <= 57
}
