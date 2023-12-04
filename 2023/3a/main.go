package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	lines := utils.GetLines("input.txt")
	grid := make([][]string, len(lines))

	re := regexp.MustCompile(`\d+`)
	symbols := regexp.MustCompile(`[^\w.]`)

	s := 0

	// Build grid
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	// Find all numbers
	for x, line := range grid {
		for _, match := range re.FindAllStringSubmatchIndex(toString(line), -1) {
			if len(match) == 0 {
				continue
			}

			// look for a symbol
			foundSymbol := false

			startY, stopY := match[0], match[1]
			nbr := line[startY:stopY]

			// check row above, A
			if x > 0 && symbols.MatchString(checkGrid(grid, x-1, startY-1, stopY+1)) {
				foundSymbol = true
			}

			// check row below, C
			if x < len(grid)-1 && symbols.MatchString(checkGrid(grid, x+1, startY-1, stopY+1)) {
				foundSymbol = true
			}

			// check column to the left, B
			if startY > 0 && symbols.MatchString(checkGrid(grid, x, startY-1, stopY-1)) {
				foundSymbol = true
			}

			// check column to the right, D
			if stopY < len(line)-1 && symbols.MatchString(checkGrid(grid, x, startY+1, stopY+1)) {
				foundSymbol = true
			}

			if foundSymbol {
				s += toNumber(toString(nbr))
			}
		}
	}

	fmt.Println(s)
}

func checkGrid(grid [][]string, x, startY, stopY int) string {
	// make sure indices are not out of range
	xStart := max(0, x)
	startYC := max(0, startY)
	stopYC := min(stopY, len(grid[0]))

	return toString(grid[xStart][startYC:stopYC])
}

func toNumber(s string) int {
	nbr, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return nbr
}

func toString(s []string) string {
	var str string
	for _, char := range s {
		str += char
	}
	return str
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
