package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/avidlaud/adventofcode2022/go/utils"
)

//go:embed input.txt
var input string

func main() {
	trimmedInput := strings.TrimSuffix(input, "\n")
	fmt.Println("Part 1:")
	fmt.Println(part1(trimmedInput))

	fmt.Println("Part 2:")
	fmt.Println(part2(trimmedInput))
}

func part1(input string) int {
	return getMaxElfCalories(getElfCalories(input))
}

func part2(input string) int {
	elves := getElfCalories(input)
	sort.Ints(elves)
	top3 := elves[len(elves)-3:]
	return utils.SumSlice(top3)
}

func getElfCalories(input string) []int {
	lines := strings.Split(input, "\n\n")
	elves := make([]int, 0)
	for _, line := range lines {
		runningSum := 0
		parsedCalories := utils.MapSlice(strings.Split(line, "\n"), strconv.Atoi)
		for _, calories := range parsedCalories {
			runningSum += calories
		}
		elves = append(elves, runningSum)
	}
	return elves
}

func getMaxElfCalories(elves []int) int {
	_, max := utils.MaxIntSlice(elves)
	return max
}
