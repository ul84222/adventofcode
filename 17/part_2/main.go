package main

import (
	"fmt"
	"math"
)

func main() {
	target := []int{2, 4, 1, 3, 7, 5, 4, 7, 0, 3, 1, 5, 5, 5, 3, 0}
	maxL, maxBase := 0, 0
	base := int(math.Pow(2, 3*15))

	inc := 15
	for i := 0; i < 100000000; i++ {
		output := execute(base)
		count := compare(target, output)
		if count == len(output) {
			fmt.Println("FOUND:", base)
			fmt.Println("FOUND output: ", output)
			fmt.Println("FOUND target: ", target)
			fmt.Println("ITERATION: ", i)
			break
		} else if count > maxL {
			maxL = count
			maxBase = base
			inc = 15 - count
			fmt.Printf("len: %v a: %v inc: %v\n", maxL, maxBase, inc)
		}

		base = base + int(math.Pow(2, 3*float64(inc)))
	}
}

func compare(base, test []int) int {
	if len(base) != len(test) {
		panic(fmt.Errorf("base %v test %v", len(base), len(test)))
	}
	len := len(base)
	for i := 0; i < len; i++ {
		if test[len-i-1] != base[len-i-1] {
			return i
		}
	}
	return len
}

func execute(a int) []int {
	i, val := 0, 0
	output := []int{}
	for a != 0 {
		val, a = program(a)
		output = append(output, val)
		i++
	}
	return output
}

func program(a int) (int, int) {
	b := a % 8
	b = b ^ 3
	c := int(float64(a) / math.Pow(2, float64(b)))
	b = b ^ c
	b = b ^ 5
	return b % 8, a / 8
}
