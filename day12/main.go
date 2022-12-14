package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/avidlaud/adventofcode2022/go/utils"
)

//go:embed input.txt
var input string

type coord struct {
	x, y int
}

const (
	start = -1
	end   = 26
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(part1(input))

	fmt.Println("Part 2:")
	fmt.Println(part2(input))
}

func part1(input string) int {
	m := parseMap(input, false)
	startCoord := getStartCoord(m)
	endCoord := getEndCoord(m)
	q := utils.NewQueue[coord](10000)
	q.Enqueue(startCoord)
	visited := map[coord]int{}
	visited[startCoord] = 0
	for {
		node := q.Deq()
		if node == endCoord {
			break
		}
		for _, move := range getMoves(m, node) {
			if _, ok := visited[move]; !ok {
				visited[move] = visited[node] + 1
				q.Enqueue(move)
			}
		}
	}

	return visited[endCoord]
}

func part2(input string) int {
	m := parseMap(input, true)
	startCoords := getAllACoords(m)
	endCoord := getEndCoord(m)
	shortestPath := 9999999999999
	for _, startCoord := range startCoords {
		q := utils.NewQueue[coord](10000)
		q.Enqueue(startCoord)
		visited := map[coord]int{}
		visited[startCoord] = 0
		for {
			node, hasItem := q.Dequeue()
			if !hasItem {
				visited[endCoord] = 999999
				break
			}
			if node == endCoord {
				break
			}
			for _, move := range getMoves(m, node) {
				if _, ok := visited[move]; !ok {
					visited[move] = visited[node] + 1
					q.Enqueue(move)
				}
			}
		}
		shortestPath = utils.Min(shortestPath, visited[endCoord])
	}
	return shortestPath
}

func parseMap(s string, isPartTwo bool) [][]int {
	lines := strings.Split(s, "\n")
	m := make([][]int, len(lines))
	for i := range m {
		m[i] = make([]int, len(lines[0]))
	}
	for i, line := range lines {
		for j, c := range line {
			m[i][j] = runeToInt(c, isPartTwo)
		}
	}
	return m
}

func runeToInt(r rune, isPartTwo bool) int {
	if r == rune('E') {
		return end
	}
	if r == rune('S') {
		if isPartTwo {
			return 0
		}
		return start
	}
	return int(r - 'a')
}

func getStartCoord(m [][]int) coord {
	return getCoordOfUnique(m, start)
}

func getEndCoord(m [][]int) coord {
	return getCoordOfUnique(m, end)
}

func getCoordOfUnique(m [][]int, val int) coord {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == val {
				return coord{i, j}
			}
		}
	}
	return coord{x: -1, y: -1}
}

func getMoves(m [][]int, pos coord) []coord {
	movement := []coord{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	}
	validMoves := []coord{}
	for _, move := range movement {
		newPos := coord{x: pos.x + move.x, y: pos.y + move.y}
		if newPos.x >= 0 && newPos.x < len(m) && newPos.y >= 0 && newPos.y < len(m[0]) {
			if m[newPos.x][newPos.y]-m[pos.x][pos.y] <= 1 {
				validMoves = append(validMoves, newPos)
			}
		}
	}
	return validMoves
}

func getAllACoords(m [][]int) []coord {
	coords := []coord{}
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 {
				coords = append(coords, coord{x: i, y: j})
			}
		}
	}
	return coords
}
