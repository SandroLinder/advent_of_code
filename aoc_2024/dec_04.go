package main

import (
	"fmt"
)

type Direction struct {
	x int
	y int
}

func (l Direction) String() string {
	return fmt.Sprintf("Direction: x: %d y: %d \n", l.x, l.y)
}

var allDirections = []Direction{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}, {x: 1, y: -1}, {x: 1, y: 1}, {x: -1, y: -1}, {x: -1, y: 1}}

func isWord(x, y int, input []string, dir Direction, word string) bool {
	nextX := x + dir.x
	nextY := y + dir.y
	for i := 0; i < len(word); i++ {
		if nextX >= 0 && nextY >= 0 && nextX < len(input[0]) && nextY < len(input) {
			next := input[nextY][nextX]
			expected := word[i]

			if next != expected {
				return false
			}
		} else {
			return false
		}

		nextX = nextX + dir.x
		nextY = nextY + dir.y
	}

	fmt.Println(y, x)
	return true
}

func solveExercise1Day04(input []string) {
	result := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			for _, dir := range allDirections {
				if input[y][x] == 'X' {
					if isWord(x, y, input, dir, "MAS") {
						result++
					}
				}
			}
		}
	}
	fmt.Println(result)
}

func solveExercise2Day04(input []string) {
	result := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == 'A' {
				if isMasToTheRight(x, y, input) && isMasToTheLeft(x, y, input) {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}

func isMasToTheRight(x int, y int, input []string) bool {
	yMinusOne := y - 1
	yPlusOne := y + 1

	xMinusOne := x - 1
	xPlusOne := x + 1

	if yMinusOne < 0 || yPlusOne >= len(input) || xMinusOne < 0 || xPlusOne >= len(input[0]) {
		return false
	}

	return (input[y-1][x-1] == 'M' && input[y+1][x+1] == 'S') || input[y-1][x-1] == 'S' && input[y+1][x+1] == 'M'
}

func isMasToTheLeft(x int, y int, input []string) bool {
	yMinusOne := y - 1
	yPlusOne := y + 1

	xMinusOne := x - 1
	xPlusOne := x + 1

	if yMinusOne < 0 || yPlusOne >= len(input) || xMinusOne < 0 || xPlusOne >= len(input[0]) {
		return false
	}

	return (input[y+1][x-1] == 'M' && input[y-1][x+1] == 'S') || input[y+1][x-1] == 'S' && input[y-1][x+1] == 'M'
}
