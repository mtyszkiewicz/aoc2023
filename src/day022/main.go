package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

type ColorCount struct {
	Count int    `parser:"@Int"`
	Color string `parser:"@('red' | 'green' | 'blue')"`
}

type Draw struct {
	ColorCounts []*ColorCount `parser:"@@ (',' @@)*"`
}

type Game struct {
	Id    int     `parser:"'Game' @Int':'"`
	Draws []*Draw `parser:"@@ (';' @@)*"`
}

func (game *Game) IsPossible(maxColors map[string]int) bool {
	for _, draw := range game.Draws {
		for _, colorCount := range draw.ColorCounts {
			if colorCount.Count > maxColors[colorCount.Color] {
				return false
			}
		}
	}
	return true
}

func (game *Game) Power() int {
	maxColors := make(map[string]int)

	for _, draw := range game.Draws {
		for _, colorCount := range draw.ColorCounts {
			if colorCount.Count > maxColors[colorCount.Color] {
				maxColors[colorCount.Color] = colorCount.Count
			}
		}
	}

	result := 1
	for _, maxCount := range maxColors {
		result *= maxCount
	}
	return result
}

func main() {
	parser := participle.MustBuild[Game]()
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		game, err := parser.ParseString("", scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		result += game.Power()
	}
	fmt.Println(result)
}
