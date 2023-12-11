package main

import (
	"fmt"
)

func main() {
	input := ReadFileIntoArray("resources/dec_11/example.txt")
	//input := ReadFileIntoArray("resources/dec_11/input.txt")

	expandedInput := expandInput(input)

	for _, line := range expandedInput {
		fmt.Println(line)
	}
}

func expandInput(input []string) []string {
	var returnVal []string
	for i := 0; i < len(input); i++ {
		var expand = true

		for j := 0; j < len(input); j++ {
			char := input[j][i]
			if char == '#' {
				expand = false
				break
			}
		}

		if expand {
			for j := 0; j < len(input); j++ {
				runes := []rune(input[j])
				runes = append(runes[:i], append([]rune{'.'}, runes[i:]...)...)
				input[j] = string(runes)
			}
		}
	}

	for i := 0; i < len(input); i++ {
		var expand = true
		tmpInput := input[i]
		for j := 0; j < len(input); j++ {
			if input[i][j] == '#' {
				expand = false
				break
			}
		}

		if expand {
			returnVal = append(returnVal, tmpInput)
			returnVal = append(returnVal, tmpInput)
		} else {
			returnVal = append(returnVal, tmpInput)
		}
	}

	return returnVal
}
