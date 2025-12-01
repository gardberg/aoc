package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Pair struct {
	x, y int
}

func main() {
	lines := utils.GetLines("input.txt")
	grid := make([][]string, len(lines))

	// Contains all adjacent numbers for each symbol *
	gear_map := make(map[Pair][]int)
	s := 0

	re := regexp.MustCompile(`\d+`)
	gear_re := regexp.MustCompile(`\*`)

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

			// look for characters: * and get their indices

			startY, stopY := match[0], match[1]
			nbr := line[startY:stopY]

			if x > 0 {
				substr, startYC := checkGrid(grid, x-1, startY-1, stopY+1)
				found := gear_re.FindAllStringSubmatchIndex(substr, -1)	
				for _, gear_idx := range found {
					gear_map[Pair{x-1, gear_idx[0] + startYC}] = append(gear_map[Pair{x-1, gear_idx[0] + startYC}], toNumber(toString(nbr)))
				}
			}

			if x < len(grid)-1 {
				substr, startYC := checkGrid(grid, x+1, startY-1, stopY+1)
				found := gear_re.FindAllStringSubmatchIndex(substr, -1)	
				for _, gear_idx := range found {
					gear_map[Pair{x+1, gear_idx[0] + startYC}] = append(gear_map[Pair{x+1, gear_idx[0] + startYC}], toNumber(toString(nbr)))
				}
			}

			if startY > 0 {
				substr, startYC := checkGrid(grid, x, startY-1, stopY-1)
				found := gear_re.FindAllStringSubmatchIndex(substr, -1)	
				for _, gear_idx := range found {
					gear_map[Pair{x, gear_idx[0] + startYC}] = append(gear_map[Pair{x, gear_idx[0] + startYC}], toNumber(toString(nbr)))
				}
			}

			if stopY < len(line)-1 {
				substr, startYC := checkGrid(grid, x, startY+1, stopY+1)
				found := gear_re.FindAllStringSubmatchIndex(substr, -1)	
				for _, gear_idx := range found {
					gear_map[Pair{x, gear_idx[0] + startYC}] = append(gear_map[Pair{x, gear_idx[0] + startYC}], toNumber(toString(nbr)))
				}
			}

			
		}
	}

	for _, v := range gear_map {
		if len(v) == 2 {
			s += v[0] * v[1]
		}
	}

	fmt.Println(s)

}

func checkGrid(grid [][]string, x, startY, stopY int) (string, int) {
	// make sure indices are not out of range
	xStart := max(0, x)
	startYC := max(0, startY)
	stopYC := min(stopY, len(grid[0]))

	return toString(grid[xStart][startYC:stopYC]), startYC
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
