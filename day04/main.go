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

type Range struct {
	start, end, length int
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
	count := 0
	for _, line := range lines {
		elves := strings.Split(line, ",")
		ranges := utils.MapSliceNoErr(elves, parseRange)
		var longerRange Range
		var shorterRange Range
		if ranges[0].length >= ranges[1].length {
			longerRange = ranges[0]
			shorterRange = ranges[1]
		} else {
			longerRange = ranges[1]
			shorterRange = ranges[0]
		}
		if longerRange.start <= shorterRange.start && longerRange.end >= shorterRange.end {
			count += 1
		}
	}
	return count
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		elves := strings.Split(line, ",")
		ranges := utils.MapSliceNoErr(elves, parseRange)
		var earlierRange Range
		var laterRange Range
		if ranges[0].start <= ranges[1].start {
			earlierRange = ranges[0]
			laterRange = ranges[1]
		} else {
			earlierRange = ranges[1]
			laterRange = ranges[0]
		}
		if earlierRange.end >= laterRange.start {
			count += 1
		}
	}
	return count
}

func parseRange(s string) Range {
	toks := strings.Split(s, "-")
	nums := utils.MapSlice(toks, strconv.Atoi)
	return Range{nums[0], nums[1], nums[1] - nums[0]}
}
