package main

import (
	"fmt"
	"sort"
	"strings"
)

func NewHandWithJoker(cards string, bet int) Hand {
	hand := Hand{cards: cards, bet: bet}

	strength := 0

	if hand.isDistinct() {
		strength = 1
	} else if hand.isTwoPair() {
		strength = 3
	} else if hand.isFullHouse() {
		strength = 5
	} else if hand.isXOfAKind(2) {
		strength = 2
	} else if hand.isXOfAKind(3) {
		strength = 4
	} else if hand.isXOfAKind(4) {
		strength = 6
	} else if hand.isXOfAKind(5) {
		strength = 7
	}

	hand.strength = strength
	return hand
}

func getCardStrengthWithJoker(card uint8) int {
	switch card {
	case 'J':
		return 0
	case '2':
		return 1
	case '3':
		return 2
	case '4':
		return 3
	case '5':
		return 4
	case '6':
		return 5
	case '7':
		return 6
	case '8':
		return 7
	case '9':
		return 8
	case 'T':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	default:
		panic(fmt.Sprintf("card not recognized '%c'", card))
	}
}

func compareCardsWithJoker(a string, b string) bool {
	for i := 0; i < len(a); i++ {
		strengthA := getCardStrengthWithJoker(a[i])
		strengthB := getCardStrengthWithJoker(b[i])

		if strengthA < strengthB {
			return false
		} else if strengthB < strengthA {
			return true
		}
	}

	return false
}

func main() {
	input := ReadFileIntoArray("resources/dec_07/example.txt")
	//input := ReadFileIntoArray("resources/dec_07/example2.txt")
	//input := ReadFileIntoArray("resources/dec_07/input.txt")

	var hands []Hand
	for _, line := range input {
		split := strings.Split(line, " ")

		hands = append(hands, NewHandWithJoker(split[0], StringToNumberOfIntegers(split[1])[0]))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].strength > hands[j].strength
	})

	var rankedHands []Hand

	var strengthMap = make(map[int][]Hand, len(hands))
	for i := 0; i < len(hands); i++ {
		hand := hands[i]

		strengthMap[hand.strength] = append(strengthMap[hand.strength], hand)
	}

	var keys []int

	for k, _ := range strengthMap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, k := range keys {
		entry := strengthMap[k]
		sort.SliceStable(entry, func(i, j int) bool {
			return compareCardsWithJoker(entry[i].cards, entry[j].cards)
		})

		rankedHands = append(rankedHands, entry...)
	}

	sum := 0
	for acc, hand := range rankedHands {
		rank := len(rankedHands) - acc
		sum += hand.bet * rank
	}
	println(sum)
}
