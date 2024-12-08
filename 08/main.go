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

	antinodes := make(map[Pair]bool)
	boundary := Pair{len(antennas), len(antennas[0])}
	for _, group := range groups {
		findAntinodes(group, boundary, antinodes)
	}
	fmt.Println("Result: ", len(antinodes))
}

func findAntinodes(antennas []Pair, boundary Pair, dest map[Pair]bool) {
	propose := func(i int, j int) {
		if i >= 0 && j >= 0 && i < boundary.I && j < boundary.J {
			dest[Pair{i, j}] = true
		}
	}
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			left := antennas[i]
			right := antennas[j]

			posI := left.I + left.I - right.I
			posJ := left.J + left.J - right.J
			propose(posI, posJ)

			posI = right.I + right.I - left.I
			posJ = right.J + right.J - left.J
			propose(posI, posJ)
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
