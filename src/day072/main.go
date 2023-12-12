package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mtyszkiewicz/aoc2023/pkg/day07"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jockersAllowed := true

	players := []*day07.Player{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(data[1])
		players = append(players, day07.NewPlayer(data[0], bid, jockersAllowed))
	}

	slices.SortFunc(players, day07.ComparePlayers)
	result := 0
	for i, p := range players {
		result += (i + 1) * p.Bid
	}
	fmt.Println(result)
}
