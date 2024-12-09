package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Start  int
	Len    int
	FileId int
}

const (
	EmptyBlockFileId = -1
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	blocks, err := parse(input)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	blocks = fragmentV2(blocks)
	spent := time.Since(start)
	fmt.Println(checksum(blocks))
	fmt.Printf("spent: %s\n", spent)
}

func print(blocks []Block) {
	for _, block := range blocks {
		for i := 0; i < block.Len; i++ {
			if block.FileId == EmptyBlockFileId {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", block.FileId)
			}
		}
	}
	fmt.Println()
}

func fragment(blocks []Block) []Block {
	i := 0
	j := len(blocks) - 1

	for i < j {
		if blocks[i].FileId != EmptyBlockFileId {
			i++
			continue
		}
		if blocks[j].FileId == EmptyBlockFileId {
			j--
			continue
		}

		if blocks[i].Len == blocks[j].Len {
			blocks[i].FileId = blocks[j].FileId
			blocks[j].FileId = EmptyBlockFileId
			i++
			j--
		} else if blocks[i].Len > blocks[j].Len {
			newBlock := Block{
				Start:  blocks[i].Start + blocks[j].Len,
				Len:    blocks[i].Len - blocks[j].Len,
				FileId: EmptyBlockFileId,
			}
			blocks[i].FileId = blocks[j].FileId
			blocks[i].Len = blocks[j].Len
			blocks[j].FileId = EmptyBlockFileId
			blocks = slices.Insert(blocks, i+1, newBlock)
			i++
		} else {
			newBlock := Block{
				Start:  blocks[j].Start + blocks[i].Len,
				Len:    blocks[i].Len,
				FileId: EmptyBlockFileId,
			}
			blocks[i].FileId = blocks[j].FileId
			blocks[j].Len = blocks[j].Len - blocks[i].Len
			blocks = slices.Insert(blocks, j+1, newBlock)
			i++
		}
	}
	return blocks
}

func fragmentV2(blocks []Block) []Block {
	i := 0
	j := len(blocks) - 1

	for j >= 0 {
		if i >= j {
			i = 0
			j--
		}
		if blocks[i].FileId != EmptyBlockFileId {
			i++
			continue
		}
		if blocks[j].FileId == EmptyBlockFileId {
			j--
			continue
		}

		if blocks[i].Len == blocks[j].Len {
			blocks[i].FileId = blocks[j].FileId
			blocks[j].FileId = EmptyBlockFileId
			i = 0
			j--
		} else if blocks[i].Len > blocks[j].Len {
			newBlock := Block{
				Start:  blocks[i].Start + blocks[j].Len,
				Len:    blocks[i].Len - blocks[j].Len,
				FileId: EmptyBlockFileId,
			}
			blocks[i].FileId = blocks[j].FileId
			blocks[i].Len = blocks[j].Len
			blocks[j].FileId = EmptyBlockFileId
			blocks = slices.Insert(blocks, i+1, newBlock)
			i = 0
		} else {
			i++
		}
	}
	return blocks
}
func checksum(blocks []Block) int {
	checksum := 0
	for _, block := range blocks {
		if block.FileId == EmptyBlockFileId {
			continue
		}
		for i := block.Start; i < block.Start+block.Len; i++ {
			checksum += i * block.FileId
		}
	}
	return checksum
}

func parse(input string) ([]Block, error) {
	blocks := make([]Block, 0, len(input))
	pointer := 0
	for i, el := range strings.TrimSpace(input) {
		number, err := strconv.Atoi(string(el))
		if err != nil {
			return nil, err
		}
		if i%2 == 0 {
			blocks = append(blocks, Block{pointer, number, i / 2})
		} else {
			blocks = append(blocks, Block{pointer, number, EmptyBlockFileId})
		}
		pointer += number
	}
	return blocks, nil
}
