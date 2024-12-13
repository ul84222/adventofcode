package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	costA = 3
	costB = 1
	extra = 10000000000000
)

type Pair struct {
	x int
	y int
}

type Puzzle struct {
	buttonA Pair
	buttonB Pair
	prize   Pair
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parsePuzzle(input)

	start := time.Now()
	result := 0
	for _, p := range puzzle {
		result += p.solve()
	}
	elapsed := time.Since(start)
	fmt.Println("Result: ", result)
	fmt.Println("Elapsed: ", elapsed)
}

func (p *Puzzle) solve() int {
	min := func(l int, r int) int {
		if l > r {
			return r
		}
		return l
	}

	bCount := min(p.prize.x/p.buttonB.x, p.prize.y/p.buttonB.y)
	for bCount >= 0 {
		xRemaining := p.prize.x - (bCount * p.buttonB.x)
		yRemaining := p.prize.y - (bCount * p.buttonB.y)

		xDiff := xRemaining % p.buttonA.x
		yDiff := yRemaining % p.buttonA.y
		aCountX := xRemaining / p.buttonA.x
		aCountY := yRemaining / p.buttonA.y
		if xDiff == 0 && yDiff == 0 {
			if aCountX == aCountY {
				return costB*bCount + costA*xRemaining/p.buttonA.x
			}
		}

		aCountDiff := max(aCountX, aCountY) - min(aCountX, aCountY)
		xStep := (aCountDiff * p.buttonA.x) / p.buttonB.x
		yStep := (aCountDiff * p.buttonA.y) / p.buttonB.y
		step := min(xStep, yStep)
		if step == 0 {
			step = 1
		}
		bCount -= step
	}
	return 0
}

func parsePuzzle(input string) []Puzzle {
	lines := strings.Split(input, "\n")
	pattern := regexp.MustCompile(`X.(\d+), Y.(\d+)`)
	parse := func(parts []string) Pair {
		x, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		return Pair{x, y}
	}

	puzzles := []Puzzle{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		a := pattern.FindStringSubmatch(lines[i])
		b := pattern.FindStringSubmatch(lines[i+1])
		prize := pattern.FindStringSubmatch(lines[i+2])
		prizePair := parse(prize)
		puzzles = append(puzzles, Puzzle{parse(a), parse(b), Pair{prizePair.x + extra, prizePair.y + extra}})
		i += 3
	}
	return puzzles
}
