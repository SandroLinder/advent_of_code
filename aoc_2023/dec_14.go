package main

import "fmt"

func main() {
	//input := ReadFileIntoArray("resources/dec_14/example.txt")
	input := ReadFileIntoArray("resources/dec_14/input.txt")

	var arr = make([][]uint8, len(input))

	for i := 0; i < len(input); i++ {
		arr[i] = make([]uint8, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			arr[i][j] = input[i][j]
		}
	}

	for i := 0; i < 1000; i++ {
		// North
		moveRocksUp(arr)
		//printArray(arr)
		// West
		moveRocksLeft(arr)
		//printArray(arr)
		// South
		moveRocksDown(arr)
		//printArray(arr)
		// East
		moveRocksRight(arr)
		//printArray(arr)
	}

	sum := 0
	for acc, row := range arr {
		for i := 0; i < len(row); i++ {
			if row[i] == 'O' {
				sum += len(arr) - acc
			}
		}
	}

	println(sum)
}

func printArray(arr [][]uint8) {
	for _, row := range arr {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}

	fmt.Println()
}

func moveRocksLeft(arr [][]uint8) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			if arr[i][j] == 'O' {
				newCol := -1

				for k := j - 1; k >= 0; k-- {
					if arr[i][k] != 'O' && arr[i][k] != '#' {
						newCol = k
					} else {
						break
					}
				}

				if newCol != -1 {
					arr[i][newCol] = 'O'
					arr[i][j] = '.'
				}
			}
		}
	}
}

func moveRocksRight(arr [][]uint8) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr[i]) - 2; j >= 0; j-- {
			if arr[i][j] == 'O' {
				newCol := -1

				for k := j + 1; k < len(arr[i]); k++ {
					if arr[i][k] != 'O' && arr[i][k] != '#' {
						newCol = k
					} else {
						break
					}
				}

				if newCol != -1 {
					arr[i][newCol] = 'O'
					arr[i][j] = '.'
				}
			}
		}
	}
}

func moveRocksUp(arr [][]uint8) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 'O' {
				newRow := -1

				for k := i - 1; k >= 0; k-- {
					if arr[k][j] != 'O' && arr[k][j] != '#' {
						newRow = k
					} else {
						break
					}
				}

				if newRow != -1 {
					arr[newRow][j] = 'O'
					arr[i][j] = '.'
				}
			}
		}
	}
}

func moveRocksDown(arr [][]uint8) {
	for i := len(arr) - 2; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 'O' {
				newRow := -1

				for k := i + 1; k < len(arr[i]); k++ {
					char := arr[k][j]
					if char != 'O' && char != '#' {
						newRow = k
					} else {
						break
					}
				}

				if newRow != -1 {
					arr[newRow][j] = 'O'
					arr[i][j] = '.'
				}
			}
		}
	}
}
