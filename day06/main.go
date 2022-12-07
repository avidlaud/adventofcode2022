package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	for i := 0; i < len(input)-4; i++ {
		if allCharactersUnique(input[i : i+4]) {
			return i + 4
		}
	}
	return -1
}

func part2(input string) int {
	for i := 0; i < len(input)-14; i++ {
		if allCharactersUnique(input[i : i+14]) {
			return i + 14
		}
	}
	return -1
}

func allCharactersUnique(s string) bool {
	letters := map[rune]bool{}
	for _, r := range s {
		if letters[r] {
			return false
		}
		letters[r] = true
	}
	return true
}
