package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Round struct {
	red   int
	blue  int
	green int
}

func (r Round) String() string {
	return fmt.Sprintf("Round: blue: %d red: %d, green: %d\n", r.blue, r.red, r.green)
}

type Game struct {
	gameId int
	rounds []Round
}

func (g *Game) addRound(r Round) {
	g.rounds = append(g.rounds, r)
}

func (g Game) String() string {
	return fmt.Sprintf("Game %s: \n Rounds: %v\n", g.gameId, g.rounds)
}

func parseGames(input []string) []Game {
	var games []Game

	for _, row := range input {
		split := strings.Split(row, ":")
		gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])

		game := Game{gameId: gameId}

		rounds := strings.Split(split[1], ";")

		for _, round := range rounds {
			colors := strings.Split(round, ",")
			tmp := Round{
				blue:  0,
				green: 0,
				red:   0,
			}

			for _, color := range colors {
				dice := strings.Split(color, " ")
				numberOf, _ := strconv.Atoi(dice[1])
				colorOfDice := dice[2]

				switch colorOfDice {
				case "red":
					tmp.red = numberOf
				case "blue":
					tmp.blue = numberOf
				case "green":
					tmp.green = numberOf
				}
			}

			game.addRound(tmp)
		}

		games = append(games, game)
	}

	return games
}

func solveEx01Day02(input []string) {
	numberOfBlueDiceInBag := 14
	numberOfRedDiceInBag := 12
	numberOfGreenDiceInBag := 13

	games := parseGames(input)

	sumOfGameIds := 0
	for _, game := range games {
		allRoundsArePossible := true
		for _, round := range game.rounds {
			if round.blue > numberOfBlueDiceInBag ||
				round.red > numberOfRedDiceInBag || round.green > numberOfGreenDiceInBag {
				allRoundsArePossible = false
			}
		}

		if allRoundsArePossible {
			sumOfGameIds += game.gameId
		}
	}

	println(sumOfGameIds)
}

func solveEx02Day02(input []string) {
	games := parseGames(input)

	sumOfPowers := 0
	for _, game := range games {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, round := range game.rounds {
			if round.red > 0 && round.red > maxRed {
				maxRed = round.red
			}

			if round.blue > 0 && round.blue > maxBlue {
				maxBlue = round.blue
			}

			if round.green > 0 && round.green > maxGreen {
				maxGreen = round.green
			}
		}

		sumOfPowers += maxRed * maxGreen * maxBlue
	}

	println(sumOfPowers)
}

func main() {
	var input = ReadFileIntoArray("resources/dec_02/ex_01.txt")

	solveEx01Day02(input)
	solveEx02Day02(input)
}
