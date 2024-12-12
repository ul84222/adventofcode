package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parsePuzzle(input)

	start := time.Now()
	regions := Regions{make(map[rune][]Region)}
	for i, line := range puzzle {
		for j, plot := range line {
			pair := Pair{i, j}
			regions.chooseFor(plot, pair).add(pair, puzzle)
		}
	}
	fmt.Println("Cost:", regions.cost())
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}

type Pair struct {
	i int
	j int
}

type Region struct {
	t     rune
	plots map[Pair]bool
}

type Regions struct {
	regions map[rune][]Region
}

func (r Regions) chooseFor(t rune, pair Pair) Region {
	regions := r.regions[t]
	for _, region := range regions {
		if region.plots[Pair{pair.i - 1, pair.j}] {
			return region
		}
		if region.plots[Pair{pair.i + 1, pair.j}] {
			return region
		}
		if region.plots[Pair{pair.i, pair.j - 1}] {
			return region
		}
		if region.plots[Pair{pair.i, pair.j + 1}] {
			return region
		}
	}
	region := Region{t, make(map[Pair]bool)}
	r.regions[t] = append(regions, region)
	return region
}

func (r Region) add(pair Pair, field [][]rune) {
	if pair.i < 0 || pair.j < 0 || pair.i >= len(field) || pair.j >= len(field[pair.i]) {
		return
	}
	if r.t != field[pair.i][pair.j] {
		return
	}
	if r.plots[pair] {
		return
	}

	r.plots[pair] = true

	r.add(Pair{pair.i + 1, pair.j}, field)
	r.add(Pair{pair.i - 1, pair.j}, field)
	r.add(Pair{pair.i, pair.j + 1}, field)
	r.add(Pair{pair.i, pair.j - 1}, field)
}

func (r Regions) cost() int {
	cost := 0
	for _, regions := range r.regions {
		for _, region := range regions {
			cost += region.cost()
		}
	}
	return cost
}

func (r Region) cost() int {
	return r.perimeter() * len(r.plots)
}

func (r Region) perimeter() int {
	p := 0
	for pair, _ := range r.plots {
		if !r.plots[Pair{pair.i - 1, pair.j}] {
			p++
		}
		if !r.plots[Pair{pair.i, pair.j - 1}] {
			p++
		}
		if !r.plots[Pair{pair.i + 1, pair.j}] {
			p++
		}
		if !r.plots[Pair{pair.i, pair.j + 1}] {
			p++
		}
	}
	return p
}

func parsePuzzle(input string) [][]rune {
	lines := strings.Split(input, "\n")
	puzzle := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		puzzle = append(puzzle, []rune(line))
	}
	return puzzle
}
