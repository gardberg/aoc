package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/15.2/input.txt")

	sensors, beacons := get_input(file)

	pos_lines := make(map[int]bool)
	neg_lines := make(map[int]bool)

	dists := make([]int, 0)
	for i := range sensors {
		dists = append(dists, dist(sensors[i], beacons[i]))
	}

	for i, p := range sensors {
		d := dist(sensors[i], beacons[i])
		//pos_lines = append(pos_lines, p.y-p.x+d+1)
		//pos_lines = append(pos_lines, p.y-p.x-d-1)
		pos_lines[p.y-p.x+d+1] = true
		pos_lines[p.y-p.x-d-1] = true

		//neg_lines = append(neg_lines, p.y+p.x+d+1)
		//neg_lines = append(neg_lines, p.y+p.x-d-1)
		neg_lines[p.y+p.x+d+1] = true
		neg_lines[p.y+p.x-d-1] = true
	}

	bound := 4000000
	//bound := 20
	for p1 := range pos_lines {
		for p2 := range neg_lines {
			p := pos{x: (p2 - p1)/2, y: (p1 + p2)/2}

			// if p is in bounds
			if p.x > 0 && p.x < bound && p.y > 0 && p.y < bound {
				// if all sensors are futher than d from p
				// then p is a valid point
				valid := true
				for i := range sensors {
					if dist(sensors[i], p) <= dists[i] {
						valid = false
						break
					}
				}

				if valid {
					fmt.Println(4000000*p.x+p.y)
				}
				
			}

		}
	}
	

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

func dist(a, b pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}