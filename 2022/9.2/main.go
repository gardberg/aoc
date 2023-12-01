package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type point struct {
	x int
	y int
 }

var pos_set map[point]bool

var N int = 10

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/9.2/input9.txt")

	pos_set = make(map[point]bool)
	pos_set[point{x: 0, y: 0}] = true

	tails := make([]point, N)

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		dir := row[0]
		dist, _ := strconv.Atoi(row[2:])
		//fmt.Println(string(dir), dist)

		for dist > 0 {
			//past_point := tails[0]
			tails[0] = update_head(string(dir), tails[0])

			for i := range tails[:len(tails)-1] {

				// tails[i+1], past_point = update_tail(tails[i], tails[i+1], past_point)
				tails[i+1] = update_tail(tails[i+1], tails[i])
			}

			dist--
			pos_set[tails[len(tails)-1]] = true
		}
		
	}

	fmt.Println(len(pos_set))
	// fmt.Print(pos_set)
	// fmt.Println()
	// print_matrix()

}

func is_adj(head, tail point) bool {
	dx := abs_diff(head.x, tail.x)
	dy := abs_diff(head.y, tail.y)

	return (dx <= 1 && dy <= 1)
}

func abs_diff(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

 func update_head(dir string, head point) point {
	new_x := head.x
	new_y := head.y

	switch dir {
	case "R":
		new_x += 1
	case "L":
		new_x -= 1
	case "U":
		new_y += 1
	case "D":
		new_y -= 1
	}

	return point{x: new_x, y: new_y}
 }


func update_tail2(head, tail, past_point point) (point, point) {
	new_tail := tail
	new_past_point := past_point
	if !is_adj(head, tail) {
		// update the tail to past_point
		new_past_point = point{x: past_point.x, y: past_point.y}
		new_tail = point{x: past_point.x, y: past_point.y}
	}

	return new_tail, new_past_point
}

// can't we just use the one infronts past position? apparently not
func update_tail(tail point, head point)(newTail point){
	newTail = tail
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,1},point{-1,2}, point{0,2}, point{1,2}, point{2,1}, point{2,2}, point{-2,2}:
		newTail.y++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{1,2},point{2,1}, point{2,0}, point{2,-1}, point{1,-2}, point{2,2}, point{2,-2}:
		newTail.x++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{2,-1},point{1,-2}, point{0,-2}, point{-1,-2}, point{-2,-1}, point{2,-2}:
		newTail.y--
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{-1,-2},point{-2,-1}, point{-2,-0}, point{-2,1}, point{-1,2}, point{-2,2}:
		newTail.x--
	}
	return
}