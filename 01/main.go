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
		panic("slies have different lengths")
	}

	slices.Sort(l)
	slices.Sort(r)

	fmt.Println("l", l)
	fmt.Println("r", r)

	result := 0
	for i := 0; i < len(l); i++ {
		diff := l[i] - r[i]
		if diff > 0 {
			result += diff
		} else {
			result -= diff
		}
	}
	fmt.Println("Result: ", result)
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
			return nil, nil, err
		}

		r, err = parseAndAppend(parts[1], r)
		if err != nil {
			return nil, nil, err
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
