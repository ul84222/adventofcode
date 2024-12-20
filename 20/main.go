package main

import (
	"fmt"
	"os"
	"strings"
)

type Map [][]rune

type Pair struct {
	i, j int
}

type Path struct {
	i, j    int
	cost    int
	steps   int
	cheated bool
	mem     map[Pair]int
}

type Paths struct {
	paths []Path
	ci    int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parse(input)

	p := puzzle.shortestPath(false, -1)[0]
	fmt.Println("shortest: ", p.steps)

	r := puzzle.shortestPath(true, p.cost-12)
	fmt.Println("cheat shortest: ", len(r))
}

func (puzzle Map) shortestPath(cheat bool, lte int) []Path {
	paths := Paths{[]Path{}, 0}

	safe := func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) {
			return false
		}
		if puzzle[i][j] != '#' {
			return true
		}
		return false
	}
	next := func(p *Path, incI, incJ int) *Path {
		i, j := p.i+incI, p.j+incJ
		if p.mem[Pair{i, j}] != 0 {
			return nil
		}
		if !safe(i, j) {
			if cheat && !p.cheated && safe(i+incI, j+incJ) {
				n := Path{i + incI, j + incJ, p.cost + 2, p.steps + 2, true, copy(p.mem)}
				return &n
			}
			return nil
		}
		n := Path{i, j, p.cost + 1, p.steps + 1, p.cheated, copy(p.mem)}
		return &n
	}
	scheduleNext := func(p *Path, incI, incJ int) {
		n := next(p, incI, incJ)
		if n != nil {
			paths.add(*n)
		}
	}

	startI, startJ := puzzle.start()
	endI, endJ := puzzle.end()
	curr := &Path{startI, startJ, 0, 0, false, make(map[Pair]int)}

	t := make([]Path, 0, 1)
	for curr != nil {
		if curr.i == endI && curr.j == endJ {
			if lte == -1 {
				t = append(t, *curr)
				return t
			}
			if lte >= curr.cost {
				t = append(t, *curr)
				curr = paths.cheapest()
				continue
			} else {
				return t
			}
		}

		curr.mem[Pair{curr.i, curr.j}] = curr.cost

		scheduleNext(curr, 1, 0)
		scheduleNext(curr, -1, 0)
		scheduleNext(curr, 0, 1)
		scheduleNext(curr, 0, -1)

		curr = paths.cheapest()
	}
	return t
}

func copy(mem map[Pair]int) map[Pair]int {
	c := make(map[Pair]int, len(mem))
	for k, v := range mem {
		c[k] = v
	}
	return c
}

func (m *Map) start() (int, int) {
	return m.position('S')
}

func (m Map) end() (int, int) {
	return m.position('E')
}

func (m Map) position(r rune) (int, int) {
	for i, line := range m {
		for j, it := range line {
			if r == it {
				return i, j
			}
		}
	}
	panic("cant find " + string(r) + "in map")
}

func (r *Paths) add(p Path) {
	r.paths = append(r.paths, p)
	if r.paths[r.ci].cost > p.cost {
		r.ci = len(r.paths) - 1
	}
}

func (r *Paths) cheapest() *Path {
	if len(r.paths) == 0 {
		return nil
	}
	res := r.paths[r.ci]
	r.paths = append(r.paths[:r.ci], r.paths[r.ci+1:]...)
	r.findCheapest()
	return &res
}

func (r *Paths) findCheapest() {
	ci := 0
	for i, it := range r.paths {
		if r.paths[ci].cost > it.cost {
			ci = i
		}
	}
	r.ci = ci
}

func (r *Paths) hasCheaper(p Path) bool {
	return len(r.paths) > 0 && r.paths[r.ci].cost < p.cost
}

func parse(input string) Map {
	lines := strings.Split(input, "\n")
	m := make(Map, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		m = append(m, []rune(line))
	}
	return m
}
