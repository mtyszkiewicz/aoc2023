package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day06"
)

func main() {
	parser := participle.MustBuild[day06.Scoreboard]()
	scoreboard, _ := parser.Parse("", os.Stdin)
	result := 1
	for i := 0; i < len(scoreboard.Times); i++ {
		result *= day06.CountWinOptions(scoreboard.Times[i], scoreboard.Distances[i])
	}
	fmt.Println(result)
}
