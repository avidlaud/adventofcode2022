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

type Position struct {
	x, y int
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
	headPos, tailPos := Position{0, 0}, Position{0, 0}
	visited := make(map[Position]bool)
	for _, line := range lines {
		toks := strings.Split(line, " ")
		count, err := strconv.Atoi(toks[1])
		if err != nil {
			panic(err)
		}
		if toks[0] == "R" {
			for i := 0; i < count; i++ {
				moveRight(&headPos)
				follow(&headPos, &tailPos)
				visited[tailPos] = true
			}
		} else if toks[0] == "L" {
			for i := 0; i < count; i++ {
				moveLeft(&headPos)
				follow(&headPos, &tailPos)
				visited[tailPos] = true
			}
		} else if toks[0] == "U" {
			for i := 0; i < count; i++ {
				moveUp(&headPos)
				follow(&headPos, &tailPos)
				visited[tailPos] = true
			}
		} else {
			for i := 0; i < count; i++ {
				moveDown(&headPos)
				follow(&headPos, &tailPos)
				visited[tailPos] = true
			}
		}
	}
	return len(visited)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	knots := make([]Position, 10)
	visited := make(map[Position]bool)
	for _, line := range lines {
		toks := strings.Split(line, " ")
		count, err := strconv.Atoi(toks[1])
		if err != nil {
			panic(err)
		}
		if toks[0] == "R" {
			for a := 0; a < count; a++ {
				moveRight(&knots[0])
				for i := 0; i < len(knots)-1; i++ {
					follow(&knots[i], &knots[i+1])
				}
				visited[knots[len(knots)-1]] = true
			}
		} else if toks[0] == "L" {
			for a := 0; a < count; a++ {
				moveLeft(&knots[0])
				for i := 0; i < len(knots)-1; i++ {
					follow(&knots[i], &knots[i+1])
				}
				visited[knots[len(knots)-1]] = true

			}
		} else if toks[0] == "U" {
			for a := 0; a < count; a++ {
				moveUp(&knots[0])
				for i := 0; i < len(knots)-1; i++ {
					follow(&knots[i], &knots[i+1])
				}
				visited[knots[len(knots)-1]] = true
			}
		} else {
			for a := 0; a < count; a++ {
				moveDown(&knots[0])
				for i := 0; i < len(knots)-1; i++ {
					follow(&knots[i], &knots[i+1])
				}
				visited[knots[len(knots)-1]] = true
			}
		}
	}
	return len(visited)
}

func moveDown(pos *Position) {
	pos.y -= 1
}

func moveUp(pos *Position) {
	pos.y += 1
}

func moveLeft(pos *Position) {
	pos.x -= 1
}

func moveRight(pos *Position) {
	pos.x += 1
}

func follow(head, tail *Position) {
	dx := head.x - tail.x
	dy := head.y - tail.y

	if utils.Abs(dx) > 1 || utils.Abs(dy) > 1 {
		tail.x = tail.x + utils.Sign(dx)
		tail.y = tail.y + utils.Sign(dy)
	}
}
