package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	GREATER = iota
	SMALLER
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Process struct {
	conditions []Condition
	name       string
}

type Condition struct {
	property    string
	operation   int
	value       int
	destination string
}

func main() {
	input := ReadFileIntoString("resources/dec_19/example.txt")
	//input := ReadFileIntoString("resources/dec_19/input.txt")

	split := strings.Split(input, "\r\n\r\n")

	//processes := split[0]

	var parts []Part

	partsString := strings.Split(split[1], "\r\n")
	for i := 0; i < len(partsString); i++ {
		parts = append(parts, parseParts(partsString[i]))
	}

	processes := make(map[string]Process)

	processesString := strings.Split(split[0], "\r\n")

	var processNames []string

	for i := 0; i < len(processesString); i++ {
		p := parseProcess(processesString[i])
		processes[p.name] = p
		processNames = append(processNames, p.name)
	}

	for _, part := range parts {
		startProcess := ""
		for _, processName := range processNames {
			process := processes[processName]
			for _, cond := range process.conditions {
				if cond.operation == GREATER {
					if getProperty(cond.property, part) > cond.value {
						startProcess = cond.destination
						break
					}
				} else if cond.operation == SMALLER {
					if getProperty(cond.property, part) < cond.value {
						startProcess = cond.destination
						break
					}
				} else {
					startProcess = cond.destination
					break
				}

			}
		}

		fmt.Println(startProcess)
	}
}

func getProperty(property string, part Part) int {
	switch property {
	case "x":
		return part.x
	case "m":
		return part.m
	case "a":
		return part.a
	case "s":
		return part.s
	default:
		panic("not reachable")
	}
}

func parseProcess(process string) Process {
	result := Process{}
	current := 0
	name := ""
	for process[current] != '{' {
		name = name + string(process[current])
		current++
	}
	result.name = name
	current++
	temp := process[current : len(process)-1]

	conditions := strings.Split(temp, ",")

	result.conditions = parseConditions(conditions)

	return result
}

func parseConditions(conditions []string) []Condition {
	var result []Condition

	for _, line := range conditions {
		condition := Condition{}
		current := 0
		if strings.Contains(line, "<") {
			condition.operation = SMALLER
			condition.property = string(line[current])
			current += 2
			value := ""
			for line[current] != ':' {
				value = value + string(line[current])
				current++
			}

			condition.value, _ = strconv.Atoi(value)

			current++
			destination := ""
			for current < len(line) {
				destination += string(line[current])
				current++
			}

			condition.destination = destination
		} else if strings.Contains(line, ">") {
			condition.operation = GREATER
			condition.property = string(line[current])
			current += 2
			value := ""
			for line[current] != ':' {
				value = value + string(line[current])
				current++
			}

			condition.value, _ = strconv.Atoi(value)

			current++
			destination := ""
			for current < len(line) {
				destination += string(line[current])
				current++
			}

			condition.destination = destination
		} else {
			destination := ""
			for current < len(line) {
				destination += string(line[current])
				current++
			}

			condition.operation = -1
			condition.destination = destination
		}

		result = append(result, condition)
	}

	return result
}

func parseParts(part string) Part {
	result := Part{}

	for i := 0; i < len(part); i++ {
		if part[i] == '{' || part[i] == '}' {
			continue
		}

		if part[i] == 'x' {
			i += 2
			x := ""
			for IsDigit(part[i]) {
				x = x + string(part[i])
				i++
			}

			result.x, _ = strconv.Atoi(x)
		}

		if part[i] == 'm' {
			i += 2
			m := ""
			for IsDigit(part[i]) {
				m = m + string(part[i])
				i++
			}

			result.m, _ = strconv.Atoi(m)
		}

		if part[i] == 'a' {
			i += 2
			a := ""
			for IsDigit(part[i]) {
				a = a + string(part[i])
				i++
			}

			result.a, _ = strconv.Atoi(a)
		}

		if part[i] == 's' {
			i += 2
			s := ""
			for IsDigit(part[i]) {
				s = s + string(part[i])
				i++
			}

			result.s, _ = strconv.Atoi(s)
		}
	}

	return result
}
