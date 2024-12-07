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

	equations, err := parseEquations(input)
	if err != nil {
		panic(err)
	}

	result := 0
	for _, equation := range equations {
		if equation.IsValid() {
			result += equation.TestValue
		}
	}
	fmt.Println("Result: ", result)

}

type Equation struct {
	TestValue int
	Numbers   []int
}

func (e *Equation) IsValid() bool {
	return verifyEquation(e.Numbers[0], e.Numbers[1:], e.TestValue)

}
func verifyEquation(left int, right []int, target int) bool {
	if len(right) == 0 {
		return left == target
	}
	return verifyEquation(left*right[0], right[1:], target) || verifyEquation(left+right[0], right[1:], target)
}

func parseEquations(input string) ([]Equation, error) {
	lines := strings.Split(input, "\n")
	equations := make([]Equation, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("cant parse line: line=%s", line)
		}

		rightParts := strings.Fields(parts[1])
		numbers := make([]int, 0, len(rightParts))
		for _, rightPart := range rightParts {
			if val, err := strconv.Atoi(rightPart); err != nil {
				return nil, err
			} else {
				numbers = append(numbers, val)
			}
		}

		if val, err := strconv.Atoi(parts[0]); err != nil {
			return nil, err
		} else {
			equations = append(equations, Equation{val, numbers})
		}
	}
	return equations, nil
}
