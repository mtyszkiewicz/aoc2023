package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day06"
)

func main() {
	parser := participle.MustBuild[day06.Scoreboard]()
	bytes, _ := io.ReadAll(os.Stdin)
	data := strings.ReplaceAll(string(bytes), " ", "")
	scoreboard, _ := parser.ParseString("", data)
	result := day06.CountWinOptions(scoreboard.Times[0], scoreboard.Distances[0])
	fmt.Println(result)
}
