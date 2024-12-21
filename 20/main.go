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
	p       Pair
	attempt int
	seen    map[Pair]bool
}

type Paths struct {
	paths []Path
	ci    int
}

const (
	shouldSaveAtLeast = 100
	maxCheatAttempts  = 20
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

var tasks = []Task{}

func (puzzle Map) cheatOpportunities(p *Path, lte int) int {
	count := 0
	for it, cost := range p.mem {
		seen := make(map[Pair]bool)
		dest := make(map[Pair]int)
		seen[it] = true

		tasks = append(tasks, Task{Pair{it.i + 1, it.j}, 1, seen})
		tasks = append(tasks, Task{Pair{it.i - 1, it.j}, 1, seen})
		tasks = append(tasks, Task{Pair{it.i, it.j + 1}, 1, seen})
		tasks = append(tasks, Task{Pair{it.i, it.j - 1}, 1, seen})

		for len(tasks) != 0 {
			task := tasks[0]
			tasks = tasks[1:]

			puzzle.process(p, task, cost, dest)
		}

		for _, c := range dest {
			if c <= lte {
				count++
			}
		}
	}

	return count
}

func (puzzle Map) process(p *Path, task Task, initialCost int, dest map[Pair]int) {
	pair := task.p

	if task.attempt > maxCheatAttempts {
		return
	}

	if task.seen[pair] {
		return
	}
	task.seen[pair] = true

	if pair.i < 0 || pair.j < 0 || pair.i >= len(puzzle) || pair.j >= len(puzzle[pair.i]) {
		return
	}

	endCost, found := p.mem[pair]
	if found {
		newCost := initialCost + (p.cost - endCost) + task.attempt
		if curr, found := dest[pair]; !found || curr > newCost {
			dest[pair] = newCost
		}
	}

	tasks = append(tasks, Task{Pair{pair.i + 1, pair.j}, task.attempt + 1, task.seen})
	tasks = append(tasks, Task{Pair{pair.i - 1, pair.j}, task.attempt + 1, task.seen})
	tasks = append(tasks, Task{Pair{pair.i, pair.j + 1}, task.attempt + 1, task.seen})
	tasks = append(tasks, Task{Pair{pair.i, pair.j - 1}, task.attempt + 1, task.seen})
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
