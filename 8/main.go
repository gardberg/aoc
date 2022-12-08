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
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/8/input8.txt")

	var grid [][]int
	var outer int
	grid, outer, rows, cols = make_grid(file)

	n_visible := 0
	n_visible += outer

	for i, r := range grid[1 : len(grid)-1] {
		for j := range r[1 : len(r)-1] {
			ok := is_visible(grid, i+1, j+1)
			if ok {
				n_visible += 1
			}
			//fmt.Printf("%v at (%v, %v) visible: %v", grid[i+1][j+1], i+1, j+1, ok)

			//fmt.Println()
		}
	}

	fmt.Println(n_visible)

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

func is_visible(grid [][]int, row int, col int) bool {
	//visible := true

	// left
	l := col
	for {
		if grid[row][col] <= grid[row][l-1] {
			break
		}
		if l < 2 {
			return true
		}
		l--
	}

	// up
	u := row
	for {
		if grid[row][col] <= grid[u-1][col] {
			break
		}
		if u < 2 {
			return true
		}
		u--
	}

	// right
	r := col
	for {
		if grid[row][col] <= grid[row][r+1] {
			break
		}
		if r+3 > cols {
			return true
		}
		r++
	}

	// down
	d := row
	for {
		if grid[row][col] <= grid[d+1][col] {
			break
		}
		if d+3 > rows {
			return true
		}
		d++
	}

	return false
}
