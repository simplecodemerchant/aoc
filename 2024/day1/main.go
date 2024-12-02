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
	} else {
		p2(input)
	}
}

func getData(input string) (*[]int, *[]int, int, error) {

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
			return nil, nil, 0, err
		}

		rval, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, nil, 0, err
		}

		left[i] = lval
		right[i] = rval
	}

	slices.Sort(left)
	slices.Sort(right)

	return &left, &right, line_len, nil
}

func p1(input string) {
	left, right, line_len, err := getData(input)
	if err != nil {
		panic(err)
	}

	total := 0
	for i := 0; i < line_len; i++ {
		total = total + Abs((*left)[i], (*right)[i])
	}

	fmt.Println(total)
}

func Abs(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func p2(input string) {
	left, right, line_len, err := getData(input)
	if err != nil {
		panic(err)
	}

	total := 0

	left_num := 0
	right_num := 0

	l := 0
	r := 0
	similarity_score := 0

	for l < line_len && r < line_len {
		left_num = (*left)[l]
		right_num = (*right)[r]

		if left_num == right_num {
			similarity_score++
			r++
		} else if right_num > left_num {
			total += left_num * similarity_score
			similarity_score = 0
			l++
		} else {
			r++
		}
	}
	fmt.Println(total)
}
