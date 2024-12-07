package main

import (
	"advent_of_code/aoc_2024/util"
	"fmt"
	"slices"
)

type Rule struct {
	after *[]int
}

type Rule2 struct {
	left  int
	right int
}

type Update struct {
	elements []int
}

func parseInRules(input []string) (map[int]*Rule, []Update) {
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
	in, updates := parseInRules(input)

	sum := 0
	for _, update := range updates {
		valid := checkUpdate(in, update)

		if valid {
			middle := len(update.elements) / 2
			sum += update.elements[middle]
		}
	}

	fmt.Println(sum)
}

func parseOutRules(input []string) (map[int]*Rule, []Update) {
	index := 0

	rules := make(map[int]*Rule)
	for input[index] != "" {
		split := StringToNumberOfIntegersBySymbol(input[index], "|")

		old, ok := rules[split[1]]

		if !ok {
			rule := &Rule{after: &[]int{split[0]}}
			rules[split[1]] = rule
		} else {
			*old.after = append(*old.after, split[0])
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

func solveExercise2Day05(input []string) {
	inRules, updates := parseInRules(input)
	outRules, updates := parseOutRules(input)
	sum := 0
	for _, update := range updates {
		if checkUpdate(inRules, update) {
			continue
		}

		middle := sortGraph(inRules, outRules, update)
		sum += middle
	}

	fmt.Println(sum)
}

func walkGraph(rules map[int][]int, curr int, nodes []int, path *util.Stack[int]) bool {
	path.Push(curr)

	if path.Length == len(nodes) {
		return true
	}

	list := rules[curr]

	for _, elem := range list {
		if !slices.Contains(nodes, elem) {
			continue
		}

		if walkGraph(rules, elem, nodes, path) {
			return true
		}
	}

	_, _ = path.Pop()

	return false
}

func sortGraph(inRules map[int]*Rule, outRules map[int]*Rule, update Update) int {
	var all []int
	for k, v := range inRules {
		if !slices.Contains(all, k) {
			all = append(all, k)
		}

		for _, val := range *v.after {
			if !slices.Contains(all, val) {
				all = append(all, val)
			}
		}
	}

	in := buildAdjacencyList(inRules, update.elements)
	out := buildAdjacencyList(outRules, update.elements)
	sum := 0

	start := 0

	for key, value := range out {
		if value == nil || len(value) == 0 {
			start = key
			break
		}
	}

	path := util.NewStack[int]()
	walkGraph(in, start, update.elements, path)
	sum += path.ToArray()[path.Length/2]

	return sum
}

func buildAdjacencyList(rules map[int]*Rule, nodes []int) map[int][]int {
	list := make(map[int][]int)

	for _, elem := range nodes {
		var adj []int
		rule := rules[elem]

		if rule != nil {
			for _, v := range *rule.after {
				if slices.Contains(nodes, v) {
					adj = append(adj, v)
				}
			}
			list[elem] = adj
		} else {
			list[elem] = make([]int, 0)
		}

	}

	return list
}

func buildAdjacencyList2(rules map[int]*Rule, nodes []int) map[int][]int {
	list := make(map[int][]int)

	for _, elem := range nodes {
		var adj []int
		rule := rules[elem]

		if rule != nil {
			adj = append(adj, *rule.after...)
			list[elem] = adj
		} else {
			list[elem] = make([]int, 0)
		}

	}

	return list
}

func checkUpdate(rules map[int]*Rule, update Update) bool {
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
	return valid
}
