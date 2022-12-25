package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type pos struct {
	x int
	y int
}

var minx, miny, maxx, maxy int

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/15/test.txt")

	sensors, beacons := get_input(file)

	//fmt.Println(sensors)
	//fmt.Println(beacons)

	// Print the largest and smallest x and y values in
	// sensors and beacons

	// x (right), y (down)

	y := 2000000

	filled := map[int]bool{}

	for i := range sensors {
		d := dist(sensors[i], beacons[i])

		// if sensor is within d of line
		if d >= abs(sensors[i].y - y) {
			d -= abs(sensors[i].y - y)

			for j := sensors[i].x - d; j <= sensors[i].x + d; j++ {
				filled[j] = true
			}
		}

	}
	// Remove counted beacons
	for i := range beacons {
		if beacons[i].y == y {
			delete(filled, beacons[i].x)
		}
	}

	fmt.Println(len(filled))

	
	
	

}

// Get x and y from input of form
// "Sensor at x=9, y=16: closest beacon is at x=10, y=16"
func get_input(file []byte) (sensors []pos, beacons []pos) {
	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {

		rs := strings.Split(row, " ")
		sx, _ := strconv.Atoi(rs[2][2:len(rs[2])-1])
		sy, _ := strconv.Atoi(rs[3][2:len(rs[3])-1])

		bx, _ := strconv.Atoi(rs[8][2:len(rs[8])-1])
		by, _ := strconv.Atoi(rs[9][2:len(rs[9])])

		sensors = append(sensors, pos{
			x: sx,
			y: sy,
		})
		beacons = append(beacons, pos{
			x: bx,
			y: by,
		})
	}

	return sensors, beacons
}

// function to get the largest and smallest x and y values
// in an array of positions
func get_min_max(sensors, beacons []pos) (minx, miny, maxx, maxy int) {
	minx = sensors[0].x
	miny = sensors[0].y
	maxx = sensors[0].x
	maxy = sensors[0].y

	for i, s := range sensors {
		d := dist(s, beacons[i])
		minx = min(minx, s.x-d)
		miny = min(miny, s.y-d)
		maxx = max(maxx, s.x+d)
		maxy = max(maxy, s.y+d)
		
	}

	return minx, miny, maxx, maxy
}

func make_grid(sensors, beacons []pos) [][]bool {
	minx, miny, maxx, maxy = get_min_max(sensors, beacons)

	fmt.Printf("minx: %d, maxx: %d, miny: %d, maxy: %d\n", minx, maxx, miny, maxy)

	// grid := make([][]bool, maxy-miny+1)
	// for i := range grid {
	// 	fmt.Printf("Making row %v of %v\n", i+1, len(grid))
	// 	grid[i] = make([]bool, maxx-minx+1)
	// }
	n := maxy - miny + 1
	m := maxx - minx + 1

	grid := make([][]bool, n)
	rows := make([]bool, n*m)
	for i := 0; i < n; i++ {
		fmt.Printf("Making row %v of %v\n", i+1, n)
		grid[i] = rows[i*m : (i+1)*m]
	}

	return grid

}

func print_grid(grid [][]bool) {

	for i, row := range grid {
		if i+miny >= 0 && i+miny < 10 {
			fmt.Print(i+miny, "   ")
		} else if i+miny < -9 {
			fmt.Print(i+miny, " ")
		} else {
			fmt.Print(i+miny, "  ")
		}
		for _, col := range row {

			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func fill_grid(grid [][]bool, sensors, beacons []pos) [][]bool {
	l := len(sensors)

	for i, s := range sensors {
		fmt.Printf("On sensor %v of %v\n", i+1, l)
		b := beacons[i]
		d := dist(s, b)

		grid[s.y-miny][s.x-minx] = true

		// fill all squares within distance d of sensor
		for y := s.y - d; y <= s.y+d; y++ {
			for x := s.x - d; x <= s.x+d; x++ {
				if dist(s, pos{x, y}) <= d {
					grid[y-miny][x-minx] = true
				}
			}
		}

	}

	// set all beacon positions to false
	for _, b := range beacons {
		grid[b.y-miny][b.x-minx] = false
	}

	return grid

}

func dist(a, b pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}