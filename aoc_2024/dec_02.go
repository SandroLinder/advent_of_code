package main

import "fmt"

func solveExercise1Day02(input []string) {
	sum := 0
	for _, i := range input {
		intInput := StringToNumberOfIntegers(i)

		diffs := make([]int, len(intInput)-1)
		for j := 0; j < len(intInput)-1; j++ {
			diffs[j] = intInput[j] - intInput[j+1]
		}

		safe := true
		dir := 0

		for _, diff := range diffs {
			if Abs(diff) > 3 || diff == 0 {
				safe = false
				break
			}

			if dir != 0 {
				if (diff < 0 && dir != -1) || (diff > 0 && dir != 1) {
					safe = false
					break
				}
			}

			if diff < 0 {
				dir = -1
			} else {
				dir = 1
			}
		}
		if safe {
			sum += 1
		}
	}

	println(sum)
}

func ArrayWithoutX(arr []int, x int) []int {
	if x < 0 || x >= len(arr) {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		return newArr
	}

	newArr := make([]int, 0, len(arr)-1)

	newArr = append(newArr, arr[:x]...)
	newArr = append(newArr, arr[x+1:]...)

	return newArr
}

func checkSafe(arr []int) bool {
	diffs := make([]int, len(arr)-1)
	for j := 0; j < len(arr)-1; j++ {
		diffs[j] = arr[j] - arr[j+1]
	}

	safe := true
	dir := 0

	for _, diff := range diffs {
		if Abs(diff) > 3 || diff == 0 {
			safe = false
			break
		}

		if dir != 0 {
			if (diff < 0 && dir != -1) || (diff > 0 && dir != 1) {
				safe = false
				break
			}
		}

		if diff < 0 {
			dir = -1
		} else {
			dir = 1
		}
	}

	return safe
}

func solveExercise2Day02(input []string) {
	sum := 0
	for _, i := range input {
		intInput := StringToNumberOfIntegers(i)
		var newInputs [][]int

		safe := checkSafe(intInput)

		if safe {
			sum += 1
			continue
		} else {
			for k := 0; k < len(intInput); k++ {
				newInput := ArrayWithoutX(intInput, k)
				newInputs = append(newInputs, newInput)
			}

			for _, temp := range newInputs {
				safe = checkSafe(temp)
				if safe {
					sum += 1
					break
				}
			}
		}
		fmt.Println(intInput)
		fmt.Println(safe)
	}

	println(sum)
}
