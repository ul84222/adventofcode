package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	fmt.Println("input: ", input)
	runes := []rune(input)
	result := 0
	for i := 0; i < len(runes); {
		readResult, next := read(runes, i)
		result += readResult
		i = next
	}
	fmt.Println("Result: ", result)
}

func read(input []rune, curr int) (int, int) {
	if curr > len(input)-7 { // len("mul(1,2)") == 8
		return 0, curr + 1
	}
	if input[curr] != 'm' || input[curr+1] != 'u' || input[curr+2] != 'l' || input[curr+3] != '(' {
		return 0, curr + 1
	}
	left, pointer, err := readNumber(input, ',', curr+4)
	if err != nil {
		return 0, curr + 1
	}
	right, pointer, err := readNumber(input, ')', pointer+1)
	if err != nil {
		return 0, curr + 1
	}

	result := left * right
	fmt.Printf("'%s' -> %d\n", string(input[curr:pointer+1]), result)
	return result, pointer + 1
}

func readNumber(input []rune, delimiter rune, curr int) (int, int, error) {
	pointer := curr
	for ; pointer < len(input) && input[pointer] != delimiter; pointer++ {
	}
	if pointer >= len(input) {
		return 0, 0, fmt.Errorf("cant find number")
	}

	raw := string(input[curr:pointer])
	if number, err := strconv.Atoi(raw); err != nil {
		return 0, 0, fmt.Errorf("cant find number")
	} else if number >= 1000 {
		return 0, 0, fmt.Errorf("cant find number")
	} else {
		return number, pointer, nil
	}
}
