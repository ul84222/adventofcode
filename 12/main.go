package main

import (
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("Cost V1:", regions.cost())
	fmt.Println("Cost V2:", regions.costV2())
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

func (r Regions) costV2() int {
	cost := 0
	for _, regions := range r.regions {
		for _, region := range regions {
			cost += region.costV2()
		}
	}
	return cost
}

func (r Region) cost() int {
	return r.perimeter() * len(r.plots)
}

func (r Region) costV2() int {
	return r.sidesCount() * len(r.plots)
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

func (r Region) sidesCount() int {
	sides := make(map[string]bool)
	hist := make(map[string]bool)
	mark := func(direction string, i int, j int) bool {
		key := ""
		rail := 0
		switch direction {
		case "u":
			key = "u" + strconv.Itoa(i)
			rail = j
		case "d":
			key = "d" + strconv.Itoa(i)
			rail = j
		case "l":
			key = "l" + strconv.Itoa(j)
			rail = i
		case "r":
			key = "r" + strconv.Itoa(j)
			rail = i
		}

		sides[key] = true
		hist[key+string(rail)] = true
		return !hist[key+string(rail-1)] && !hist[key+string(rail+1)]
	}
	count := 0
	maxI := 0
	maxJ := 0
	for pair, _ := range r.plots {
		if pair.i > maxI {
			maxI = pair.i
		}
		if pair.j > maxJ {
			maxJ = pair.j
		}
	}
	for i := 0; i <= maxI; i++ {
		for j := 0; j <= maxJ; j++ {
			pair := Pair{i, j}
			if !r.plots[pair] {
				continue
			}

			if !r.plots[Pair{i - 1, j}] {
				if mark("u", i, j) {
					count++
				}
			}
			if !r.plots[Pair{i, j - 1}] {
				if mark("l", i, j) {
					count++
				}
			}
			if !r.plots[Pair{i + 1, j}] {
				if mark("d", i, j) {
					count++
				}
			}
			if !r.plots[Pair{i, j + 1}] {
				if mark("r", i, j) {
					count++
				}
			}
		}
	}
	return count
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
