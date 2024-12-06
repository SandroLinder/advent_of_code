package main

import (
	"advent_of_code/aoc_2024/util"
	"fmt"
)

var up = Direction{x: 0, y: -1}
var down = Direction{x: 0, y: 1}
var left = Direction{x: -1, y: 0}
var right = Direction{x: 1, y: 0}

var guardDirections = map[uint8]Direction{
	'^': up,
	'v': down,
	'<': left,
	'>': right,
}

type Pos struct {
	x int
	y int
}

func findStartDirection(input []string) (Direction, Pos) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			direction, ok := guardDirections[input[y][x]]
			if ok {
				return direction, Pos{y: y, x: x}
			}
		}
	}

	panic("no guard found")
}

func walk(input []string, current Pos, direction Direction, path *util.Stack[Pos]) bool {
	nextPos := Pos{x: current.x + direction.x, y: current.y + direction.y}
	if nextPos.x < 0 || nextPos.y < 0 || nextPos.x >= len(input[0]) || nextPos.y >= len(input) {
		path.Push(current)
		return true
	}

	path.Push(current)

	if input[nextPos.y][nextPos.x] == '#' {
		var newDir Direction
		if direction == up {
			newDir = right
		} else if direction == down {
			newDir = left
		} else if direction == left {
			newDir = up
		} else if direction == right {
			newDir = down
		}

		if walk(input, current, newDir, path) {
			return true
		}
	} else {
		if walk(input, nextPos, direction, path) {
			return true
		}
	}

	_, _ = path.Pop()

	return false
}

func solveExercise1Day06(input []string) {
	direction, start := findStartDirection(input)
	path := util.NewStack[Pos]()

	walk(input, start, direction, path)

	pathArray := path.ToArray()

	for _, elem := range pathArray {
		row := input[elem.y]
		newRow := []rune(row)
		newRow[elem.x] = 'X'
		input[elem.y] = string(newRow)
	}

	for _, row := range input {
		fmt.Println(row)
	}

	fmt.Println(path.UniqueItems())
}

func solveExercise2Day06(input []string) {
	println(input)
}
