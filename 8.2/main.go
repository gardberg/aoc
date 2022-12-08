package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
)

var rows int
var cols int

// How many trees are visible from the edge?
func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/8.2/input8.txt")

	var grid [][]int
	grid, _, rows, cols = make_grid(file)

	score := 0

	for i, r := range grid[1 : len(grid)-1] {
		for j := range r[1 : len(r)-1] {
			temp_score := get_score(grid, i+1, j+1)
			if temp_score > score {
				score = temp_score
			}
			//fmt.Printf("%v at (%v, %v) visible: %v", grid[i+1][j+1], i+1, j+1, ok)

			//fmt.Println()
		}
	}

	fmt.Println(score)

}

func make_grid(file []uint8) ([][]int, int, int, int) {
	n_rows := 0
	n_cols := 0
	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {

		n_cols = len(row)
		n_rows += 1
	}

	outer := 2*n_rows + 2*n_cols - 4

	grid := make([][]int, n_rows)
	for i := range grid {
		grid[i] = make([]int, n_cols)
	}

	for i, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		for j, c := range row {
			grid[i][j] = int(c - 48)
		}
	}

	return grid, outer, n_rows, n_cols
}

func get_score(grid [][]int, row int, col int) int {
	//visible := true

	score_l := 1
	// left
	for l := col; l >= 2; l-- {
		if grid[row][col] <= grid[row][l-1] {
			break
		}
		score_l += 1
	}

	// up
	score_u := 1
	
	for u := row; u >= 2; u-- {
		if grid[row][col] <= grid[u-1][col] {
			break
		}
		score_u += 1
	}

	// right
	score_r := 1
	for r := col; r + 3 <= cols; r++ {
		if grid[row][col] <= grid[row][r+1] {
			break
		}
		score_r += 1
	}

	// down
	score_d := 1
	for d := row; d+3 <= rows; d++ {
		if grid[row][col] <= grid[d+1][col] {
			break
		}
		score_d += 1
	}

	return score_l * score_u * score_d * score_r
}
