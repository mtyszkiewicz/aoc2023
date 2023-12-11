package day03

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	point := new(Point)
	point.x = x
	point.y = y
	return point
}

type EngineSchematic struct {
	grid    []string
	Parts   map[Point]int
	Symbols map[Point]rune
}

func (schema *EngineSchematic) NeighbouringParts(p Point) []int {
	result := mapset.NewSet[int]()
	for _, dx := range []int{1, 0, -1} {
		for _, dy := range []int{1, 0, -1} {
			partNumber, ok := schema.Parts[*NewPoint(p.x+dx, p.y+dy)]
			if ok {
				result.Add(partNumber)
			}
		}
	}
	return result.ToSlice()
}

func LoadEngineSchematic(scanner *bufio.Scanner) *EngineSchematic {
	schema := new(EngineSchematic)
	schema.Parts = map[Point]int{}
	schema.Symbols = map[Point]rune{}

	numberRegexp, _ := regexp.Compile("[0-9]+")

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		schema.grid = append(schema.grid, line)

		for _, loc := range numberRegexp.FindAllStringIndex(line, -1) {
			for j := loc[0]; j < loc[1]; j++ {
				schema.Parts[*NewPoint(i, j)], _ = strconv.Atoi(line[loc[0]:loc[1]])
			}
		}

		for j, c := range line {
			if strings.ContainsRune("+%@&$#*/=-", c) {
				schema.Symbols[*NewPoint(i, j)] = c
			}
		}
	}
	return schema
}
