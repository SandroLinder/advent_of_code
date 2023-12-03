package main

import (
	"bufio"
	"fmt"
	"os"
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

func ReverseString(s string) string {
	reversedString := ""
	end := len(s)
	for end > 0 {
		reversedString = fmt.Sprint(reversedString, string(s[end-1]))
		end--
	}

	return reversedString
}
