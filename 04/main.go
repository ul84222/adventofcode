package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	field := parse(input)
	xmasCount := 0
	shapedXmasCount := 0
	for i, _ := range field {
		for j, _ := range field[i] {
			if isHorizontalHit(field, i, j) {
				xmasCount++
			}
			if isVerticalHit(field, i, j) {
				xmasCount++
			}
			if isDownDiagonalHit(field, i, j) {
				xmasCount++
			}
			if isUpDiagonalHit(field, i, j) {
				xmasCount++
			}
			if isShapedXmas(field, i, j) {
				shapedXmasCount++
			}
		}
	}
	fmt.Println("Found xmas: ", xmasCount)
	fmt.Println("Found shaped xmas: ", shapedXmasCount)
}

func isHorizontalHit(field [][]rune, i int, j int) bool {
	if j+3 >= len(field[i]) {
		return false
	}
	if field[i][j] == 'X' && field[i][j+1] == 'M' && field[i][j+2] == 'A' && field[i][j+3] == 'S' {
		return true
	}
	if field[i][j] == 'S' && field[i][j+1] == 'A' && field[i][j+2] == 'M' && field[i][j+3] == 'X' {
		return true
	}
	return false
}

func isVerticalHit(field [][]rune, i int, j int) bool {
	if i+3 >= len(field) {
		return false
	}
	if field[i][j] == 'X' && field[i+1][j] == 'M' && field[i+2][j] == 'A' && field[i+3][j] == 'S' {
		return true
	}
	if field[i][j] == 'S' && field[i+1][j] == 'A' && field[i+2][j] == 'M' && field[i+3][j] == 'X' {
		return true
	}
	return false
}

func isDownDiagonalHit(field [][]rune, i int, j int) bool {
	if i+3 >= len(field) || j+3 >= len(field[i]) {
		return false
	}
	if field[i][j] == 'X' && field[i+1][j+1] == 'M' && field[i+2][j+2] == 'A' && field[i+3][j+3] == 'S' {
		return true
	}
	if field[i][j] == 'S' && field[i+1][j+1] == 'A' && field[i+2][j+2] == 'M' && field[i+3][j+3] == 'X' {
		return true
	}
	return false
}

func isUpDiagonalHit(field [][]rune, i int, j int) bool {
	if i-3 < 0 || j+3 >= len(field[i]) {
		return false
	}
	if field[i][j] == 'X' && field[i-1][j+1] == 'M' && field[i-2][j+2] == 'A' && field[i-3][j+3] == 'S' {
		return true
	}
	if field[i][j] == 'S' && field[i-1][j+1] == 'A' && field[i-2][j+2] == 'M' && field[i-3][j+3] == 'X' {
		return true
	}
	return false
}

func isShapedXmas(field [][]rune, i int, j int) bool {
	if i+2 >= len(field) || j+2 >= len(field[i]) {
		return false
	}

	leftMatch := false
	if field[i][j] == 'M' && field[i+1][j+1] == 'A' && field[i+2][j+2] == 'S' {
		leftMatch = true
	}
	if field[i][j] == 'S' && field[i+1][j+1] == 'A' && field[i+2][j+2] == 'M' {
		leftMatch = true
	}

	rightMatch := false
	if field[i][j+2] == 'M' && field[i+1][j+1] == 'A' && field[i+2][j] == 'S' {
		rightMatch = true
	}
	if field[i][j+2] == 'S' && field[i+1][j+1] == 'A' && field[i+2][j] == 'M' {
		rightMatch = true
	}
	return leftMatch && rightMatch
}

func parse(input string) [][]rune {
	var result [][]rune

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		result = append(result, []rune(line))
	}
	return result
}
