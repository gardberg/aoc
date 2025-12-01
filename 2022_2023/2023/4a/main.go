package main

import (
	"fmt"
	"aoc/utils"
	"strings"
	"math"
	"regexp"
)

func main() {
	lines := utils.GetLines("input.txt")

	s := 0.0
	for i, line := range(lines) {
		score := 0
		sp := strings.Split(line, "|")
		winning := strings.Split(sp[0], ": ")[1]
		nbrs := sp[1:]

		nbr_regex := regexp.MustCompile(`\d+`)

		winningArr := nbr_regex.FindAllString(winning, -1)
		nbrsArr := nbr_regex.FindAllString(nbrs[0], -1)

		winningInts := utils.ToInt(winningArr)
		nbrsInts := utils.ToInt(nbrsArr)

		fmt.Println(winningInts)
		fmt.Println(nbrsInts)

		for _, nbr := range(nbrsInts) {
			for _, w := range(winningInts) {
				if nbr == w {
					score++
				}
			}
		}

		fmt.Printf("%d: %d\n", i, score)

		if score > 0 {
			s += math.Pow(2, float64(score)-1)
		}


	}

	fmt.Println(s)
}