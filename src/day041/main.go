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

func Score(matchCount int) int {
	if matchCount == 0 {
		return 0
	} else {
		return int(math.Pow(float64(2), float64(matchCount-1)))
	}
}

func main() {
	parser := participle.MustBuild[day04.ScratchCard]()
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		card, err := parser.ParseString("", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		result += Score(card.MatchCount())
	}
	fmt.Println(result)
}
