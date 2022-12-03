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
	lines := strings.Split(input, "\n")
	prioritySum := 0

	for _, line := range lines {
		prioritySum += getCommonItem(line)
	}
	return prioritySum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	badgeSum := 0
	for i := 0; i < len(lines)/3; i++ {
		firstBitset := rucksackBitset(lines[i*3])
		secondBitset := rucksackBitset(lines[i*3+1])
		thirdBitset := rucksackBitset(lines[i*3+2])

		commonBitset := firstBitset & secondBitset & thirdBitset

		badgeSum += getIntFromBitset(commonBitset)
	}
	return badgeSum
}

func getCommonItem(line string) int {
	firstHalf := line[:len(line)/2]
	secondHalf := line[len(line)/2:]

	firstHalfBitset := rucksackBitset(firstHalf)
	secondHalfBitset := rucksackBitset(secondHalf)

	commonBitset := firstHalfBitset & secondHalfBitset
	return getIntFromBitset(commonBitset)
}

func rucksackBitset(line string) int64 {
	bitset := int64(0)
	for _, r := range []rune(line) {
		bitset = bitset | getBitPosition(runeToInt(r))
	}
	return bitset
}

func getBitPosition(pos int) int64 {
	return int64(1) << (pos - 1)
}

func getIntFromBitset(bitset int64) int {
	pos := 0
	for bitset != 0 {
		bitset = bitset >> 1
		pos++
	}
	return pos
}

func runeToInt(r rune) int {
	x := int(r)
	a := int(rune('a'))
	A := int(rune('A'))
	if x < a {
		return x - A + 1 + 26
	}
	return x - a + 1
}
