package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

type Node struct {
	Left  string `parser:"'(' @Ident ','"`
	Right string `parser:"@Ident ')'"`
}

type Location struct {
	Current    string `parser:"@Ident"`
	Directions *Node  `parser:"'=' @@"`
}

type Pouch struct {
	Instructions string     `parser:"@Ident"`
	Locations    []Location `parser:"@@*"`
}

type DesertMap struct {
	Instructions   string
	Network        map[string]Node
	StartLocations []string
}

func (p *Pouch) GetDesertMap() *DesertMap {
	dm := new(DesertMap)
	dm.Instructions = p.Instructions
	dm.Network = make(map[string]Node)
	for _, loc := range p.Locations {
		dm.Network[loc.Current] = *loc.Directions
		if loc.Current[2] == 'A' {
			dm.StartLocations = append(dm.StartLocations, loc.Current)
		}
	}
	return dm
}

func (dm *DesertMap) FindCyclePeriod(current string) int {
	stepCount := 0
	for {
		direction := dm.Instructions[stepCount%len(dm.Instructions)]

		switch direction {
		case 'L':
			current = dm.Network[current].Left
		case 'R':
			current = dm.Network[current].Right
		}
		stepCount++
		if current[2] == 'Z' {
			return stepCount
		}
	}
}

func gcd(a, b int) int {
	// Euclidean algorithm
	for b != 0 {
		var temp = b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a, b int) int {
	return (a * b / gcd(a, b))
}

func lcmm(args []int) int {
	// Recursively iterate through pairs of arguments
	// i.e. lcm(args[0], lcm(args[1], lcm(args[2], args[3])))

	if len(args) == 2 {
		return lcm(args[0], args[1])
	} else {
		var arg0 = args[0]
		args = args[1:]
		return lcm(arg0, lcmm(args))
	}
}

func main() {
	parser := participle.MustBuild[Pouch]()
	pouch, err := parser.Parse("", os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	desertMap := pouch.GetDesertMap()

	periods := []int{}
	for _, current := range desertMap.StartLocations {
		period := desertMap.FindCyclePeriod(current)
		periods = append(periods, period)
		if current == "AAA" {
			fmt.Printf("Part 1: %d\n", period)
		}
	}
	fmt.Printf("Part 2: %d\n", lcmm(periods))
}
