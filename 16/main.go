package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pair struct {
	i int
	j int
}

type Puzzle [][]rune

type Route struct {
	i    int
	j    int
	d    rune
	cost int
	path map[Pair]bool
}

type Routes struct {
	routes []Route
	ci     int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle := parsePuzzle(input)

	now := time.Now()
	puzzle.solve()
	elapsed := time.Since(now)
	fmt.Println("Elapsed: ", elapsed)
}

func (p Puzzle) solve() {
	i, j := startPoisition(p)
	m := make(map[Pair]bool)
	m[Pair{i, j}] = true
	route := Route{i, j, '>', 0, m}
	routes := &Routes{[]Route{}, 0}

	lowestCost := -1
	found := []Route{}
	for {
		route = next(p, route, routes)
		if p[route.i][route.j] == 'E' {
			if lowestCost == -1 {
				lowestCost = route.cost
			}
			if lowestCost == route.cost {
				found = append(found, route)
			} else {
				break
			}
		}
	}
	fmt.Println("Route: ", lowestCost)
	fmt.Println("Found: ", len(found))

	distinct := make(map[Pair]bool)
	for _, f := range found {
		for k, v := range f.path {
			distinct[k] = v
		}
	}
	fmt.Println("Result: ", len(distinct))
}

func next(p Puzzle, r Route, routes *Routes) Route {
	applicable := func(incI, incJ int) bool {
		i := r.i + incI
		j := r.j + incJ

		test := p[i][j] != '#' && !r.path[Pair{i, j}]
		return test
	}

	d := r.clockwise()
	incI, incJ := increments(d)
	if applicable(incI, incJ) {
		f := r.fork()
		f[Pair{r.i + incI, r.j + incJ}] = true
		route := Route{r.i + incI, r.j + incJ, d, r.cost + 1001, f}
		routes.add(route)
	}

	d = r.counterClockwise()
	incI, incJ = increments(d)
	if applicable(incI, incJ) {
		f := r.fork()
		f[Pair{r.i + incI, r.j + incJ}] = true
		route := Route{r.i + incI, r.j + incJ, d, r.cost + 1001, f}
		routes.add(route)
	}

	d = r.opposite()
	incI, incJ = increments(d)
	if applicable(incI, incJ) {
		f := r.fork()
		f[Pair{r.i + incI, r.j + incJ}] = true
		route := Route{r.i + incI, r.j + incJ, d, r.cost + 2002, f}
		routes.add(route)
	}

	incI, incJ = increments(r.d)
	if applicable(incI, incJ) {
		r.i += incI
		r.j += incJ
		r.cost++
		r.path[Pair{r.i, r.j}] = true
		if routes.hasCheaper(r) {
			routes.add(r)
			return routes.cheapest()
		} else {
			return r
		}
	} else {
		return routes.cheapest()
	}
}

func (p Puzzle) show() {
	fmt.Println()
	for _, line := range p {
		fmt.Println(string(line))
	}
}

func (r *Routes) add(route Route) {
	r.routes = append(r.routes, route)
	latest := len(r.routes) - 1
	if r.routes[r.ci].cost > r.routes[latest].cost {
		r.ci = latest
	}
}

func (r *Routes) cheapest() Route {
	res := r.routes[r.ci]
	r.routes = append(r.routes[:r.ci], r.routes[r.ci+1:]...)
	r.findCheapest()
	return res
}

func (r *Routes) hasCheaper(route Route) bool {
	return len(r.routes) > 0 && r.routes[r.ci].cost < route.cost
}

func (r *Routes) findCheapest() {
	ci := 0
	for i, val := range r.routes {
		if r.routes[ci].cost > val.cost {
			ci = i
		}
	}
	r.ci = ci
}

func (r *Route) fork() map[Pair]bool {
	fork := make(map[Pair]bool)
	for k, v := range r.path {
		fork[k] = v
	}
	return fork
}

func (r *Route) clockwise() rune {
	switch r.d {
	case '>':
		return 'v'
	case '<':
		return '^'
	case '^':
		return '>'
	case 'v':
		return '<'
	default:
		panic(fmt.Errorf("unknown direction: %v", r.d))
	}
}

func (r *Route) counterClockwise() rune {
	switch r.d {
	case '>':
		return '^'
	case '<':
		return 'v'
	case '^':
		return '<'
	case 'v':
		return '>'
	default:
		panic(fmt.Errorf("unknown direction: %v", r.d))
	}
}

func (r *Route) opposite() rune {
	switch r.d {
	case '>':
		return '<'
	case '<':
		return '>'
	case '^':
		return 'v'
	case 'v':
		return '^'
	default:
		panic(fmt.Errorf("unknown direction: %v", r.d))
	}
}
func increments(d rune) (int, int) {
	switch d {
	case '>':
		return 0, 1
	case '<':
		return 0, -1
	case '^':
		return -1, 0
	case 'v':
		return 1, 0
	default:
		panic(fmt.Errorf("unknown direction: %v", string(d)))
	}
}

func startPoisition(puzzle [][]rune) (int, int) {
	for i, line := range puzzle {
		for j, el := range line {
			if el == 'S' {
				return i, j
			}
		}
	}
	panic(fmt.Errorf("cant find start position"))
}

func parsePuzzle(input string) Puzzle {
	lines := strings.Split(input, "\n")
	puzzle := make([][]rune, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		puzzle = append(puzzle, []rune(line))
	}
	return puzzle
}
