package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/avidlaud/adventofcode2022/go/utils"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:")
	fmt.Println(part1(input))

	fmt.Println("Part 2:")
	fmt.Println(part2(input))
}

func part1(input string) int {
	importantCycles := []int{20, 60, 100, 140, 180, 220}
	lines := strings.Split(input, "\n")
	cycles := make([]int, 0)
	regVal := 1
	for _, line := range lines {
		toks := strings.Split(line, " ")
		if len(toks) == 2 {
			cycles = append(cycles, regVal)
			cycles = append(cycles, regVal)
			addend, err := strconv.Atoi(toks[1])
			if err != nil {
				panic(err)
			}
			regVal += addend
		} else {
			cycles = append(cycles, regVal)
		}
	}
	cycles = append(cycles, regVal)
	score := 0
	for _, cycle := range importantCycles {
		score += cycle * cycles[cycle-1]
	}
	return score
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	cycles := make([]int, 0)
	regVal := 1
	for _, line := range lines {
		toks := strings.Split(line, " ")
		if len(toks) == 2 {
			cycles = append(cycles, regVal)
			cycles = append(cycles, regVal)
			addend, err := strconv.Atoi(toks[1])
			if err != nil {
				panic(err)
			}
			regVal += addend
		} else {
			cycles = append(cycles, regVal)
		}
	}
	cycles = append(cycles, regVal)
	output := ""
	for i := 0; i < 240; i++ {
		hPos := i % 40
		spritePos := cycles[i] % 40
		if utils.Abs(spritePos-hPos) <= 1 {
			output += "#"
		} else {
			output += "."
		}
		if hPos == 39 {
			output += "\n"
		}
	}
	return output
}
