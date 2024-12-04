package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	levels, err := parse(input)
	if err != nil {
		panic(err)
	}

	safeLavels := 0
	for _, level := range levels {
		if safe, error := isSafeLevel(level); error != nil {
			panic(error)
		} else if safe {
			safeLavels++
		}
	}
	fmt.Println("Result: ", safeLavels)
}

func isSafeLevel(level []int) (bool, error) {
	safe, unsafeCell, err := findUnsafeCell(level)
	if safe {
		return safe, err
	}

	safe, _, err = findUnsafeCell(cut(level, 0))
	if safe {
		return safe, err
	}

	safe, _, err = findUnsafeCell(cut(level, unsafeCell-1))
	if safe {
		return safe, err
	}

	safe, _, err = findUnsafeCell(cut(level, unsafeCell))
	return safe, err
}

func findUnsafeCell(level []int) (bool, int, error) {
	if len(level) == 0 {
		return false, 0, fmt.Errorf("unexpected level length. length must not be zero")
	}
	if len(level) == 1 {
		return true, 0, nil
	}

	ascending := true
	if level[0] > level[1] {
		ascending = false
	}

	for i := 1; i < len(level); i++ {
		var diff int
		if ascending {
			diff = level[i] - level[i-1]
		} else {
			diff = level[i-1] - level[i]
		}

		if diff < 1 {
			return false, i, nil
		}
		if diff > 3 {
			return false, i, nil
		}
	}
	return true, 0, nil
}

func cut(slice []int, ind int) []int {
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[0:ind]...)
	result = append(result, slice[ind+1:]...)
	return result
}

func parse(input string) ([][]int, error) {
	var result [][]int
	var err error

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		parsedLine := make([]int, 0, len(parts))
		for _, part := range parts {
			if parsedLine, err = parseAndAppend(part, parsedLine); err != nil {
				return nil, fmt.Errorf("error parsing line %q: %w", line, err)
			}
		}
		result = append(result, parsedLine)
	}
	return result, nil
}

func parseAndAppend(part string, slice []int) ([]int, error) {
	number, err := strconv.Atoi(part)
	if err != nil {
		return nil, err
	}
	return append(slice, number), nil
}
