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

	puzzle.wide()
	puzzle.applyMoves()
	puzzle.printMap()
	elapsed := time.Since(now)
	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Coordinate: ", puzzle.coordinates())
}

func (p *Puzzle) wide() {
	wideRoom := make([][]rune, 0, len(p.room))
	for _, l := range p.room {
		wideL := make([]rune, 0, len(p.room)*2)
		for _, r := range l {
			switch r {
			case '#':
				wideL = append(wideL, '#', '#')
			case 'O':
				wideL = append(wideL, '[', ']')
			case '.':
				wideL = append(wideL, '.', '.')
			case '@':
				wideL = append(wideL, '@', '.')
			default:
				panic(fmt.Errorf("unknown rune: %v", r))
			}
		}
		wideRoom = append(wideRoom, wideL)
	}
	p.room = wideRoom
	p.currI, p.currJ = findCurrPosition(p.room)
}

func (p *Puzzle) applyMoves() {
	p.printMap()
	for _, m := range p.moves {
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

	obstacles := 0
	for ; p.room[i][j] != '#' && p.room[i][j] != '.'; i, j = i+incI, j+incJ {
		obstacles++
	}
	if p.room[i][j] == '#' {
		return false
	}

	if obstacles == 0 {
		p.room[p.currI][p.currJ] = '.'
		p.currI += incI
		p.currJ += incJ
		p.room[p.currI][p.currJ] = '@'
		return true
	}

	if incJ != 0 {
		for ; j != p.currJ; j -= incJ {
			p.room[p.currI][j] = p.room[p.currI][j-incJ]
		}
		p.room[p.currI][p.currJ] = '.'
		p.currI += incI
		p.currJ += incJ
		p.room[p.currI][p.currJ] = '@'
		return true
	}
	if incI != 0 {
		mem := make(map[Pair]bool)
		canMove := verticalPush(p.room, p.currI+incI, p.currJ, incI, true, mem)
		if canMove {
			clear(mem)
			verticalPush(p.room, p.currI+incI, p.currJ, incI, false, mem)

			p.room[p.currI][p.currJ] = '.'
			p.currI += incI
			p.currJ += incJ
			p.room[p.currI][p.currJ] = '@'
			return true
		}
		return false
	}
	panic(fmt.Errorf("Unreachable state"))
}

func verticalPush(m [][]rune, i, j, incI int, dryRun bool, mem map[Pair]bool) bool {
	if m[i][j] == ']' {
		j--
	}

	if mem[Pair{i, j}] {
		return true
	} else {
		mem[Pair{i, j}] = true
	}

	nextI := i + incI
	if m[nextI][j] == '#' || m[nextI][j+1] == '#' {
		return false
	}
	if m[nextI][j] == '.' && m[nextI][j+1] == '.' {
		if dryRun {
			return true
		}
		m[nextI][j] = '['
		m[nextI][j+1] = ']'
		m[i][j] = '.'
		m[i][j+1] = '.'

		return true
	}
	leftMoved := m[nextI][j] == '.' || verticalPush(m, nextI, j, incI, dryRun, mem)
	rightMoved := m[nextI][j+1] == '.' || verticalPush(m, nextI, j+1, incI, dryRun, mem)
	if leftMoved && rightMoved {
		if dryRun {
			return true
		}
		m[nextI][j] = '['
		m[nextI][j+1] = ']'
		m[i][j] = '.'
		m[i][j+1] = '.'
		return true
	}
	return false
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
			if p.room[i][j] == '[' {
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
