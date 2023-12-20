package main

import (
	"aoc_2023/util"
	"fmt"
	"github.com/mitchellh/hashstructure/v2"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func directionToString(dir int) string {
	switch dir {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		panic("unhandled default case")
	}
}

type Field struct {
	y int
	x int
}

type Next struct {
	field     Field
	direction int
}

func (n Next) String() string {
	return fmt.Sprintf("Field: %v, Direction: %s", n.field, directionToString(n.direction))
}

func toArray(input []string) [][]uint8 {
	var arr = make([][]uint8, len(input))

	for j, line := range input {
		arr[j] = make([]uint8, len(line))
		for i := 0; i < len(line); i++ {
			arr[j][i] = line[i]
		}
	}

	return arr
}

func solveProblem16Part1() {
	input := toArray(ReadFileIntoArray("resources/dec_16/example.txt"))
	//input := toArray(ReadFileIntoArray("resources/dec_16/input.txt"))

	var nexts = util.NewQueue(20)

	var alreadyVisited = make(map[uint64]bool)

	_ = nexts.Enqueue(Next{field: Field{x: 0, y: 0}, direction: RIGHT})

	for nexts.CountItemsInQueue() > 0 {
		nexts.Print()

		state := nexts.Dequeue().(Next)

		nextField := state.field

		if nextField.x < 0 || nextField.y < 0 || nextField.x >= len(input[0]) || nextField.y >= len(input) {
			continue
		}

		symbol := input[nextField.y][nextField.x]

		var actualNexts []Next

		if symbol == '|' && state.direction == RIGHT {
			if state.field.y > 0 {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x, y: nextField.y - 1}, direction: UP})
			}
			if state.field.y < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x, y: nextField.y + 1}, direction: DOWN})
			}
		}

		if symbol == '|' && state.direction == DOWN {
			if state.field.y < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x, y: nextField.y + 1}, direction: DOWN})
			}
		}

		if symbol == '-' && state.direction == DOWN {
			if state.field.x > 0 {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x - 1, y: nextField.y}, direction: LEFT})
			}
			if state.field.x < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x + 1, y: nextField.y}, direction: RIGHT})
			}
		}

		if symbol == '-' && state.direction == UP {
			if state.field.x > 0 {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x - 1, y: nextField.y}, direction: LEFT})
			}
			if state.field.x < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x + 1, y: nextField.y}, direction: RIGHT})
			}
		}

		if symbol == '-' && state.direction == RIGHT {
			if state.field.x < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x + 1, y: nextField.y}, direction: RIGHT})
			}
		}

		if symbol == '\\' && state.direction == RIGHT {
			if state.field.y < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x, y: nextField.y + 1}, direction: DOWN})
			}
		}

		if symbol == '/' && state.direction == RIGHT {
			if state.field.y > 0 {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x, y: nextField.y - 1}, direction: UP})
			}
		}

		if symbol == '/' && state.direction == UP {
			if state.field.x < len(input) {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x + 1, y: nextField.y}, direction: RIGHT})
			}
		}

		if symbol == '-' && state.direction == LEFT {
			if state.field.x > 0 {
				actualNexts = append(actualNexts, Next{field: Field{x: nextField.x - 1, y: nextField.y}, direction: LEFT})
			}
		}

		if symbol == '.' {
			next := getNextField(state)
			if next.x < 0 || next.y < 0 || next.x >= len(input[0]) || next.y >= len(input) {
				continue
			}

			actualNexts = append(actualNexts, Next{field: next, direction: state.direction})
		}

		for _, n := range actualNexts {
			hashNew, _ := hashstructure.Hash(n, hashstructure.FormatV2, nil)

			_, ok := alreadyVisited[hashNew]

			if ok {
				continue
			} else {
				_ = nexts.Enqueue(n)
				alreadyVisited[hashNew] = true
			}
		}
	}

}

func getNextField(state Next) Field {
	switch state.direction {
	case UP:
		return Field{x: state.field.x, y: state.field.y - 1}
	case DOWN:
		return Field{x: state.field.x, y: state.field.y + 1}
	case LEFT:
		return Field{x: state.field.x - 1, y: state.field.y}
	case RIGHT:
		return Field{x: state.field.x + 1, y: state.field.y}
	default:
		panic("not reachable")
	}
}
