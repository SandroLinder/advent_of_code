package main

func solveDec10Part1() {
	//var input = ReadFileIntoArray("resources/dec_10/example.txt")
	//var input = ReadFileIntoArray("resources/dec_10/example2.txt")
	//var input = ReadFileIntoArray("resources/dec_10/example3.txt")
	var input = ReadFileIntoArray("resources/dec_10/input.txt")

	var inputAsArray = make([][]int32, len(input))

	for i, line := range input {
		var chars []int32
		for _, char := range line {
			chars = append(chars, char)
		}
		inputAsArray[i] = chars
	}

	pipes := make(map[string][]string)

	pipes["F"] = []string{"D", "R"}
	pipes["L"] = []string{"U", "R"}
	pipes["J"] = []string{"U", "L"}
	pipes["7"] = []string{"D", "L"}
	pipes["-"] = []string{"R", "L"}
	pipes["|"] = []string{"U", "D"}

	moves := make(map[string][]int)

	moves["U"] = []int{-1, 0}
	moves["D"] = []int{1, 0}
	moves["L"] = []int{0, -1}
	moves["R"] = []int{0, 1}

	inverseMove := make(map[string]string)

	inverseMove["U"] = "D"
	inverseMove["D"] = "U"
	inverseMove["L"] = "R"
	inverseMove["R"] = "L"

	x, y := 0, 0
	newS := ""
	for i := 0; i < len(inputAsArray); i++ {
		for j := 0; j < len(inputAsArray[i]); j++ {
			if inputAsArray[i][j] == 'S' {
				y = i
				x = j

				upPossible := false
				leftPossible := false
				downPossible := isDownPossible(inputAsArray[y+1][x])
				rightPossible := isRightPossible(inputAsArray[y][x+1])

				if y > 0 {
					upPossible = isUpPossible(inputAsArray[y-1][x])
				}

				if x > 0 {
					leftPossible = isLeftPossible(inputAsArray[y][x-1])
				}

				if upPossible && downPossible {
					newS = "|"
				} else if upPossible && leftPossible {
					newS = "J"
				} else if upPossible && rightPossible {
					newS = "L"
				} else if leftPossible && rightPossible {
					newS = "-"
				} else if downPossible && rightPossible {
					newS = "F"
				} else if downPossible && leftPossible {
					newS = "7"
				}
			}
		}
	}

	move := pipes[newS][0]
	execMove := moves[pipes[newS][0]]

	currentY := y + execMove[0]
	currentX := x + execMove[1]

	distance := 0

	for currentY != y || currentX != x {
		symbol := string(inputAsArray[currentY][currentX])

		for _, m := range pipes[symbol] {
			inv := inverseMove[m]
			if inv == move {
				continue
			}

			move = m
			execMove = moves[m]
			break
		}

		currentY = currentY + execMove[0]
		currentX = currentX + execMove[1]
		distance++
	}

	println((distance + 1) / 2)
}

func isUpPossible(char int32) bool {
	if char == '|' || char == 'F' || char == '7' {
		return true
	}

	return false
}

func isDownPossible(char int32) bool {
	if char == '|' || char == 'J' || char == 'L' {
		return true
	}

	return false
}

func isRightPossible(char int32) bool {
	if char == '-' || char == 'J' || char == '7' {
		return true
	}

	return false
}

func isLeftPossible(char int32) bool {
	if char == '-' || char == 'F' || char == 'L' {
		return true
	}

	return false
}
