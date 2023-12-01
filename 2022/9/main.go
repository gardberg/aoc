package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)



type Pos struct {
	key int
	value int
 }

var pos_set map[Pos]bool

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/9/input9.txt")

	pos_set = make(map[Pos]bool)
	pos_set[Pos{key: 0, value: 0}] = true

	head_x := 0
	head_y := 0

	tail_x := 0
	tail_y := 0

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		dir := row[0]
		dist, _ := strconv.Atoi(row[2:])
		fmt.Println(string(dir), dist)

		head_x, head_y, tail_x, tail_y = update_head_iter(string(dir), dist, head_x, head_y, tail_x, tail_y)
		
	}

	fmt.Println(len(pos_set))

}

func is_adj(head_x, head_y, tail_x, tail_y int) bool {
	dx := abs_diff(head_x, tail_x)
	dy := abs_diff(head_y, tail_y)

	return (dx <= 1 && dy <= 1)
}

func abs_diff(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

 func update_head(dir string, dist, head_x, head_y int) (int, int) {
	new_x := head_x
	new_y := head_y

	switch dir {
	case "R":
		new_x += dist
	case "L":
		new_x -= dist
	case "U":
		new_y += dist
	case "D":
		new_y -= dist
	}

	return new_x, new_y
 }

func update_head_iter(dir string, dist, hx, hy, tx, ty int) (int, int, int, int) {

	head_x, head_y := hx, hy
	tail_x, tail_y := tx, ty

	for i := 0; i < dist; i++ {
		head_past_x, head_past_y := head_x, head_y
		head_x, head_y = update_head(dir, 1, head_x, head_y)

		if !is_adj(head_x, head_y, tail_x, tail_y) {
			tail_x, tail_y = head_past_x, head_past_y
			p := Pos{key: tail_x, value: tail_y}

			// update pos set
			pos_set[p] = true

		}

		fmt.Printf("Head: (%v, %v), Tail: (%v, %v)", head_x, head_y, tail_x, tail_y)
		fmt.Println()

	}

	return head_x, head_y, tail_x, tail_y
}