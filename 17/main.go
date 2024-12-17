package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	computer := parse(input)
	fmt.Println("Computer: ", computer)
	computer.execute()
}

func (c *Computer) execute() {
	opcode := 0
	operand := 0
	for c.pointer < len(c.program) {
		opcode = c.program[c.pointer]
		operand = c.program[c.pointer+1]
		c.pointer += 2
		instruction(opcode)(c, operand)
	}
}

type Computer struct {
	a       int
	b       int
	c       int
	program []int
	pointer int
}

type Instruction func(c *Computer, operand int)

func instruction(opcode int) Instruction {
	switch opcode {
	case 0:
		return adv
	case 1:
		return bxl
	case 2:
		return bst
	case 3:
		return jnz
	case 4:
		return bxc
	case 5:
		return out
	case 6:
		return bdv
	case 7:
		return cdv
	default:
		panic(fmt.Errorf("got unexpected opcode: %v", opcode))
	}
}

func (c *Computer) comboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic(fmt.Errorf("unexpected operand: %v", operand))
	}
}

func adv(c *Computer, operand int) {
	combo := c.comboOperand(operand)
	c.a = int(math.Trunc(float64(c.a) / math.Pow(2, float64(combo))))
}

func bxl(c *Computer, operand int) {
	c.b = c.b ^ operand
}

func bst(c *Computer, operand int) {
	c.b = c.comboOperand(operand) % 8
}

func jnz(c *Computer, operand int) {
	if c.a == 0 {
		return
	}
	c.pointer = operand
}

func bxc(c *Computer, operand int) {
	c.b = c.b ^ c.c
}

func out(c *Computer, operand int) {
	val := c.comboOperand(operand) % 8
	fmt.Printf("%v,", val)
}

func bdv(c *Computer, operand int) {
	combo := c.comboOperand(operand)
	c.b = int(float64(c.a) / math.Pow(2, float64(combo)))
}

func cdv(c *Computer, operand int) {
	combo := c.comboOperand(operand)
	c.c = int(math.Trunc(float64(c.a) / math.Pow(2, float64(combo))))
}

func parse(input string) Computer {
	atoi := func(str string) int {
		if num, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			return num
		}
	}
	readRegister := func(register, str string) int {
		pattern := regexp.MustCompile(`Register ` + register + `: (-?\d+)`)
		parts := pattern.FindStringSubmatch(str)
		return atoi(parts[1])
	}
	readProgram := func(str string) []int {
		pattern := regexp.MustCompile(`Program: (.*)`)
		parts := pattern.FindStringSubmatch(str)
		program := []int{}
		for _, it := range strings.Split(parts[1], ",") {
			program = append(program, atoi(it))
		}
		return program
	}

	lines := strings.Split(input, "\n")
	return Computer{
		readRegister("A", lines[0]),
		readRegister("B", lines[1]),
		readRegister("C", lines[2]),
		readProgram(lines[4]),
		0,
	}
}
