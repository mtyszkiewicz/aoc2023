package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parse(line string) int {
	l := strings.IndexFunc(line, unicode.IsDigit)
	if l == -1 {
		return 0
	}
	r := strings.LastIndexFunc(line, unicode.IsDigit)
	result, _ := strconv.Atoi(string(line[l]) + string(line[r]))
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		result := parse(line)
		total += result
	}
	fmt.Println(total)
}
