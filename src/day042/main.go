package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day04"
)

func main() {
	parser := participle.MustBuild[day04.ScratchCard]()
	scanner := bufio.NewScanner(os.Stdin)

	matchCounts := []int{}
	for scanner.Scan() {
		card, err := parser.ParseString("", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		matchCounts = append(matchCounts, card.MatchCount())
	}

	result := 0
	copyCounts := make([]int, len(matchCounts))
	for i := 0; i < len(matchCounts); i++ {
		copyCounts[i] = 0
	}
	for i, count := range matchCounts {
		for j := i + 1; j < int(math.Min(float64(i+1+count), float64(len(matchCounts)))); j++ {
			copyCounts[j] += 1 + copyCounts[i]
		}
		result += 1 + copyCounts[i]
	}
	fmt.Println(result)
}
