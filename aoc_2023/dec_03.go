package main

import (
	"fmt"
	"strconv"
)

type Location struct {
	x int
	y int
}

func (l Location) String() string {
	return fmt.Sprintf("Location: x: %d y: %d \n", l.x, l.y)
}

var directions = []Location{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}, {x: 1, y: -1}, {x: 1, y: 1}, {x: -1, y: -1}, {x: -1, y: 1}}

func main() {
	var input = ReadFileIntoArray("resources/dec_03/example01.txt")
	//var input = ReadFileIntoArray("resources/dec_03/input.txt")

	parsedInput := parseInput(input)

	sum := 0
	for i := 0; i < len(parsedInput); i++ {
		row := parsedInput[i]
		for j := 0; j < len(row); j++ {
			element := row[j]
			if isStar(element) {
				adjacentLocations := findAdjacentLocations(j, i, len(parsedInput), len(row))
				for _, location := range adjacentLocations {
					if IsDigit(parsedInput[location.y][location.x]) {
						left := location.x
						right := location.x

						number := ""
						for left > 0 && IsDigit(parsedInput[location.y][left]) {
							number = number + string(parsedInput[location.y][left])
							left--
						}
						println("Left: " + number)

						numberRight := ""
						for right < len(parsedInput[i]) && IsDigit(parsedInput[location.y][right]) {
							numberRight = numberRight + string(parsedInput[location.y][right])
							right++
						}
						println("Right: " + numberRight)
					}
				}
			}
		}
	}

	println(sum)
}

func findAdjacentLocations(x int, y int, height int, width int) []Location {
	var adjacentLocations []Location
	for _, direction := range directions {
		newY := y + direction.y
		if newY > height-1 || newY < 0 {
			continue
		}

		newX := x + direction.x
		if newX > width-1 || newX < 0 {
			continue
		}

		adjacentLocations = append(adjacentLocations, Location{x: x + direction.x, y: y + direction.y})
	}
	return adjacentLocations
}

func isStar(char uint8) bool {
	return char == '*'
}

func solvePart01() {
	var input = ReadFileIntoArray("resources/dec_03/example01.txt")
	//var input = ReadFileIntoArray("resources/dec_03/input.txt")

	parsedInput := parseInput(input)
	sumAdjacent := 0

	for i := 0; i < len(parsedInput); i++ {
		row := parsedInput[i]
		number := ""
		anyNumberAdjacentToSymbol := false
		for j := 0; j < len(row); j++ {
			char := row[j]
			if IsDigit(char) {
				isAdjacent := checkAdjacent(parsedInput, i, j, len(parsedInput), len(row))
				if isAdjacent {
					anyNumberAdjacentToSymbol = true
				}

				number = number + string(char)
			} else {
				if anyNumberAdjacentToSymbol {
					numberAsInt, _ := strconv.Atoi(number)
					sumAdjacent += numberAsInt
				}
				number = ""
				anyNumberAdjacentToSymbol = false
			}
		}

		if anyNumberAdjacentToSymbol {
			numberAsInt, _ := strconv.Atoi(number)
			sumAdjacent += numberAsInt
		}
	}
	fmt.Println(sumAdjacent)
}

func checkAdjacent(input [][]uint8, i int, j int, height int, width int) bool {
	for _, direction := range directions {
		y := i + direction.y
		if y > height-1 || y < 0 {
			continue
		}

		x := j + direction.x
		if x > width-1 || x < 0 {
			continue
		}

		symbol := input[i+direction.y][j+direction.x]
		if !IsDigit(symbol) && symbol != '.' {
			return true
		}
	}

	return false
}

func parseInput(input []string) [][]uint8 {
	result := make([][]uint8, len(input))

	for i := range result {
		result[i] = make([]uint8, len(input[0]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			char := input[i][j]
			result[i][j] = char
		}
	}

	return result
}
