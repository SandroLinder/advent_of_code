package main

import "fmt"

type Pattern struct {
	pattern []string
	height  int
	width   int
}

func solveProblem13Part1() {
	input := ReadFileIntoArray("resources/dec_13/example.txt")
	//input := ReadFileIntoArray("resources/dec_13/input.txt")

	var patterns []Pattern
	pattern := Pattern{}
	for i := 0; i < len(input)-1; i++ {
		if input[i] == "" {
			patterns = append(patterns, pattern)
			pattern = Pattern{}
			continue
		}

		pattern.pattern = append(pattern.pattern, input[i])
	}

	patterns = append(patterns, pattern)

	var matchingRows []int
	var matchingCols []int

	for _, pat := range patterns {
		i := 0
		if len(pat.pattern)%2 == 0 {
			i = 1
		}

		for i < len(pat.pattern)/2 {
			rowsAreMatching := true
			for j := 0; j < len(pat.pattern[i]); j++ {
				rowStart := pat.pattern[i]
				rowEnd := pat.pattern[len(pat.pattern)-1-i]
				if rowStart[j] != rowEnd[j] {
					rowsAreMatching = false
					break
				}
			}

			if rowsAreMatching {
				matchingRows = append(matchingRows, i+1)
			}
			i++
		}

		/*for j := 0; j < len(pat.pattern[0])-1; j++ {
			colMatching := true
			for i := 0; i < len(pat.pattern); i++ {
				if pat.pattern[i][j] != pat.pattern[i][j+1] {
					colMatching = false
					break
				}
			}

			if colMatching {
				matchingCols = append(matchingCols, j+1)
			}
		}*/
	}

	fmt.Println(matchingRows, matchingCols)
}
