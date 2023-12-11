package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mtyszkiewicz/aoc2023/pkg/day03"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	schema := day03.LoadEngineSchematic(scanner)
	result := 0
	for point, symbol := range schema.Symbols {
		if symbol == '*' {
			parts := schema.NeighbouringParts(point)
			if len(parts) == 2 {
				result += parts[0] * parts[1] // gear ratio
			}
		}
	}
	fmt.Println(result)
}
