package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Left  int
	Right int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	rules, updates, err := parse(input)
	if err != nil {
		panic(err)
	}

	resultPart1 := 0
	resultPart2 := 0
	for _, update := range updates {
		if test(update, rules) {
			median, err := median(update)
			if err != nil {
				panic(err)
			}
			resultPart1 += median
		} else {
			fixedUpdate := fixUpdate(update, rules)
			median, err := median(fixedUpdate)
			if err != nil {
				panic(err)
			}
			resultPart2 += median
		}

	}

	fmt.Println("Result part 1: ", resultPart1)
	fmt.Println("Result part 2: ", resultPart2)
}

func test(update []int, rules map[int][]Rule) bool {
	seen := make(map[int]bool)

	for _, page := range update {
		for _, rule := range rules[page] {
			if rule.Left == page && seen[rule.Right] {
				return false
			}
		}
		seen[page] = true
	}

	return true
}

func fixUpdate(update []int, rules map[int][]Rule) []int {
	seen := make(map[int]bool)
	fixed := make([]int, 0, len(update))

	shouldBeBeforePages := make(map[int]bool)
	for _, page := range update {
		clear(shouldBeBeforePages)

		for _, rule := range rules[page] {
			if rule.Left == page && seen[rule.Right] {
				shouldBeBeforePages[rule.Right] = true
			}
		}
		seen[page] = true
		fixed = insertPage(fixed, page, shouldBeBeforePages)
	}
	return fixed
}

func insertPage(dest []int, val int, before map[int]bool) []int {
	if len(before) == 0 {
		return append(dest, val)
	}

	cutInd := -1
	for i, val := range dest {
		if before[val] {
			cutInd = i
			break
		}
	}

	var result []int
	result = append(result, dest[0:cutInd]...)
	result = append(result, val)
	result = append(result, dest[cutInd:]...)
	return result
}

func median(update []int) (int, error) {
	if len(update)%2 == 0 {
		return 0, fmt.Errorf("Update has even number of pages. update=%v", update)
	}
	return update[len(update)/2], nil
}

func parse(input string) (map[int][]Rule, [][]int, error) {
	lines := strings.Split(input, "\n")
	separatorIndex := -1
	for i, line := range lines {
		if line == "" {
			separatorIndex = i
			break
		}
	}
	if separatorIndex == -1 {
		return nil, nil, fmt.Errorf("failed to read input")
	}

	rules, err := parseRules(lines[0:separatorIndex])
	if err != nil {
		return nil, nil, err
	}

	updates, err := parseUpdates(lines[separatorIndex+1:])
	if err != nil {
		return nil, nil, err
	}
	return rules, updates, nil
}

func parseRules(lines []string) (map[int][]Rule, error) {
	rules := make(map[int][]Rule)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return nil, fmt.Errorf("failed to parse rule: line=%s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		rule := Rule{left, right}
		rules[left] = append(rules[left], rule)
		rules[right] = append(rules[right], rule)
	}
	return rules, nil
}

func parseUpdates(lines []string) ([][]int, error) {
	updates := make([][]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		update := make([]int, 0, len(parts))
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return updates, nil
}
