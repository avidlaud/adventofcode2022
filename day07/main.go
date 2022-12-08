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

type Directory struct {
	subDirs     []string
	fileSizeSum int
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
	var stack utils.Stack[string]
	dirSizes := make(map[string]int)
	currentDir := ""
	for _, line := range lines {
		if line == "$ cd .." {
			stack.Pop()
		} else if strings.HasPrefix(line, "$ cd") {
			toks := strings.Split(line, " ")
			stack.Push(toks[2])
			currentDir = strings.Join(stack.Vals, "/")
		} else if dirSize, err := strconv.Atoi(strings.Split(line, " ")[0]); err == nil {
			dirSizes[currentDir] += dirSize
			for _, subDir := range getSubdirs(currentDir) {
				dirSizes[subDir] += dirSize
			}
		}
	}
	totalSum := 0
	for _, s := range dirSizes {
		if s < 100000 {
			totalSum += s
		}
	}
	return totalSum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	var stack utils.Stack[string]
	dirSizes := make(map[string]int)
	currentDir := ""
	for _, line := range lines {
		if line == "$ cd .." {
			stack.Pop()
		} else if strings.HasPrefix(line, "$ cd") {
			toks := strings.Split(line, " ")
			stack.Push(toks[2])
			currentDir = strings.Join(stack.Vals, "/")
		} else if dirSize, err := strconv.Atoi(strings.Split(line, " ")[0]); err == nil {
			dirSizes[currentDir] += dirSize
			for _, subDir := range getSubdirs(currentDir) {
				dirSizes[subDir] += dirSize
			}
		}
	}
	target := dirSizes["/"] - 40000000
	candidate := 999999999
	for _, s := range dirSizes {
		if s < candidate && s > target {
			candidate = s
		}
	}
	return candidate
}

func getSubdirs(currDir string) []string {
	toks := strings.Split(currDir, "/")
	var out []string
	for i := range toks {
		if i != 1 {
			out = append(out, strings.Join(toks[:i], "/"))
		}
	}
	return out
}
