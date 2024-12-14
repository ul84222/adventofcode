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
	positions := make(map[Pair]int)
	for _, puzzle := range puzzle {
		p := puzzle.findFuturePosition(100)
		q := p.quadrant()
		quadrants[q]++
		positions[p]++
	}
	result := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]

	// for y := 0; y < spaceHigh; y++ {
	// 	for x := 0; x < spaceWidth; x++ {
	// 		p := Pair{x, y}
	// 		if p.quadrant() == 0 {
	// 			fmt.Printf(" ")
	// 		} else if positions[p] != 0 {
	// 			// fmt.Printf("%d", positions[p])
	// 			fmt.Printf("%d", p.quadrant())
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	elapsed := time.Since(now)
	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Result: ", result)
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
