package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	maxX  = 71
	maxY  = 71
	frame = 1024
)

type Pair struct {
	x, y int
}

type Path struct {
	x, y  int
	d     rune
	cost  int
	steps int
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
	corruptions := parse(input)

	for i := len(corruptions) - 1; i > 0; i-- {
		p := shortestPath(corruptions, i)
		if p != nil {
			fmt.Println("i:", i)
			break
		}
	}
	fmt.Println()
	// fmt.Println("steps: ", p.steps)
}

type Key struct {
	x, y int
}

func shortestPath(corruptions map[Pair]int, frame int) *Path {
	mem := make(map[Key]bool)
	safe := func(x, y int) bool {
		cr := corruptions[Pair{x, y}]
		if cr != 0 && cr <= frame {
			return false
		}
		return x >= 0 && y >= 0 && x < maxX && y < maxY
	}
	next := func(p Path, d rune) *Path {
		ix, iy := increments(d)
		x, y := p.x+ix, p.y+iy
		if mem[Key{x, y}] {
			return nil
		}
		if !safe(x, y) {
			return nil
		}
		p = Path{x, y, d, p.cost + cost(d), p.steps + 1}
		return &p
	}

	curr := &Path{0, 0, '>', 0, 0}
	paths := Paths{[]Path{}, 0}
	for curr != nil && (curr.x != (maxX-1) || curr.y != (maxY-1)) {
		mem[Key{curr.x, curr.y}] = true

		n := next(*curr, left(curr.d))
		if n != nil {
			paths.add(*n)
		}

		n = next(*curr, right(curr.d))
		if n != nil {
			paths.add(*n)
		}

		n = next(*curr, curr.d)
		if n == nil {
			curr = paths.cheapest()
		} else if paths.hasCheaper(*n) {
			paths.add(*n)
			curr = paths.cheapest()
		} else {
			curr = n
		}
	}
	return curr
}

func increments(d rune) (int, int) {
	switch d {
	case '>':
		return 1, 0
	case '<':
		return -1, 0
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	default:
		panic("unexpected d: " + string(d))
	}
}

func cost(d rune) int {
	switch d {
	case '>', 'v':
		return 0
	case '<', '^':
		return 1
	default:
		panic("unexpected d: " + string(d))
	}
}

func left(d rune) rune {
	switch d {
	case '>':
		return '^'
	case '<':
		return 'v'
	case '^':
		return '<'
	case 'v':
		return '>'
	default:
		panic("unexpected d: " + string(d))
	}
}

func right(d rune) rune {
	switch d {
	case '>':
		return 'v'
	case '<':
		return '^'
	case '^':
		return '>'
	case 'v':
		return '<'
	default:
		panic("unexpected d: " + string(d))
	}
}

func show(corrupted map[Pair]int, until int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			ct := corrupted[Pair{x, y}]
			if ct != 0 && ct <= until {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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

func parse(input string) map[Pair]int {
	atoi := func(str string) int {
		if it, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			return it
		}
	}

	lines := strings.Split(input, "\n")
	positions := make(map[Pair]int)
	for i, it := range lines {
		if it == "" {
			continue
		}
		parts := strings.Split(it, ",")
		positions[Pair{atoi(parts[0]), atoi(parts[1])}] = i + 1
	}
	return positions
}
