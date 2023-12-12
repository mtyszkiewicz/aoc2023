package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day05"
)

func main() {
	parser := participle.MustBuild[day05.Almanac]()
	almanac, err := parser.Parse("", os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	locations := []int{}
	for _, value := range almanac.Seeds {
		for _, converter := range almanac.Converters {
			value = converter.Convert(value)
		}
		locations = append(locations, value)
	}
	fmt.Println(slices.Min(locations))
}
