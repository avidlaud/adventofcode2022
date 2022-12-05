package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Stack []string

func main() {
	trimmedInput := strings.TrimSuffix(input, "\n")
	fmt.Println("Part 1:")
	fmt.Println(part1(trimmedInput))

	fmt.Println("Part 2:")
	fmt.Println(part2(trimmedInput))
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	boxes := parseBoxes(lines[:8])
	movedBoxes := performInstructions(lines[10:], boxes)
	topStr := ""
	for _, box := range movedBoxes {
		hasTop, top := box.Pop()
		if hasTop {
			topStr += top
		} else {
			topStr += " "
		}
	}
	return topStr
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	boxes := parseBoxes(lines[:8])
	movedBoxes := performGroupedInstructions(lines[10:], boxes)
	topStr := ""
	for _, box := range movedBoxes {
		hasTop, top := box.Pop()
		if hasTop {
			topStr += top
		} else {
			topStr += " "
		}
	}
	return topStr
}

func (s *Stack) Push(box string) {
	*s = append(*s, box)
}

func (s *Stack) Pop() (bool, string) {
	if s.IsEmpty() {
		return false, ""
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return true, top
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func parseBoxes(lines []string) []Stack {
	stacks := make([]Stack, 9)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < 9; j++ {
			idx := (j * 4) + 1
			val := string(line[idx])
			if val != " " {
				stacks[j].Push(string(line[idx]))
			}
		}
	}
	return stacks
}

func performInstructions(lines []string, boxes []Stack) []Stack {
	for _, line := range lines {
		toks := strings.Split(line, " ")
		count, _ := strconv.Atoi(toks[1])
		src, _ := strconv.Atoi(toks[3])
		dest, _ := strconv.Atoi(toks[5])
		for i := 0; i < count; i++ {
			hasBox, pulledBox := boxes[src-1].Pop()
			if hasBox {
				boxes[dest-1].Push(pulledBox)
			}
		}
	}
	return boxes
}

func performGroupedInstructions(lines []string, boxes []Stack) []Stack {
	for _, line := range lines {
		toks := strings.Split(line, " ")
		count, _ := strconv.Atoi(toks[1])
		src, _ := strconv.Atoi(toks[3])
		dest, _ := strconv.Atoi(toks[5])
		buffer := make(Stack, 0)
		for i := 0; i < count; i++ {
			hasBox, pulledBox := boxes[src-1].Pop()
			if hasBox {
				buffer.Push(pulledBox)
			}
		}
		for !buffer.IsEmpty() {
			_, item := buffer.Pop()
			boxes[dest-1].Push(item)
		}
	}
	return boxes
}
