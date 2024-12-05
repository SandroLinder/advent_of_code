package main

import (
	"fmt"
	"slices"
)

type Rule struct {
	after *[]int
}

type Update struct {
	elements []int
}

func parseInput(input []string) (map[int]*Rule, []Update) {
	index := 0

	rules := make(map[int]*Rule)
	for input[index] != "" {
		split := StringToNumberOfIntegersBySymbol(input[index], "|")

		old, ok := rules[split[0]]

		if !ok {
			rule := &Rule{after: &[]int{split[1]}}
			rules[split[0]] = rule
		} else {
			*old.after = append(*old.after, split[1])
			slices.Sort(*old.after)
		}
		index++
	}

	index++

	var updates []Update

	for index < len(input) {
		split := StringToNumberOfIntegersBySymbol(input[index], ",")

		updates = append(updates, Update{
			elements: split,
		})
		index++
	}

	return rules, updates
}

func solveExercise1Day05(input []string) {
	rules, updates := parseInput(input)

	sum := 0
	for _, update := range updates {
		valid := true
		for i := len(update.elements) - 1; i >= 0; i-- {
			current := update.elements[i]
			for j := i - 1; j >= 0; j-- {
				checking := update.elements[j]
				if valid {
					exists, ok := rules[current]

					if !ok {
						continue
					} else {
						if slices.Contains(*exists.after, checking) {
							valid = false
						}
					}
				}
			}
		}

		if valid {
			middle := len(update.elements) / 2
			sum += update.elements[middle]
		}
	}

	fmt.Println(sum)
}

func solveExercise2Day05(input []string) {
	fmt.Print(input)
}
