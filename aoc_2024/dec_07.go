package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	expectedResult int
	values         []int
}

func parseEquations(input []string) []Equation {
	var equations []Equation

	for _, line := range input {
		eq := strings.Split(line, ":")

		expected, _ := strconv.Atoi(eq[0])
		values := StringToNumberOfIntegersBySymbol(eq[1], " ")

		equations = append(equations, Equation{expectedResult: expected, values: values})
	}

	return equations
}

func solveExercise1Day07(input []string) {
	equations := parseEquations(input)

	sum := 0
	for _, equation := range equations {
		valid := checkEquations(equation)
		fmt.Printf("Equation: %+v -> valid: %t\n", equation, valid)
		if valid {
			sum += equation.expectedResult
		}
	}

	fmt.Println(sum)
}

var operations = []string{"+", "*", "||"}

func walkEquation(values []int, goal int) bool {
	if len(values) == 1 {
		if goal == values[0] {
			return true
		} else {
			return false
		}
	}

	for _, op := range operations {
		if op == "+" {
			next := []int{values[0] + values[1]}
			next = append(next, values[2:]...)
			if walkEquation(next, goal) {
				return true
			}
		} else if op == "*" {
			next := []int{values[0] * values[1]}
			next = append(next, values[2:]...)
			if walkEquation(next, goal) {
				return true
			}
		} else {
			newValue, _ := strconv.Atoi(strconv.Itoa(values[0]) + strconv.Itoa(values[1]))
			next := []int{newValue}
			next = append(next, values[2:]...)
			if walkEquation(next, goal) {
				return true
			}
		}
	}
	return false
}

func checkEquations(equation Equation) bool {
	valid := walkEquation(equation.values, equation.expectedResult)
	return valid
}

func solveExercise2Day07(input []string) {
	equations := parseEquations(input)

	sum := 0
	for _, equation := range equations {
		valid := checkEquations(equation)
		fmt.Printf("Equation: %+v -> valid: %t\n", equation, valid)
		if valid {
			sum += equation.expectedResult
		}
	}

	fmt.Println(sum)
}
