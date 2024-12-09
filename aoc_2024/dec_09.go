package main

import (
	"fmt"
)

func checkSum(input []uint16) int64 {
	sum := int64(0)
	for i := 0; i < len(input); i++ {
		if input[i] != 20000 {
			sum += int64(i) * int64(input[i])
		}
	}
	return sum
}

func swap(arr []uint16, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func solveExercise1Day09(input string) {
	inputAsNumbers := StringToNumberOfIntegersBySymbol(input, "")
	var extendedFiles []uint16
	id := 0
	for i := 0; i < len(inputAsNumbers); i++ {
		if IsEven(i) {
			for j := 0; j < inputAsNumbers[i]; j++ {
				extendedFiles = append(extendedFiles, uint16(id))
			}
			id++
		} else {
			for j := 0; j < inputAsNumbers[i]; j++ {
				extendedFiles = append(extendedFiles, uint16(20000))
			}
		}
	}

	start := 0
	end := len(extendedFiles) - 1

	fmt.Println(extendedFiles)

	for start < end {
		for extendedFiles[start] != 20000 && start < end {
			start++
		}

		for extendedFiles[end] == 20000 && end > start {
			end--
		}

		if extendedFiles[start] == 20000 && start < end {
			swap(extendedFiles, start, end)
		}
		start++
		end--
	}
	fmt.Println()
	fmt.Println(extendedFiles)
	fmt.Println(checkSum(extendedFiles))

	//6283404590840
	//6281533534769
}

func solveExercise2Day09(input string) {
	fmt.Println(input)
}
