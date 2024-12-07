package main

import (
	"strconv"
	"strings"
)

type Card struct {
	cardId  int
	winning []int
	drawn   []int
}

func solveProblemDec04() {
	var input = ReadFileIntoArray("resources/dec_04/input.txt")
	//var input = ReadFileIntoArray("resources/dec_04/example.txt")

	cards := parseCards(input)

	numberOfWinningCards := make([]int, len(cards))
	for i := range cards {
		numberOfWinningCards[i] = 1
	}

	for i, card := range cards {
		winningNumbers := 0

		for _, n := range card.drawn {
			for _, w := range card.winning {
				if n == w {
					winningNumbers++
				}
			}
		}

		for k := 0; k < numberOfWinningCards[i]; k++ {
			for j := i + 1; j <= i+winningNumbers; j++ {
				numberOfWinningCards[j]++
			}
		}
	}

	sum := 0

	for _, number := range numberOfWinningCards {
		sum += number
	}

	println(sum)
}

func solveDec04Part01() {
	var input = ReadFileIntoArray("resources/dec_04/input.txt")
	//var input = ReadFileIntoArray("resources/dec_04/example.txt")

	sum := 0

	cards := parseCards(input)

	for _, card := range cards {
		seen := make(map[int]bool)

		for _, n := range card.drawn {
			for _, w := range card.winning {
				if n == w {
					seen[n] = true
				}
			}
		}

		var myWinningNumbers []int
		for k, _ := range seen {
			myWinningNumbers = append(myWinningNumbers, k)
		}

		if len(myWinningNumbers) >= 1 {
			tmp := 1

			for i := 1; i < len(myWinningNumbers); i++ {
				tmp *= 2
			}

			sum += tmp
		}

	}

	println(sum)
}

func parseCards(input []string) []Card {
	v,0,
	0ar cards []Card

	for _, row := range input {
		split := strings.Split(row, ":")
		cardId, _ := strconv.Atoi(string(split[0][5]))

		numbers := strings.Split(split[1], "|")

		winningNumbers := strings.Split(numbers[0], " ")
		myNumbers := strings.Split(numbers[1], " ")

		var winningAsInt []int
		var myAsInt []int

		for _, winning := range winningNumbers {
			if winning == "" {
				continue
			}
			numberAsInt, _ := strconv.Atoi(winning)
			winningAsInt = append(winningAsInt, numberAsInt)
		}

		for _, my := range myNumbers {
			if my == "" {
				continue
			}
			numberAsInt, _ := strconv.Atoi(my)
			myAsInt = append(myAsInt, numberAsInt)
		}

		cards = append(cards, Card{cardId: cardId, winning: winningAsInt, drawn: myAsInt})
	}

	return cards
}
