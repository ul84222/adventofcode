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

type Task struct {
	p       Path
	attempt int
}

type Paths struct {
	paths []Path
	ci    int
}

const (
	// shouldSaveAtLeast = 74
	shouldSaveAtLeast = 100
	// shouldSaveAtLeast = 72
	cheatMaxTime = 20
	// shouldSaveAtLeast = 64
	// cheatMaxTime      = 1
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parse(input)

	p := puzzle.shortestPath()
	fmt.Println("shortest: ", p.steps)
	fmt.Println("cheat opportunities: ", puzzle.cheatOpportunities(p, p.cost-shouldSaveAtLeast))
}

// too low - 9428
// too high 1433001
func (puzzle Map) shortestPath() *Path {
	paths := Paths{[]Path{}, 0}
	mem := make(map[Pair]bool)

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
		if mem[Pair{i, j}] {
			return nil
		}
		if !safe(i, j) {
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

	for curr != nil && (curr.i != endI || curr.j != endJ) {
		curr.mem[Pair{curr.i, curr.j}] = curr.cost
		mem[Pair{curr.i, curr.j}] = true

		scheduleNext(curr, 1, 0)
		scheduleNext(curr, -1, 0)
		scheduleNext(curr, 0, 1)
		scheduleNext(curr, 0, -1)

		curr = paths.cheapest()
	}
	curr.mem[Pair{curr.i, curr.j}] = curr.cost
	return curr
}

func (puzzle Map) cheatOpportunities(p *Path, lte int) int {
	count := 0
	for it, cost := range p.mem {
		dest := make(map[Pair]int)
		seen := make(map[Pair]bool)
		seen[it] = true
		// puzzle.process(p, Pair{it.i + 1, it.j}, cost, lte, 0, dest, seen)
		// puzzle.process(p, Pair{it.i - 1, it.j}, cost, lte, 0, dest, seen)
		// puzzle.process(p, Pair{it.i, it.j + 1}, cost, lte, 0, dest, seen)
		// puzzle.process(p, Pair{it.i, it.j - 1}, cost, lte, 0, dest, seen)

		puzzle.process(p, Pair{it.i + 1, it.j}, cost, lte, 0, dest, it)
		puzzle.process(p, Pair{it.i - 1, it.j}, cost, lte, 0, dest, it)
		puzzle.process(p, Pair{it.i, it.j + 1}, cost, lte, 0, dest, it)
		puzzle.process(p, Pair{it.i, it.j - 1}, cost, lte, 0, dest, it)
		for _, c := range dest {
			if c <= lte {
				count++
			}

		}
		// count += len(dest)
	}
	return count
}

func (puzzle Map) process(p *Path, pair Pair, initialCost, lte, attempt int, dest map[Pair]int, seen map[Pair]bool) {
	// func (puzzle Map) process(p *Path, pair Pair, initialCost, lte, attempt int, dest map[Pair]int, prev Pair) {
	// if seen[pair] {
	// 	return
	// }
	// seen[pair] = true

	if attempt > cheatMaxTime {
		return
	}

	if pair.i < 0 || pair.j < 0 || pair.i >= len(puzzle) || pair.j >= len(puzzle[pair.i]) {
		return
	}

	endCost, found := p.mem[pair]
	if found {
		newCost := p.cost - (endCost - initialCost - 1 - attempt)
		if curr := dest[pair]; curr == 0 || curr > newCost {
			// fmt.Printf("FOUND: initialCost %v cost %v p.cost %v newCost %v\n", initialCost, endCost, p.cost, newCost)
			dest[pair] = newCost
		}
		// }
		return
	}

	// if prev.i != pair.i+1 {
	// 	puzzle.process(p, Pair{pair.i + 1, pair.j}, initialCost, lte, attempt+1, dest, pair)
	// }
	// if prev.i != pair.i-1 {
	// 	puzzle.process(p, Pair{pair.i - 1, pair.j}, initialCost, lte, attempt+1, dest, pair)
	// }
	// if prev.j != pair.j+1 {
	// 	puzzle.process(p, Pair{pair.i, pair.j + 1}, initialCost, lte, attempt+1, dest, pair)
	// }
	// if prev.j != pair.j-1 {
	// 	puzzle.process(p, Pair{pair.i, pair.j - 1}, initialCost, lte, attempt+1, dest, pair)
	// }

	// puzzle.process(p, Pair{pair.i + 1, pair.j}, initialCost, lte, attempt+1, dest, copyB(seen))
	// puzzle.process(p, Pair{pair.i - 1, pair.j}, initialCost, lte, attempt+1, dest, copyB(seen))
	// puzzle.process(p, Pair{pair.i, pair.j + 1}, initialCost, lte, attempt+1, dest, copyB(seen))
	// puzzle.process(p, Pair{pair.i, pair.j - 1}, initialCost, lte, attempt+1, dest, copyB(seen))
}

func copy(mem map[Pair]int) map[Pair]int {
	c := make(map[Pair]int, len(mem))
	for k, v := range mem {
		c[k] = v
	}
	return c
}

func copyB(mem map[Pair]bool) map[Pair]bool {
	c := make(map[Pair]bool, len(mem))
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
	panic("cant find " + string(r) + " in map")
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
	m := make(Map, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		m = append(m, []rune(line))
	}
	return m
}
