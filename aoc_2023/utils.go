package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsDigit(u uint8) bool {
	return u >= 48 && u <= 57
}

func ReadFileIntoArray(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var inputArray []string

	for fileScanner.Scan() {
		inputArray = append(inputArray, fileScanner.Text())
	}

	_ = file.Close()

	return inputArray
}

func ReadFileIntoString(path string) string {
	file, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	return string(file)
}

func ReverseString(s string) string {
	reversedString := ""
	end := len(s)
	for end > 0 {
		reversedString = fmt.Sprint(reversedString, string(s[end-1]))
		end--
	}

	return reversedString
}

func StringToNumberOfIntegers(s string) []int {
	numbers := strings.Split(s, " ")
	var result []int

	for _, number := range numbers {
		if number == "" {
			continue
		}
		numberAsInt, _ := strconv.Atoi(number)
		result = append(result, numberAsInt)
	}

	return result
}

func StringToNumberOfIntegersBySymbol(s string, symbol string) []int {
	symbolToSplit := " "
	if symbol != "" {
		symbolToSplit = symbol
	}
	numbers := strings.Split(s, symbolToSplit)
	var result []int

	for _, number := range numbers {
		if number == "" {
			continue
		}
		numberAsInt, _ := strconv.Atoi(number)
		result = append(result, numberAsInt)
	}

	return result
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
