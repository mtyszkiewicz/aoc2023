package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Network map[string][]string

func (net *Network) Capture(values []string) error {
	if *net == nil {
		*net = make(map[string][]string)
	}
	(*net)[values[0]] = []string{"xd"}
	return nil
}

type DesertMap struct {
	Instructions string   `parser:"@Instr"`
	Network      *Network `parser:"@Map*"`
}

func main() {
	lexer := lexer.MustSimple([]lexer.SimpleRule{
		{"Instr", `[A-Z]+\n`},
		{"Whitespace", `[ \t\r\n]+`},
		{"Map", `([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`},
	})
	parser := participle.MustBuild[DesertMap](
		participle.Lexer(lexer),
		participle.Elide("Whitespace"),
	)
	desertMap, err := parser.Parse("", os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(desertMap.Network)
}
