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
	currI, currJ := findCurrentPosition(historianMap)
	currDirection := "u"

	visited, possibleObstructionCount, loop := traverse(historianMap, currI, currJ, currDirection, false)

	fmt.Println("Visited: ", visited)
	fmt.Println("Possible obstructions: ", possibleObstructionCount)
	fmt.Println("Loop: ", loop)
}

func traverse(historianMap [][]rune, currI int, currJ int, currDirection string, test bool) (int, int, bool) {
	visited := 0
	visitedMap := make(map[string]string)
	obstructionMap := make(map[string]bool)
	possibleObstructions := 0

	for currI >= 0 && currI < len(historianMap) && currJ >= 0 && currJ < len(historianMap[currI]) {
		roomKey := createKey(currI, currJ)
		if visitedMap[roomKey] == "" {
			visited++
		} else if strings.Contains(visitedMap[roomKey], currDirection) {
			return visited, possibleObstructions, true
		}
		if !test {
			historianMap[currI][currJ] = 'X'
		}
		visitedMap[roomKey] += currDirection

		switch currDirection {
		case "u":
			if currI == 0 || historianMap[currI-1][currJ] != '#' {
				if !test && currI != 0 && historianMap[currI-1][currJ] != 'X' {
					historianMap[currI-1][currJ] = '#'
					if _, _, loop := traverse(historianMap, currI, currJ, currDirection, true); loop {
						possibleObstructions++
						obstructionMap[createKey(currI-1, currJ)] = true
					}
					historianMap[currI-1][currJ] = 'X'
				}
				currI--
			} else {
				currDirection = "r"
				continue
			}
		case "r":
			if currJ >= len(historianMap[currI])-1 || historianMap[currI][currJ+1] != '#' {
				if !test && currJ < len(historianMap[currI])-1 && historianMap[currI][currJ+1] != 'X' {
					historianMap[currI][currJ+1] = '#'
					if _, _, loop := traverse(historianMap, currI, currJ, currDirection, true); loop {
						possibleObstructions++
						obstructionMap[createKey(currI, currJ+1)] = true
					}
					historianMap[currI][currJ+1] = 'X'
				}
				currJ++
			} else {
				currDirection = "d"
				continue
			}
		case "d":
			if currI >= len(historianMap)-1 || historianMap[currI+1][currJ] != '#' {
				if !test && currI < len(historianMap)-1 && historianMap[currI+1][currJ] != 'X' {
					historianMap[currI+1][currJ] = '#'
					if _, _, loop := traverse(historianMap, currI, currJ, currDirection, true); loop {
						possibleObstructions++
						obstructionMap[createKey(currI+1, currJ)] = true
					}
					historianMap[currI+1][currJ] = 'X'
				}
				currI++
			} else {
				currDirection = "l"
				continue
			}
		case "l":
			if currJ == 0 || historianMap[currI][currJ-1] != '#' {
				if !test && currJ > 0 && historianMap[currI][currJ-1] != 'X' {
					historianMap[currI][currJ-1] = '#'
					if _, _, loop := traverse(historianMap, currI, currJ, currDirection, true); loop {
						possibleObstructions++
						obstructionMap[createKey(currI, currJ-1)] = true
					}
					historianMap[currI][currJ-1] = 'X'
				}
				currJ--
			} else {
				currDirection = "u"
				continue
			}
		}
	}
	return visited, len(obstructionMap), false
}

func createKey(i int, j int) string {
	return strconv.Itoa(i) + "_" + strconv.Itoa(j)
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
