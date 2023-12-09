package main

func solveProblem09Part2() {
	//var input = ReadFileIntoArray("resources/dec_09/example.txt")
	//var input = ReadFileIntoArray("resources/dec_09/example2.txt")
	var input = ReadFileIntoArray("resources/dec_09/input.txt")

	result := 0
	for _, line := range input {
		numbers := StringToNumberOfIntegers(line)
		var differences [][]int

		differences = append(differences, numbers)
		acc := 1

		for !allZeroes(differences[acc-1]) {
			differences = append(differences, make([]int, len(differences[acc-1])-1))
			for i := 1; i < len(differences[acc-1]); i++ {
				diff := differences[acc-1][i] - differences[acc-1][i-1]
				differences[acc][i-1] = diff
			}
			acc++
		}

		tmp := differences[0][0]

		for j := 1; j < len(differences)-1; j += 2 {
			x := differences[j][0] - differences[j+1][0]
			tmp -= x
		}
		result += tmp
	}

	println(result)
}

func solveProblem09Part1() {
	//var input = ReadFileIntoArray("resources/dec_09/example.txt")
	//var input = ReadFileIntoArray("resources/dec_09/example2.txt")
	var input = ReadFileIntoArray("resources/dec_09/input.txt")

	result := 0
	for _, line := range input {
		numbers := StringToNumberOfIntegers(line)
		var differences [][]int

		differences = append(differences, numbers)
		acc := 1

		for !allZeroes(differences[acc-1]) {
			differences = append(differences, make([]int, len(differences[acc-1])-1))
			for i := 1; i < len(differences[acc-1]); i++ {
				diff := differences[acc-1][i] - differences[acc-1][i-1]
				differences[acc][i-1] = diff
			}
			acc++
		}

		tmp := 0

		for _, arr := range differences {
			if len(arr) > 1 && arr != nil {
				tmp = tmp + arr[len(arr)-1]
			}
		}

		result += tmp
	}

	println(result)
}

func allZeroes(arr []int) bool {
	if arr == nil {
		return false
	}

	count := 0
	for _, elem := range arr {
		if elem == 0 {
			count++
		}
	}

	return len(arr) == count
}
