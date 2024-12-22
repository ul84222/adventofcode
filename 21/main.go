package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	humanKeypad = [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}
	roboKeypad = [][]rune{
		{' ', '^', 'A'},
		{'<', 'v', '>'},
	}
	controls = map[Pair]rune{
		{1, 0}:  'v',
		{-1, 0}: '^',
		{0, 1}:  '>',
		{0, -1}: '<',
	}
	rControls = map[rune]Pair{
		'v': {1, 0},
		'^': {-1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}
	depth = 25
)

type Paths struct {
	paths []Path
	ci    int
}

type Path struct {
	i, j    int
	cost    int
	command []rune
	final   bool
}

type Pair struct {
	i, j int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	codes := parse(input)

	result := 0
	for _, code := range codes {
		seq := find(code, humanKeypad, 3, 2, depth)
		for i := 1; i <= depth; i++ {
			seq = find(seq, roboKeypad, 0, 2, depth-i)
			fmt.Println("seq: ", len(seq))
		}
		result = result + (numeric(code) * len(seq))
	}
	fmt.Println("result", result)
}

func reverse(str string) {
	i0, j0 := 3, 2
	i1, j1 := 0, 2
	i2, j2 := 0, 2
	l1, l2, l3 := 0, 0, 0
	buff := []rune{}
	for _, r := range []rune(str) {
		l3++
		if r == 'A' {
			l2++
			r = roboKeypad[i2][j2]
			if r == 'A' {
				l1++
				r = roboKeypad[i1][j1]
				buff = append(buff, r)
				if r == 'A' {
					fmt.Print(string(humanKeypad[i0][j0]))
				} else {
					inc := rControls[r]
					i0 += inc.i
					j0 += inc.j
				}
			} else {
				inc := rControls[humanKeypad[i2][j2]]
				i1 += inc.i
				j1 += inc.j
			}
		} else {
			inc := rControls[r]
			i2 += inc.i
			j2 += inc.j
		}
	}
	fmt.Printf(" : total: %v, l1: %v, l2: %v, l3: %v\n", l1+l2+l3, l1, l2, l3)
	fmt.Println("Buff: ", string(buff))
}

func numeric(code []rune) int {
	val, err := strconv.Atoi(string(code[:len(code)-1]))
	if err != nil {
		panic(err)
	}
	return val
}

type Key struct {
	code   string
	keypad int
	i, j   int
	depth  int
}

var findCache = make(map[Key][]rune)

func find(code []rune, keypad [][]rune, i, j int, depth int) []rune {
	// func find(code []rune, keypad [][]rune, i, j int, depth int) []rune {
	k := Key{
		string(code),
		len(keypad),
		i, j,
		depth,
	}
	if findCache[k] != nil {
		return findCache[k]
	}

	result := []rune{}
	part := []rune{}
	for _, c := range code {
		part, i, j = seq(c, i, j, keypad, depth, result)
		result = append(result, part...)
	}
	findCache[k] = result
	return result
}

func seq(code rune, i, j int, keypad [][]rune, depth int, prefix []rune) ([]rune, int, int) {
	paths := Paths{[]Path{}, 0}
	mem := make(map[Pair]bool)
	safe := func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(keypad) || j >= len(keypad[i]) {
			return false
		}
		if keypad[i][j] == ' ' {
			return false
		}
		return true
	}
	next := func(p *Path, incI, incJ int) *Path {
		i, j := p.i+incI, p.j+incJ
		if mem[Pair{i, j}] {
			return nil
		}
		if !safe(i, j) {
			return nil
		}

		c := make([]rune, len(p.command))
		copy(c, p.command)
		c = append(c, controls[Pair{incI, incJ}])

		final := keypad[i][j] == code
		if final {
			c = append(c, 'A')
		}

		test := append(prefix, c...)
		seq := test
		for k := 1; k <= depth; k++ {
			seq = find(seq, roboKeypad, 0, 2, depth-k)
		}
		cost := len(seq)
		n := Path{i, j, cost, c, final}
		return &n
	}
	scheduleNext := func(p *Path, incI, incJ int) {
		n := next(p, incI, incJ)
		if n != nil {
			paths.add(*n)
		}
	}

	if keypad[i][j] == code {
		return []rune{'A'}, i, j
	}

	curr := &Path{i, j, 0, []rune{}, false}
	for curr != nil && !curr.final {
		mem[Pair{curr.i, curr.j}] = true

		scheduleNext(curr, -1, 0)
		scheduleNext(curr, 1, 0)
		scheduleNext(curr, 0, -1)
		scheduleNext(curr, 0, 1)

		curr = paths.cheapest()
	}
	return curr.command, curr.i, curr.j
}

func countRune(a []rune, r rune) int {
	total := 0
	for _, it := range a {
		if it == r {
			total++
		}
	}
	return total
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

func parse(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		result = append(result, []rune(line))
	}

	return result
}
