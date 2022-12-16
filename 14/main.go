package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	//"time"
)

var xs []int
var ys []int

var points [][]string

var min_xs, max_xs, min_ys, max_ys int

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/14/input.txt")

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		// split row on ","
		l := get_points(row)
		points = append(points, l)
		fmt.Println()
	}
	// print min and max of xs and ys
	min_xs, max_xs = min_max(xs)
	min_ys, max_ys = min_max(ys)
	fmt.Printf("min x: %v, max x; %v\n", min_xs, max_xs)
	fmt.Printf("min y: %v, max y: %v\n", min_ys, max_ys)
	
	grid := get_grid(max_xs-min_xs, max_ys)

	grid[0][500-min_xs] = "+"

	for _, line := range points {
		grid = draw_line(grid, line)
	}

	print_grid(grid)
	fmt.Println()

	count := 0
	ok := true

	for ok {
		grid, ok = drop_sand(grid)
		if ok {
			count += 1
		}
		//print_grid(grid)
		//fmt.Println()
		// wait one second
		//time.Sleep(time.Second/2)
	}

	print_grid(grid)

	fmt.Println(count)

}

func drop_sand(grid [][]string) ([][]string, bool) {
	/*
	A unit of sand always falls down one step if possible. If the tile immediately below is blocked
	(by rock or sand), the unit of sand attempts to instead move diagonally one step down and to the left.
	If that tile is blocked, the unit of sand attempts to instead move diagonally one step down 
	and to the right. Sand keeps moving as long as it is able to do so, at each step trying to move down,
	then down-left, then down-right. If all three possible destinations are blocked, the unit of sand 
	comes to rest and no longer moves, at which point the next unit of sand is created back at the source.
	*/

	i := 1
	j := 500-min_xs
	for i < len(grid) {
		// check if we are out of bounds of grid
		if i >= len(grid) || j >= len(grid[i]) || j < 0 {
			return grid, false
		}


		if grid[i][j] == "#" || grid[i][j] == "o" {
			// check down and left
			if j-1 < 0 {
				return grid, false
			}
			if grid[i][j-1] == "." {
				j--
			} else if j+1 > len(grid[i]) {
				return grid, false
			} else if grid[i][j+1] == "." {
				j++
			} else {
				// both left and right blocked
				// place sand
				grid[i-1][j] = "o"
				return grid, true
			}

		
		}
		i++
	}

	return grid, false
}

func draw_line(grid [][]string, points []string) [][]string {
	// check if line is horizontal or vertical

	// (x, y) represents (col, row)
	// x: right, y: down
	for i := range points[:len(points)-1] {
		startx, _ := strconv.Atoi(strings.Split(points[i], ",")[0])
		starty, _ := strconv.Atoi(strings.Split(points[i], ",")[1])
		stopx, _ := strconv.Atoi(strings.Split(points[i+1], ",")[0])
		stopy, _ := strconv.Atoi(strings.Split(points[i+1], ",")[1])

		//fmt.Printf("Drawing line from (%v, %v) to (%v, %v)\n", startx, starty, stopx, stopy)

		startx -= min_xs
		stopx -= min_xs

		if startx == stopx {
			// vertical line

			// if going up
			if starty > stopy {
				starty, stopy = stopy, starty
			}

			for i := starty; i <= stopy; i++ {
				grid[i][startx] = "#"
			}
		} else {

			// if going left
			if startx > stopx {
				startx, stopx = stopx, startx
			}
			for i := startx; i <= stopx; i++ {
				grid[starty][i] = "#"
			}
		} 
	}

	return grid
}

func get_grid(x, y int) [][]string {
	fmt.Printf("Making grid of size %v x %v\n", x, y)
	grid := make([][]string, y+1)
	for i := range grid {
		// fill with "."
		grid[i] = make([]string, x+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	return grid
} 

func print_grid(grid [][]string) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%v ", c)
		}
		fmt.Println()
	}
}

func get_points(lines string) []string{
	l := strings.Split(lines, " -> ")
	for _, line := range l {
		xt, _ := strconv.Atoi(strings.Split(line, ",")[0])
		yt, _ := strconv.Atoi(strings.Split(line, ",")[1])
		xs = append(xs, xt)
		ys = append(ys, yt)
	}

	return l
}

func min_max(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}