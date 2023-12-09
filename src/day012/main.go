package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitsSpelled = strings.NewReplacer(
	"one", "1e",
	"two", "2o",
	"three", "3e",
	"four", "4r",
	"five", "5e",
	"six", "6x",
	"seven", "7n",
	"eight", "8t",
	"nine", "9e",
)

func parse(line string) int {
	line = digitsSpelled.Replace(line)
	line = digitsSpelled.Replace(line)
	l := strings.IndexFunc(line, unicode.IsDigit)
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
