package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day04"
)

func main() {
	parser := participle.MustBuild[day04.ScratchCard]()
	scanner := bufio.NewScanner(os.Stdin)

	matches := []int{}
	for scanner.Scan() {
		card, err := parser.ParseString("", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		matches = append(matches, card.MatchCount())
	}

	nCards := len(matches)
	copies := make([]int, nCards)
	for i := 0; i < nCards; i++ {
		copies[i] = 0
	}
	result := 0
	for i, n := range matches {
		for j := i + 1; j < i+1+n && j < nCards; j++ {
			copies[j] += 1 + copies[i]
		}
		result += 1 + copies[i]
	}
	fmt.Println(result)
}
