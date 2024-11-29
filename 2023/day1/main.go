package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed 2023_1.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var p int
	flag.IntVar(&p, "p", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", p)

	if p == 1 {
		p1(input)
	}
	// else {
	// 	p2(input)
	// }
}

func p1(input string) {
	total := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line := []rune(line)
		ll := len(line)
		first := ""
		last := ""

		for f, l := 0, ll-1; f < ll; f, l = f+1, l-1 {

			if first == "" && unicode.IsDigit(line[f]) {
				first = string(line[f])
			}
			if last == "" && unicode.IsDigit(line[l]) {
				last = string(line[l])
			}
			if first != "" && last != "" {
				break
			}
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		total += num

	}
	fmt.Println(total)
}

// func p2(input string) {
// }
