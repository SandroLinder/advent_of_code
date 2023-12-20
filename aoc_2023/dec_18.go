package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	U = "U"
	D = "D"
	L = "L"
	R = "R"
)

type Move struct {
	direction     string
	numberOfTiles int
	colour        string
}

type LavaDuct struct {
	area   [][]uint8
	maxX   int
	maxY   int
	minX   int
	minY   int
	startX int
	startY int
}

func solveProblem18Part1() {
	input := ReadFileIntoArray("resources/dec_18/example.txt")
	//input := ReadFileIntoArray("resources/dec_18/input.txt")

	var moves []Move
	for _, line := range input {
		split := strings.Split(line, " ")
		tiles, _ := strconv.Atoi(split[1])

		moves = append(moves, Move{direction: split[0], numberOfTiles: tiles, colour: split[2]})
	}

	tuct := buildLavaDuct(moves)

	moveY := tuct.startY
	moveX := tuct.startX

	up := []int{-1, 0}
	down := []int{1, 0}
	left := []int{0, -1}
	right := []int{0, 1}

	for _, move := range moves {
		newY := -1
		newX := -1
		switch move.direction {
		case U:
			newY = moveY - move.numberOfTiles
			for i := moveY; i >= newY; i-- {
				tuct.area[i][moveX] = '#'
			}
		case D:
			newY = moveY + move.numberOfTiles
			for i := moveY; i <= newY; i++ {
				tuct.area[i][moveX] = '#'
			}
		case L:
			newX = moveX - move.numberOfTiles
			for i := moveX; i >= newX; i-- {
				tuct.area[moveY][i] = '#'
			}
		case R:
			newX = moveX + move.numberOfTiles
			for i := moveX; i <= newX; i++ {
				tuct.area[moveY][i] = '#'
			}
		default:
			panic("Not reachable")
		}

		if newY != -1 {
			moveY = newY
		}

		if newX != -1 {
			moveX = newX
		}
	}

	for i := 0; i < len(tuct.area); i++ {
		for j := 0; j < len(tuct.area[i]); j++ {
			symbol := tuct.area[i][j]

			if symbol == '#' {
				continue
			}

			curX := j
			curY := i

			isEnclosedTop := false
			isEnclosedBottom := false
			isEnclosedLeft := false
			isEnclosedRight := false

			for curY >= 0 && curX >= 0 {
				if tuct.area[curY][curX] == '#' {
					isEnclosedTop = true
				}

				curY = curY + up[0]
				curX = curX + up[1]
			}

			curX = j
			curY = i

			for curY >= 0 && curX >= 0 {
				if tuct.area[curY][curX] == '#' {
					isEnclosedBottom = true
				}

				curY = curY + down[0]
				curX = curX + down[1]
			}

			curX = j
			curY = i

			for curY >= 0 && curX >= 0 {
				if tuct.area[curY][curX] == '#' {
					isEnclosedLeft = true
				}

				curY = curY + left[0]
				curX = curX + left[1]
			}

			curX = j
			curY = i

			for curY >= 0 && curX >= 0 {
				if tuct.area[curY][curX] == '#' {
					isEnclosedRight = true
				}

				curY = curY + right[0]
				curX = curX + right[1]
			}

			if isEnclosedTop && isEnclosedBottom && isEnclosedLeft && isEnclosedRight {
				tuct.area[i][j] = '#'
			}
		}
	}

	for _, line := range tuct.area {
		fmt.Println(line)
	}
}

func buildLavaDuct(moves []Move) LavaDuct {
	x := 0
	y := 0

	maxX := 0
	maxY := 0

	minX := 9999999999
	minY := 9999999999

	for _, move := range moves {
		switch move.direction {
		case U:
			y = y - move.numberOfTiles
		case D:
			y = y + move.numberOfTiles
		case L:
			x = x - move.numberOfTiles
		case R:
			x = x + move.numberOfTiles
		default:
			panic("Not reachable")
		}
		maxX = maxInt(maxX, x)
		maxY = maxInt(maxY, y)
		minX = minInt(minX, x)
		minY = minInt(minY, y)
	}

	startY := -1 * minY
	startX := -1 * minX

	arr := make([][]uint8, maxY+startY+1)

	for i := 0; i < len(arr); i++ {
		arr[i] = make([]uint8, maxX+startX+1)
		for j := 0; j < maxX+startX+1; j++ {
			arr[i][j] = '.'
		}
	}

	return LavaDuct{
		area:   arr,
		maxX:   maxX,
		maxY:   maxY,
		minY:   minY,
		minX:   minX,
		startX: startX,
		startY: startY,
	}
}
