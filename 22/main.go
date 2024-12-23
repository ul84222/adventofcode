package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bid struct {
	price int
	diff  int
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)
	secrets := parse(input)

	bids := make([][]Bid, len(secrets))
	for i := 0; i < 2000; i++ {
		for j := 0; j < len(secrets); j++ {
			secrets[j] = next(secrets[j])
			price := secrets[j] % 10
			if i == 0 {
				bids[j] = []Bid{{price, 0}}
			} else {
				prev := bids[j][i-1]
				bids[j] = append(bids[j], Bid{price, price - prev.price})
			}
		}
	}

	result := 0
	for _, s := range secrets {
		result += s
	}
	fmt.Println(result)
	fmt.Println(best(bids))
}

func best(bids [][]Bid) int {
	getKey := func(change []Bid) string {
		tmp := ""
		tmp += strconv.Itoa(change[0].diff) + ":"
		tmp += strconv.Itoa(change[1].diff) + ":"
		tmp += strconv.Itoa(change[2].diff) + ":"
		tmp += strconv.Itoa(change[3].diff) + ":"
		return tmp
	}

	max := -1
	cache := map[string]bool{}
	for i := 0; i < len(bids); i++ {
		for j := 1; j < len(bids[i])-3; j++ {
			change := bids[i][j : j+4]
			key := getKey(change)
			if cache[key] {
				continue
			}
			cache[key] = true
			tmp := count(bids[i:], change)
			if tmp > max {
				max = tmp
			}
		}
		fmt.Printf("I %v, max %v\n", i, max)
	}
	return max
}

func count(bids [][]Bid, change []Bid) int {
	total := 0
	for i := 0; i < len(bids); i++ {
		total += find(bids[i], change)
	}
	return total
}

func find(bids []Bid, change []Bid) int {
	fetch := func(i int) (int, bool) {
		if i+3 >= len(bids) {
			return 0, false
		}
		if bids[i].diff != change[0].diff ||
			bids[i+1].diff != change[1].diff ||
			bids[i+2].diff != change[2].diff ||
			bids[i+3].diff != change[3].diff {
			return 0, false
		}
		return bids[i+3].price, true
	}

	for i := 1; i < len(bids)-3; i++ {
		price, found := fetch(i)
		if found {
			return price
		}
	}
	return 0
}

func next(s int) int {
	tmp := s * 64
	s = mix(tmp, s)
	s = prune(s)

	tmp = s / 32
	s = mix(tmp, s)
	s = prune(s)

	tmp = s * 2048
	s = mix(tmp, s)
	return prune(s)
}

func mix(val, secret int) int {
	return val ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}

func parse(input string) []int {
	secrets := []int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		s, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		secrets = append(secrets, s)
	}
	return secrets
}
