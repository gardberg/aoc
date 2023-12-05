package main

import (
	"fmt"
	"aoc/utils"
	"strings"
	"regexp"
)

var nbr_lines int

func main() {
	lines := utils.GetLines("input.txt")

	winnings, nbrs := parseInput(lines)
	nbr_lines = len(nbrs)
	total_nbr_scratchcards := 0

	for i := 0; i < nbr_lines; i++ {
		s := countScratchcards(i, winnings, nbrs)
		total_nbr_scratchcards += s
	}

	fmt.Println(total_nbr_scratchcards + nbr_lines)

}

func countScratchcards(card_nbr int, winnings [][]int, nbrs [][]int) int {

	s := 0

	copies := countWinning(winnings[card_nbr], nbrs[card_nbr])

	if copies == 0 {
		return 0
	}

	for i := 0; i < copies; i++ {
		s += countScratchcards(card_nbr + 1 + i, winnings, nbrs) + 1
	}

	return s

}

func countWinning(winnings []int, nbrs []int) int {
	count := 0

	for _, winning := range winnings {
		for _, nbr := range nbrs {
			if winning == nbr {
				count++
			}
		}
	}

	return count
}

func parseInput(lines []string) ([][]int, [][]int) {
	winningsOut := make([][]int, 0)
	nbrsOut := make([][]int, 0)
	
	for _, line := range(lines) {
		sp := strings.Split(line, "|")
		winning := strings.Split(sp[0], ": ")[1]
		nbrs := sp[1:]

		nbr_regex := regexp.MustCompile(`\d+`)

		winningArr := nbr_regex.FindAllString(winning, -1)
		nbrsArr := nbr_regex.FindAllString(nbrs[0], -1)

		winningInts := utils.ToInt(winningArr)
		nbrsInts := utils.ToInt(nbrsArr)

		winningsOut = append(winningsOut, winningInts)
		nbrsOut = append(nbrsOut, nbrsInts)
	}

	return winningsOut, nbrsOut
}