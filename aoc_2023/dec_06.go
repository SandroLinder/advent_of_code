package main

import (
	"strconv"
	"strings"
)

func solveProblemDec06() {
	//var input = ReadFileIntoArray("resources/dec_06/example.txt")
	var input = ReadFileIntoArray("resources/dec_06/input.txt")

	timeSplit := strings.Split(input[0], ":")
	recordSplit := strings.Split(input[1], ":")

	times := StringToNumberOfIntegers(timeSplit[1])
	records := StringToNumberOfIntegers(recordSplit[1])

	time := ""

	for _, t := range times {
		time += strconv.Itoa(t)
	}

	record := ""

	for _, r := range records {
		record += strconv.Itoa(r)
	}

	timeAsInt, _ := strconv.Atoi(time)
	recordAsInt, _ := strconv.Atoi(record)

	numberOfPossibleWins := 0
	for j := 0; j < timeAsInt; j++ {
		distanceTravelled := (timeAsInt - j) * j
		if distanceTravelled > recordAsInt {
			numberOfPossibleWins++
		}
	}

	println(numberOfPossibleWins)

}

func solveDec06Part01() {
	//var input = ReadFileIntoArray("resources/dec_06/example.txt")
	var input = ReadFileIntoArray("resources/dec_06/input.txt")

	timeSplit := strings.Split(input[0], ":")
	recordSplit := strings.Split(input[1], ":")

	times := StringToNumberOfIntegers(timeSplit[1])
	records := StringToNumberOfIntegers(recordSplit[1])

	numberOfPossibleHoldDurations := make([]int, len(times))
	for i := 0; i < len(times); i++ {
		record := records[i]
		duration := times[i]
		for j := 0; j < duration; j++ {
			distanceTravelled := (duration - j) * j
			if distanceTravelled > record {
				numberOfPossibleHoldDurations[i]++
			}
		}
	}

	ans := 1

	for _, number := range numberOfPossibleHoldDurations {
		ans *= number
	}

	println(ans)
}
