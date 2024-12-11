package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
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
	for i := 0; i < 25; i++ {
		working := make([]int, 0, len(curr))
		for _, num := range curr {
			if num == 0 {
				working = append(working, 1)
				continue
			}
			len := decimalLen(num)
			if len%2 == 0 {
				mid := decimalOfLen(1 + len/2)
				working = append(working, num/mid)
				working = append(working, num%mid)
			} else {
				working = append(working, num*2024)
			}
		}
		curr = working
	}

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
	fmt.Println("result: ", len(curr))
}

func decimalLen(num int) int {
	if num == 0 {
		return 0
	}
	if num < 0 {
		return decimalLen(-num)
	}
	for i := 3; i < 100; i++ {
		if num < decimalOfLen(i) {
			return i - 1
		}
	}
	panic("unexpected number")
}

func decimalOfLen(len int) int {
	num := 1
	for i := 1; i < len; i++ {
		num *= 10
	}
	return num
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
