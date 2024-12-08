package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Pair struct {
	I int
	J int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	antennas := parseAntennas(input)

	groups := make(map[rune][]Pair)
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			antenna := antennas[i][j]
			if unicode.IsDigit(antenna) || unicode.IsLetter(antenna) {
				groups[antenna] = append(groups[antenna], Pair{i, j})
			}
		}
	}

	antinodes := make(map[Pair]int)
	boundary := Pair{len(antennas), len(antennas[0])}
	for _, group := range groups {
		findAntinodes(group, boundary, antinodes)
	}
	fmt.Println("Result: ", len(antinodes))
}

func findAntinodes(antennas []Pair, boundary Pair, dest map[Pair]int) {
	propose := func(i int, j int) bool {
		if i >= 0 && j >= 0 && i < boundary.I && j < boundary.J {
			dest[Pair{i, j}]++
			return true
		} else {
			return false
		}
	}
	apply := func(i int, j int, stepI int, stepJ int) {
		posI := i
		posJ := j
		for ; propose(posI, posJ); posI, posJ = posI+stepI, posJ+stepJ {
		}
	}

	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			left := antennas[i]
			right := antennas[j]

			stepI := left.I - right.I
			stepJ := left.J - right.J
			apply(left.I, left.J, stepI, stepJ)
			apply(right.I, right.J, -stepI, -stepJ)
		}
	}
}

func parseAntennas(input string) [][]rune {
	lines := strings.Split(input, "\n")
	antennas := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		antennas = append(antennas, []rune(line))
	}
	return antennas
}
