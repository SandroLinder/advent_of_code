package main

import (
	"advent_of_code/aoc_2024/util"
	"fmt"
)

func printInput(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}

type Antenna struct {
	location    util.Point
	antennaType string
}

func (a Antenna) String() string {
	return fmt.Sprintf("\ny: %d, x: %d, type:  %s\n", a.location.Y, a.location.X, a.antennaType)
}

func parseAntennas(input []string) []Antenna {
	var antennas []Antenna
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] != '.' {
				antennas = append(antennas, Antenna{location: util.Point{X: x, Y: y}, antennaType: string(input[y][x])})
			}
		}
	}

	return antennas
}

func slope(a, b util.Point) float32 {
	dy := float32(b.Y - a.Y)
	dx := float32(b.X - a.X)

	if dy == 0 || dx == 0 {
		return 0
	}
	return dx / dy
}

func samePoint(a, b util.Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func solveExercise1Day08(input []string) {
	antennas := parseAntennas(input)
	sum := 0
	validPoints := make(map[util.Point]bool)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			point := util.Point{Y: y, X: x}
			for a1Index := 0; a1Index < len(antennas); a1Index++ {
				a1 := antennas[a1Index]
				if samePoint(point, a1.location) {
					continue
				}

				for a2Index := a1Index + 1; a2Index < len(antennas); a2Index++ {
					a2 := antennas[a2Index]
					if a1.antennaType != a2.antennaType {
						continue
					}
					if samePoint(point, a2.location) {
						continue
					}

					if a1 == a2 {
						continue
					}

					slopeA1 := slope(a1.location, point)
					slopeA2 := slope(point, a2.location)
					antennaSlope := slope(a1.location, a2.location)

					if slopeA1 == antennaSlope && slopeA2 == antennaSlope {
						dA1 := ManhattanDistance(point, a1.location)
						dA2 := ManhattanDistance(point, a2.location)

						if dA1 == 2*dA2 || dA2 == 2*dA1 {
							_, ok := validPoints[point]
							if !ok {
								validPoints[point] = true
								sum += 1
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func solveExercise2Day08(input []string) {
	antennas := parseAntennas(input)
	sum := 0
	validPoints := make(map[util.Point]bool)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			point := util.Point{Y: y, X: x}
			for a1Index := 0; a1Index < len(antennas); a1Index++ {
				a1 := antennas[a1Index]
				if samePoint(point, a1.location) {
					continue
				}

				for a2Index := a1Index + 1; a2Index < len(antennas); a2Index++ {
					a2 := antennas[a2Index]
					if a1.antennaType != a2.antennaType {
						continue
					}
					if samePoint(point, a2.location) {
						continue
					}

					if a1 == a2 {
						continue
					}

					slopeA1 := slope(a1.location, point)
					slopeA2 := slope(point, a2.location)
					antennaSlope := slope(a1.location, a2.location)

					if slopeA1 == antennaSlope && slopeA2 == antennaSlope {
						_, ok := validPoints[point]
						if !ok {
							validPoints[point] = true
							sum += 1
						}
					}
				}
			}
		}
	}

	for key, _ := range validPoints {
		fmt.Printf("(%d, %d)\n", key.Y, key.X)
	}

	fmt.Println(sum)
}
