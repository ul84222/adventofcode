package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parse(input)

	count := 0
	for _, d := range puzzle.designs {
		mem := make(Mem)
		count += possibleWays(puzzle, d, mem)
	}
	fmt.Println("Found: ", count)
}

func possibleWays(p Puzzle, design []rune, mem Mem) int {
	if len(design) == 0 {
		return 1
	}

	if mem[string(design)] != 0 {
		return mem[string(design)] - 1
	}

	applicable := p.towels[design[0]]
	if applicable == nil {
		mem[string(design)] = 1
		return 0
	}

	count := 0
	for _, it := range applicable {
		if !isPrefix(design, it) {
			continue
		}

		count += possibleWays(p, design[len(it):], mem)
	}
	mem[string(design)] = count + 1
	return count
}

func isPrefix(base []rune, prefix []rune) bool {
	if len(prefix) > len(base) {
		return false
	}

	for i, it := range prefix {
		if it != base[i] {
			return false
		}
	}
	return true
}

type Mem map[string]int

type Towel []rune

type Puzzle struct {
	towels  map[rune][]Towel
	designs [][]rune
}

func parse(input string) Puzzle {
	parts := strings.Split(input, "\n")

	towels := make(map[rune][]Towel)
	for _, it := range strings.Split(parts[0], ",") {
		towel := Towel(strings.TrimSpace(it))
		towels[towel[0]] = append(towels[towel[0]], towel)
	}

	designs := [][]rune{}
	for i := 2; i < len(parts); i++ {
		if parts[i] == "" {
			continue
		}
		designs = append(designs, []rune(parts[i]))
	}
	return Puzzle{towels, designs}
}
