package main

import (
	_ "embed"
	"fmt"
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
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}
	for i, row := range lines {
		for j, c := range row {
			grid[i][j] = int(c - '0')
		}
	}
	count := 0
	for i, row := range lines {
		for j := range row {
			if checkVisible(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}
	for i, row := range lines {
		for j, c := range row {
			grid[i][j] = int(c - '0')
		}
	}
	maxScenicScore := 0
	for i, row := range lines {
		for j := range row {
			maxScenicScore = utils.Max(calculateScenicScore(grid, i, j), maxScenicScore)

		}
	}
	return maxScenicScore
}

func checkVisible(grid [][]int, i, j int) bool {
	return checkLeft(grid, i, j) || checkDown(grid, i, j) || checkRight(grid, i, j) || checkUp(grid, i, j)
}

func checkLeft(grid [][]int, i, j int) bool {
	currHeight := grid[i][j]
	for ptr := j - 1; ptr >= 0; ptr-- {
		if grid[i][ptr] >= currHeight {
			return false
		}
	}
	return true
}

func checkDown(grid [][]int, i, j int) bool {
	currHeight := grid[i][j]
	for ptr := i + 1; ptr < len(grid); ptr++ {
		if grid[ptr][j] >= currHeight {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, i, j int) bool {
	currHeight := grid[i][j]
	for ptr := j + 1; ptr < len(grid[0]); ptr++ {
		if grid[i][ptr] >= currHeight {
			return false
		}
	}
	return true
}

func checkUp(grid [][]int, i, j int) bool {
	currHeight := grid[i][j]
	for ptr := i - 1; ptr >= 0; ptr-- {
		if grid[ptr][j] >= currHeight {
			return false
		}
	}
	return true
}

func calculateScenicScore(grid [][]int, i, j int) int {
	return distanceLeft(grid, i, j) * distanceDown(grid, i, j) * distanceRight(grid, i, j) * distanceUp(grid, i, j)
}

func distanceLeft(grid [][]int, i, j int) int {
	dist := 0
	for ptr := j - 1; ptr >= 0; ptr-- {
		dist++
		if grid[i][ptr] >= grid[i][j] {
			return dist
		}
	}
	return dist
}

func distanceDown(grid [][]int, i, j int) int {
	dist := 0
	for ptr := i + 1; ptr < len(grid); ptr++ {
		dist++
		if grid[ptr][j] >= grid[i][j] {
			return dist
		}
	}
	return dist
}

func distanceRight(grid [][]int, i, j int) int {
	dist := 0
	for ptr := j + 1; ptr < len(grid[0]); ptr++ {
		dist++
		if grid[i][ptr] >= grid[i][j] {
			return dist
		}
	}
	return dist
}

func distanceUp(grid [][]int, i, j int) int {
	dist := 0
	for ptr := i - 1; ptr >= 0; ptr-- {
		dist++
		if grid[ptr][j] >= grid[i][j] {
			return dist
		}
	}
	return dist
}
