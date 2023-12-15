package main

import (
	"strconv"
	"strings"
)

const (
	ASSIGN = iota
	REMOVE
)

type Lense struct {
	focal int
	label string
}

func (l Lense) getValueToCompare() interface{} {
	return l.label
}

type Box struct {
	lenses []Lense
}

func contains(s []Lense, e Lense) bool {
	for _, a := range s {
		if a.label == e.label {
			return true
		}
	}
	return false
}

func remove(slice []Lense, s int) []Lense {
	return append(slice[:s], slice[s+1:]...)
}

func solveProblem15Part2() {
	//input := ReadFileIntoArray("resources/dec_15/example.txt")
	input := ReadFileIntoArray("resources/dec_15/input.txt")

	arr := strings.Split(input[0], ",")

	HASHMAP := make([]Box, 256)
	for _, elem := range arr {
		hashKey := 0
		label := ""
		strength := -1
		operation := -1
		if strings.HasSuffix(elem, "-") {
			label = elem[:len(elem)-1]
			hashKey = hash(label)
			operation = REMOVE
		} else {
			label = elem[:len(elem)-2]
			hashKey = hash(label)
			strength, _ = strconv.Atoi(string(elem[len(elem)-1]))
			operation = ASSIGN
		}

		lense := Lense{label: label, focal: strength}

		if operation == ASSIGN {
			if HASHMAP[hashKey].lenses == nil {
				HASHMAP[hashKey].lenses = append(HASHMAP[hashKey].lenses, Lense{label: label, focal: strength})
			} else {
				if contains(HASHMAP[hashKey].lenses, lense) {
					for i, tmp := range HASHMAP[hashKey].lenses {
						if tmp.label == lense.label {
							HASHMAP[hashKey].lenses[i] = lense
						}
					}
				} else {
					HASHMAP[hashKey].lenses = append(HASHMAP[hashKey].lenses, lense)
				}
			}
		} else {
			if HASHMAP[hashKey].lenses != nil {
				if contains(HASHMAP[hashKey].lenses, lense) {
					for i, tmp := range HASHMAP[hashKey].lenses {
						if tmp.label == lense.label {
							HASHMAP[hashKey].lenses = remove(HASHMAP[hashKey].lenses, i)
						}
					}
				}
			}
		}
	}

	ans := 0
	for acc, box := range HASHMAP {
		for accL, l := range box.lenses {
			ans += (acc + 1) * (accL + 1) * l.focal
		}
	}

	println(ans)
}

func solveProblem15Part1() {
	input := ReadFileIntoArray("resources/dec_15/example.txt")
	//input := ReadFileIntoArray("resources/dec_15/input.txt")

	arr := strings.Split(input[0], ",")

	ans := 0
	for _, elem := range arr {
		ans += hash(elem)
	}

	println(ans)
}

func hash(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result += int(s[i])
		result *= 17
		result %= 256
	}

	return result
}
