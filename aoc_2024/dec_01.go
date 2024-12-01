package main

import (
	"advent_of_code/aoc_2024/util"
	"container/heap"
	"strconv"
	"strings"
)

func parse(input string) (*util.MinHeap, *util.MinHeap) {
	left := &util.MinHeap{}
	*left = []int{}
	heap.Init(left)

	right := &util.MinHeap{}
	*right = []int{}
	heap.Init(right)

	rows := strings.Split(input, "\n")

	for _, r := range rows {
		row := strings.Fields(r)
		leftElement, _ := strconv.Atoi(row[0])
		rightElement, _ := strconv.Atoi(row[1])

		heap.Push(left, leftElement)
		heap.Push(right, rightElement)
	}

	return left, right
}

func solveExercise1Day01(input string) {
	left, right := parse(input)

	sum := 0
	for left.Len() > 0 {
		leftElem := heap.Pop(left).(int)
		rightElem := heap.Pop(right).(int)
		dist := Abs(leftElem - rightElem)
		sum += dist
	}

	println(sum)
}

func solveExercise2Day01(input string) {
	var left []int
	right := make(map[int]int)

	rows := strings.Split(input, "\n")

	for _, r := range rows {
		row := strings.Fields(r)
		leftElement, _ := strconv.Atoi(row[0])
		rightElement, _ := strconv.Atoi(row[1])

		left = append(left, leftElement)

		mr, ok := right[rightElement]

		if !ok {
			right[rightElement] = 1
		} else {
			right[rightElement] = mr + 1
		}
	}

	similarityScore := 0
	for _, elem := range left {
		rightCount, _ := right[elem]
		similarityScore += elem * rightCount
	}

	println(similarityScore)

}
