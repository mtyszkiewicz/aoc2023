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
	for point, _ := range schema.Symbols {
		parts := schema.NeighbouringParts(point)
		for _, value := range parts {
			result += value
		}
	}
	fmt.Println(result)
}
