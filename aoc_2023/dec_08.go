package main

import (
	"strings"
)

func main() {
	//var input = ReadFileIntoArray("resources/dec_08/example.txt")
	//var input = ReadFileIntoArray("resources/dec_08/example2.txt")
	//var input = ReadFileIntoArray("resources/dec_08/example3.txt")
	var input = ReadFileIntoArray("resources/dec_08/input.txt")

	points := make(map[string][]string, len(input)-2)

	instructions := input[0]

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "=")

		head := split[0]

		lenBody := len(split[1]) - 1
		body := strings.Split(split[1][2:lenBody], ", ")

		points[strings.TrimSpace(head)] = body
	}

	var currentStates []string

	for k, _ := range points {
		if k[2] == 'A' {
			currentStates = append(currentStates, k)
		}
	}

	count := 0

	for !allStatesAreEndStates(currentStates) {
		for i := 0; i < len(instructions); i++ {
			var temp []string
			for _, state := range currentStates {
				entry := points[state]
				if instructions[i] == 'L' {
					temp = append(temp, entry[0])
				} else if instructions[i] == 'R' {
					temp = append(temp, entry[1])
				}
			}
			currentStates = temp
		}
		count++
		println(count)
	}
}

func allStatesAreEndStates(states []string) bool {
	result := true

	for _, state := range states {
		if state[2] != 'Z' {
			result = false
		}
	}

	return result
}

func solveProb08Part1() {
	//var input = ReadFileIntoArray("resources/dec_08/example.txt")
	//var input = ReadFileIntoArray("resources/dec_08/example2.txt")
	var input = ReadFileIntoArray("resources/dec_08/input.txt")

	points := make(map[string][]string, len(input)-2)

	instructions := input[0]

	for i := 2; i < len(input); i++ {
		split := strings.Split(input[i], "=")

		head := split[0]

		lenBody := len(split[1]) - 1
		body := strings.Split(split[1][2:lenBody], ", ")

		points[strings.TrimSpace(head)] = body
	}

	goalState := "ZZZ"
	currentState := "AAA"
	count := 0

	for currentState != goalState {
		for i := 0; i < len(instructions); i++ {
			entry := points[currentState]
			if instructions[i] == 'L' {
				currentState = entry[0]
			} else if instructions[i] == 'R' {
				currentState = entry[1]
			}
			count++
		}
	}

	println(count)
}
