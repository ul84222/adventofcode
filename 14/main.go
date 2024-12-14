package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	spaceWidth = 101
	spaceHigh  = 103
	// spaceWidth = 11
	// spaceHigh  = 7
)

type Pair struct {
	x int
	y int
}

type Robot struct {
	position Pair
	velocity Pair
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parsePuzzle(input)
	now := time.Now()

	quadrants := []int{0, 0, 0, 0, 0}
	for _, puzzle := range puzzle {
		p := puzzle.findFuturePosition(100)
		q := p.quadrant()
		quadrants[q]++
	}
	result := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
	elapsed := time.Since(now)
	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Result: ", result)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 100000; i++ {
		positions := make(map[Pair]bool)

		xMin := spaceWidth / 3
		xMax := spaceWidth - xMin
		yMin := spaceHigh / 3
		yMax := spaceHigh - yMin
		treshold := int(float64(len(puzzle)) * 0.50)

		count := 0
		for _, puzzle := range puzzle {
			p := puzzle.findFuturePosition(i)
			positions[p] = true

			if p.x > xMin && p.x < xMax && yMin < p.y && yMax > p.y {
				count++
			}
		}
		if count < treshold {
			continue
		}
		for y := 0; y < spaceHigh; y++ {
			for x := 0; x < spaceWidth; x++ {
				p := Pair{x, y}
				if positions[p] {
					fmt.Printf("X")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println("i:", i)
		fmt.Println()
		fmt.Println()
		scanner.Scan()
	}
}

func isPossibleTree(positions map[Pair]bool) bool {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if positions[Pair{x, y}] {
				return false
			}
			if positions[Pair{x, spaceHigh - y}] {
				return false
			}
			if positions[Pair{spaceWidth - x, y}] {
				return false
			}
			if positions[Pair{spaceWidth - x, spaceHigh - y}] {
				return false
			}
		}
	}
	return true
}

func (r *Robot) findFuturePosition(steps int) Pair {
	x := r.velocity.x*steps + r.position.x
	y := r.velocity.y*steps + r.position.y

	resolve := func(x, size int) int {
		pos := x % size
		if pos >= 0 {
			return pos
		}
		return size + pos
	}

	return Pair{resolve(x, spaceWidth), resolve(y, spaceHigh)}
}

func (p *Pair) quadrant() int {
	medianX := spaceWidth / 2
	medianY := spaceHigh / 2

	switch {
	case p.x == medianX || p.y == medianY:
		return 0
	case p.x > medianX && p.y < medianY:
		return 1
	case p.x < medianX && p.y < medianY:
		return 2
	case p.x < medianX && p.y > medianY:
		return 3
	case p.x > medianX && p.y > medianY:
		return 4
	default:
		panic(fmt.Errorf("unexpected state: %v", p))
	}
}

func parsePuzzle(input string) []Robot {
	lines := strings.Split(input, "\n")
	puzzle := make([]Robot, 0, len(lines))
	rg := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	atoi := func(str string) int {
		if val, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			return val
		}
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := rg.FindStringSubmatch(line)
		robot := Robot{
			Pair{atoi(parts[1]), atoi(parts[2])},
			Pair{atoi(parts[3]), atoi(parts[4])},
		}
		puzzle = append(puzzle, robot)
	}
	return puzzle
}
