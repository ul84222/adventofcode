package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Map = [][]rune

type Pair struct {
	i int
	j int
}

type Puzzle struct {
	room  Map
	moves []rune
	currI int
	currJ int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parsePuzzle(input)
	now := time.Now()

	puzzle.applyMoves()
	puzzle.printMap()
	elapsed := time.Since(now)
	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Coordinate: ", puzzle.coordinates())
}

func (p *Puzzle) applyMoves() {
	for _, m := range p.moves {
		// fmt.Printf("Move: >%v<\n", string(m))
		switch m {
		case '^':
			p.move(-1, 0)
		case 'v':
			p.move(1, 0)
		case '>':
			p.move(0, 1)
		case '<':
			p.move(0, -1)
		default:
			panic(fmt.Errorf("unknown move: >%v<", string(m)))
		}
	}
}

func (p *Puzzle) move(incI, incJ int) bool {
	i := p.currI + incI
	j := p.currJ + incJ
	isWallMoved := false

	for ; p.room[i][j] != '#' && p.room[i][j] != '.'; i, j = i+incI, j+incJ {
		if p.room[i][j] == 'O' {
			isWallMoved = true
		}
	}

	if p.room[i][j] == '#' {
		return false
	}

	if isWallMoved {
		p.room[i][j] = 'O'
	}
	p.room[p.currI][p.currJ] = '.'

	p.currI += incI
	p.currJ += incJ
	p.room[p.currI][p.currJ] = '@'

	return true
}

func (p *Puzzle) printMap() {
	for _, line := range p.room {
		fmt.Println(string(line))
	}
}

func (p *Puzzle) coordinates() int {
	result := 0
	for i := 0; i < len(p.room); i++ {
		for j := 0; j < len(p.room[i]); j++ {
			if p.room[i][j] == 'O' {
				result += 100*i + j
			}

		}
	}
	return result
}

func parsePuzzle(input string) Puzzle {
	lines := strings.Split(input, "\n")
	room := [][]rune{}
	i := 0
	for ; lines[i] != ""; i++ {
		room = append(room, []rune(lines[i]))
	}

	moves := make([]rune, 0, 0)
	i++
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		moves = append(moves, []rune(lines[i])...)
	}
	i, j := findCurrPosition(room)
	return Puzzle{room, moves, i, j}
}

func findCurrPosition(room Map) (int, int) {
	for i := 0; i < len(room); i++ {
		for j := 0; j < len(room[i]); j++ {
			if room[i][j] == '@' {
				return i, j
			}
		}
	}
	panic(fmt.Errorf("Cant find robot"))
}
