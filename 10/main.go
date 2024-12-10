package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	I int
	J int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle, err := parsePuzzle(input)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	resultV1 := 0
	resultV2 := 0
	dest := make(map[Point]int)
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == 0 {
				clear(dest)
				countTrails(puzzle, i, j, 0, dest)
				resultV1 += len(dest)

				for _, cnt := range dest {
					resultV2 += cnt
				}
			}
		}
	}
	spent := time.Since(start)
	fmt.Printf("result V1: %d\n", resultV1)
	fmt.Printf("result V2: %d\n", resultV2)
	fmt.Printf("spent: %s\n", spent)
}

func countTrails(puzzle [][]int, i int, j int, expected int, dest map[Point]int) {
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) {
		return
	}
	curr := puzzle[i][j]
	if curr != expected {
		return
	}
	if curr == 9 {
		dest[Point{i, j}] += 1
		return
	}
	countTrails(puzzle, i-1, j, expected+1, dest)
	countTrails(puzzle, i+1, j, expected+1, dest)
	countTrails(puzzle, i, j-1, expected+1, dest)
	countTrails(puzzle, i, j+1, expected+1, dest)
}

func parsePuzzle(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	puzzle := make([][]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		row := []int{}
		for _, el := range line {
			if num, err := strconv.Atoi(string(el)); err != nil {
				return nil, err
			} else {
				row = append(row, num)
			}
		}
		puzzle = append(puzzle, row)
	}

	return puzzle, nil
}
