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
		if isPossible(puzzle, d, mem) {
			count++
		}
	}
	fmt.Println("Found: ", count)
}

func isPossible(p Puzzle, design []rune, mem Mem) bool {
	if len(design) == 0 {
		return true
	}

	if mem[string(design)] {
		return false
	}

	applicable := p.towels[design[0]]
	if applicable == nil {
		mem[string(design)] = true
		return false
	}

	for _, it := range applicable {
		if !isPrefix(design, it) {
			continue
		}

		if isPossible(p, design[len(it):], mem) {
			return true
		}
	}
	mem[string(design)] = true
	return false
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

type Mem map[string]bool

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
