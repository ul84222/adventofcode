package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	puzzle, err := parsePuzzle(input)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	curr := puzzle

	result := 0
	for _, num := range curr {
		result += blink(num, 75)
	}

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
	fmt.Println("result: ", result)
}

type Pair struct {
	num       int
	remainint int
}

var (
	mem = make(map[Pair]int)
)

func blink(num int, remaining int) int {
	if remaining == 0 {
		return 1
	}
	if num == 0 {
		return blink(1, remaining-1)
	}
	len := decimalLen(num)
	if len%2 == 0 {
		pair := Pair{num, remaining}
		if mem[pair] != 0 {
			return mem[pair]
		}
		mid := decimalOfLen(1 + len/2)
		result := blink(num/mid, remaining-1) + blink(num%mid, remaining-1)
		mem[pair] = result
		return result
	} else {
		return blink(num*2024, remaining-1)
	}
}

func decimalLen(num int) int {
	if num == 0 {
		return 0
	}
	if num < 0 {
		return decimalLen(-num)
	}
	for i := 2; i < 100; i++ {
		if num < decimalOfLen(i) {
			return i - 1
		}
	}
	panic("unexpected number")
}

func decimalOfLen(len int) int {
	return int(math.Pow10(len - 1))
}

func parsePuzzle(input string) ([]int, error) {
	parts := strings.Fields(strings.TrimSpace(input))
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
