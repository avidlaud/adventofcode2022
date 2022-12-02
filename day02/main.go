package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/avidlaud/adventofcode2022/go/utils"
)

//go:embed input.txt
var input string

type Round struct {
	opp, you Shape
}

const lostScore = 0
const drawScore = 3
const winScore = 6

type Shape int

const (
	Rock     Shape = 1
	Paper          = 2
	Scissors       = 3
)

var roundScores = map[Round]int{
	{Rock, Rock}:         drawScore,
	{Rock, Paper}:        winScore,
	{Rock, Scissors}:     lostScore,
	{Paper, Rock}:        lostScore,
	{Paper, Paper}:       drawScore,
	{Paper, Scissors}:    winScore,
	{Scissors, Rock}:     winScore,
	{Scissors, Paper}:    lostScore,
	{Scissors, Scissors}: drawScore,
}

var shapeDictionary = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func main() {
	trimmedInput := strings.TrimSuffix(input, "\n")
	fmt.Println("Part 1:")
	fmt.Println(part1(trimmedInput))

	fmt.Println("Part 2:")
	fmt.Println(part2(trimmedInput))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	rounds := utils.MapSliceNoErr(lines, lineToRound)
	totalScore := 0
	for _, round := range rounds {
		// Score from round outcome
		totalScore += roundScores[round]
		// Base score from shape
		totalScore += int(round.you)
	}
	return totalScore
}

func lineToRound(line string) Round {
	moves := strings.Split(line, " ")
	return Round{shapeDictionary[moves[0]], shapeDictionary[moves[1]]}
}

func lineWithInstructionToRound(line string) Round {
	toks := strings.Split(line, " ")
	oppMove := shapeDictionary[toks[0]]
	instruction := toks[1]
	var myMove Shape
	switch oppMove {
	case Rock:
		switch instruction {
		case "X":
			myMove = Scissors
		case "Y":
			myMove = Rock
		case "Z":
			myMove = Paper
		}
	case Paper:
		switch instruction {
		case "X":
			myMove = Rock
		case "Y":
			myMove = Paper
		case "Z":
			myMove = Scissors
		}
	case Scissors:
		switch instruction {
		case "X":
			myMove = Paper
		case "Y":
			myMove = Scissors
		case "Z":
			myMove = Rock

		}
	}
	return Round{oppMove, myMove}
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	rounds := utils.MapSliceNoErr(lines, lineWithInstructionToRound)
	totalScore := 0
	for _, round := range rounds {
		// Score from round outcome
		totalScore += roundScores[round]
		// Base score from shape
		totalScore += int(round.you)
	}
	return totalScore
}
