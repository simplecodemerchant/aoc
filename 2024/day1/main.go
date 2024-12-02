package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed 2024_1.txt
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
	re := regexp.MustCompile("[ ]+")

	lines := strings.Split(input, "\n")

	line_len := len(lines)

	left := make([]int, line_len)
	right := make([]int, line_len)

	for i := 0; i < line_len; i++ {
		line := lines[i]
		vals := re.Split(line, -1)

		lval, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}

		rval, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		left[i] = lval
		right[i] = rval
	}

	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := 0; i < line_len; i++ {
		total = total + Abs(left[i], right[i])
	}

	fmt.Println(total)
}

func Abs(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// func p2(input string) {
// }
