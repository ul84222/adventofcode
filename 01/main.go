package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	l, r, err := parse(input)
	if err != nil {
		panic(err)
	}

	if len(l) != len(r) {
		panic("slises have different lengths")
	}

	slices.Sort(l)
	slices.Sort(r)

	result := 0
	for i := 0; i < len(l); i++ {
		diff := l[i] - r[i]
		if diff > 0 {
			result += diff
		} else {
			result -= diff
		}
	}
	fmt.Println("Part 1: result: ", result)

	score := 0
	indexL := 0
	indexR := 0
	appearCount := 0
	for indexL < len(l) && indexR < len(r) {
		if l[indexL] == r[indexR] {
			appearCount++
			indexR++
		} else {
			score += multiplier(indexL, l) * appearCount * l[indexL]
			appearCount = 0
			if l[indexL] > r[indexR] {
				indexR++
			} else {
				indexL++
			}
		}
	}
	if appearCount != 0 {
		score += appearCount * l[indexL]
	}
	fmt.Println("Part 2: score: ", score)
}

func multiplier(index int, slice []int) int {
	multiplier := 1
	for i := index + 1; i < len(slice) && slice[i] == slice[index]; i++ {
		multiplier++
	}
	return multiplier
}

func parse(input string) ([]int, []int, error) {
	var l []int
	var r []int
	var err error

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("unexpected line: %s", line)
		}

		l, err = parseAndAppend(parts[0], l)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing line %q: %w", line, err)
		}

		r, err = parseAndAppend(parts[1], r)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing line %q: %w", line, err)
		}
	}
	return l, r, nil
}

func parseAndAppend(part string, slice []int) ([]int, error) {
	number, err := strconv.Atoi(part)
	if err != nil {
		return nil, err
	}
	return append(slice, number), nil
}
