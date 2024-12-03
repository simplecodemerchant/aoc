package main

import (
	_ "embed"
	"flag"
	"fmt"
	"iter"
	"regexp"
	"strconv"
	"strings"
)

//go:embed 2024_2.txt
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

func getLines(input string) iter.Seq[bool] {
	return func(yield func(bool) bool) {
		re := regexp.MustCompile("[ ]+")

		lines := strings.Split(input, "\n")
		for _, line := range lines {
			good := true
			up := true

			data := re.Split(line, -1)

			prev, _ := strconv.Atoi(data[0])
			next, _ := strconv.Atoi(data[1])

			if prev > next {
				up = false
			}

			for _, d := range data[1:] {
				num, _ := strconv.Atoi(d)

				if num == prev {
					good = false
					break
				} else if up && (num-prev > 3 || num-prev < 0) {
					good = false
					break
				} else if !up && (prev-num > 3 || prev-num < 0) {
					good = false
					break
				}
				prev = num
			}

			if !yield(good) {
				return
			}
		}
	}

}

func p1(input string) {
	total := 0
	for good := range getLines(input) {
		if good {
			total++
		}
	}
	fmt.Println(total)
}

// func p2(input string) {
// }
