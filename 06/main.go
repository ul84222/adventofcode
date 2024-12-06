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
	historianMap := parse(input)

	visited := 0
	visitedMap := make(map[string]bool)

	currI, currJ := findCurrentPosition(historianMap)
	currDirection := "u"
	for currI >= 0 && currI < len(historianMap) && currJ >= 0 && currJ < len(historianMap[currI]) {
		roomKey := strconv.Itoa(currI) + "_" + strconv.Itoa(currJ)
		if !visitedMap[roomKey] {
			visited++
			historianMap[currI][currJ] = 'X'
			visitedMap[roomKey] = true
		}

		if currDirection == "u" {
			if currI == 0 || historianMap[currI-1][currJ] != '#' {
				currI--
			} else {
				currDirection = "r"
				continue
			}
		}

		if currDirection == "r" {
			if currJ >= len(historianMap[currI])-1 || historianMap[currI][currJ+1] != '#' {
				currJ++
			} else {
				currDirection = "d"
				continue
			}
		}

		if currDirection == "d" {
			if currI >= len(historianMap)-1 || historianMap[currI+1][currJ] != '#' {
				currI++
			} else {
				currDirection = "l"
				continue
			}
		}

		if currDirection == "l" {
			if currJ == 0 || historianMap[currI][currJ-1] != '#' {
				currJ--
			} else {
				currDirection = "u"
				continue
			}
		}
	}
	fmt.Println("Visited: ", visited)
}

func printMap(historianMap [][]rune) {
	for _, row := range historianMap {
		for _, el := range row {
			fmt.Printf("%v ", string(el))
		}
		fmt.Println()
	}
}

func findCurrentPosition(historianMap [][]rune) (int, int) {
	for i := 0; i < len(historianMap); i++ {
		for j := 0; j < len(historianMap[i]); j++ {
			if historianMap[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func parse(input string) [][]rune {
	lines := strings.Split(input, "\n")
	historianMap := make([][]rune, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		historianMap = append(historianMap, []rune(line))
	}
	return historianMap
}
