package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/mtyszkiewicz/aoc2023/pkg/day05"
	"github.com/schollz/progressbar/v3"
)

func main() {
	parser := participle.MustBuild[day05.Almanac]()
	almanac, err := parser.Parse("", os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	lowestLocation := math.MaxInt
	for i := 0; i < len(almanac.Seeds); i += 2 {
		rangeStart := almanac.Seeds[i]
		rangeEnd := rangeStart + almanac.Seeds[i+1]

		fmt.Printf("Processing seed range (%d/%d):\n", i/2+1, len(almanac.Seeds)/2)
		bar := progressbar.Default(int64(rangeEnd-rangeStart), "Seeds searched")
		for seedValue := rangeStart; seedValue < rangeEnd; seedValue++ {
			value := seedValue
			for _, converter := range almanac.Converters {
				value = converter.Convert(value)
			}
			lowestLocation = min(lowestLocation, value)
			if seedValue%100000 == 0 {
				bar.Add(100000)
			}
		}
		bar.Finish()
	}
	fmt.Println(lowestLocation)
}
