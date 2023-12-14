package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type History struct {
	Values []int `parser:"@Number*"`
}

func PredictNext(values []int) int {
	diff := make([]int, len(values)-1)
	nonzero := true
	for i := 1; i < len(values); i++ {
		diff[i-1] = values[i] - values[i-1]
		if diff[i-1] != 0 {
			nonzero = false
		}
	}
	if nonzero {
		return values[len(values)-1]
	}
	return values[len(values)-1] + PredictNext(diff)
}

func PredictPrev(values []int) int {
	diff := make([]int, len(values)-1)
	nonzero := true
	for i := 1; i < len(values); i++ {
		diff[i-1] = values[i] - values[i-1]
		if diff[i-1] != 0 {
			nonzero = false
		}
	}
	if nonzero {
		return values[0]
	}
	return values[0] - PredictPrev(diff)
}

func main() {
	var lexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Number", `[-+]?(\d*\.)?\d+`},
		{"whitespace", `[ \t]+`},
	})
	parser := participle.MustBuild[History](participle.Lexer(lexer))
	scanner := bufio.NewScanner(os.Stdin)

	resultP1 := 0
	resultP2 := 0
	for scanner.Scan() {
		history, err := parser.ParseString("", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		resultP1 += PredictNext(history.Values)
		resultP2 += PredictPrev(history.Values)
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", resultP1, resultP2)
}
