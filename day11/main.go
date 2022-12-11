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

type monkey struct {
	Items          utils.Queue
	Operation      func(int) int
	Test           func(int) int
	ModNum         int
	NumInspections int
}

func main() {
	fmt.Println("Part 1:")
	fmt.Println(part1(input))

	fmt.Println("Part 2:")
	fmt.Println(part2(input))
}

func part1(input string) int {
	monkeyDefs := strings.Split(input, "\n\n")
	monkeys := make([]monkey, 0)
	for _, m := range monkeyDefs {
		monkeys = append(monkeys, parseMonkey(m))
	}
	numRounds := 20
	for round := 0; round < numRounds; round++ {
		for i, monkey := range monkeys {
			for {
				item, hasItem := monkey.Items.Dequeue()
				if !hasItem {
					break
				}
				monkeys[i].NumInspections = monkeys[i].NumInspections + 1
				// Apply func
				itemVal := monkey.Operation(item)
				// Divide val
				divVal := itemVal / 3
				monkeys[monkey.Test(divVal)].Items.Enqueue(divVal)
			}
		}
	}
	return getMonkeyBusinessLevel(monkeys)
}

func part2(input string) int {
	monkeyDefs := strings.Split(input, "\n\n")
	monkeys := make([]monkey, 0)
	for _, m := range monkeyDefs {
		monkeys = append(monkeys, parseMonkey(m))
	}
	mod := 1
	for _, m := range monkeys {
		mod *= m.ModNum
	}
	numRounds := 10000
	for round := 0; round < numRounds; round++ {
		for i, monkey := range monkeys {
			for {
				item, hasItem := monkey.Items.Dequeue()
				if !hasItem {
					break
				}
				monkeys[i].NumInspections = monkeys[i].NumInspections + 1
				// Apply func
				itemVal := monkey.Operation(item)
				modVal := itemVal % mod
				monkeys[monkey.Test(modVal)].Items.Enqueue(modVal)
			}
		}
	}
	return getMonkeyBusinessLevel(monkeys)
}

func parseMonkey(s string) monkey {
	lines := strings.Split(s, "\n")
	opFunc := parseOperation(lines[2])
	testFunc, modVal := parseTest(lines[3:6])
	m := monkey{
		*utils.NewQueue(1000),
		opFunc,
		testFunc,
		modVal,
		0,
	}
	// Starting items
	items := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
	for _, e := range items {
		item, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		m.Items.Enqueue(item)
	}
	return m
}

func parseOperation(s string) func(int) int {
	toks := strings.Split(s, " ")
	// Only interested in the last two
	ops := toks[len(toks)-2:]
	if ops[1] == "old" {
		if ops[0] == "+" {
			return func(x int) int {
				return x + x
			}
		} else if ops[0] == "*" {
			return func(x int) int {
				return x * x
			}
		} else {
			panic("Unknown op")
		}
	} else {
		val, err := strconv.Atoi(ops[1])
		if err != nil {
			panic(err)
		}
		if ops[0] == "+" {
			return func(x int) int {
				return x + val
			}
		} else if ops[0] == "*" {
			return func(x int) int {
				return x * val
			}
		} else {
			panic("Unknown op")
		}
	}
}

func parseTest(s []string) (func(int) int, int) {
	divToks := strings.Split(s[0], " ")
	divVal, err := strconv.Atoi(divToks[len(divToks)-1])
	if err != nil {
		panic(err)
	}
	trueToks := strings.Split(s[1], " ")
	trueMonkey, err := strconv.Atoi(trueToks[len(trueToks)-1])
	if err != nil {
		panic(err)
	}
	falseToks := strings.Split(s[2], " ")
	falseMonkey, err := strconv.Atoi(falseToks[len(falseToks)-1])
	if err != nil {
		panic(err)
	}
	return func(x int) int {
		if x%divVal == 0 {
			return trueMonkey
		}
		return falseMonkey
	}, divVal
}

func getMonkeyBusinessLevel(monkeys []monkey) int {
	top1 := 0
	top2 := 0
	for _, monkey := range monkeys {
		insp := monkey.NumInspections
		if insp > top2 {
			if insp > top1 {
				top1, top2 = insp, top1
			} else {
				top2 = insp
			}
		}
	}
	return top1 * top2
}
