package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/10/input10.txt")

	X := 0
	cycle := 0

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		//fmt.Printf("register: %v, cycle: %v", X, cycle)
		//fmt.Println()
		s, v := get_op(row)
		//fmt.Println(s, v)

		X, cycle = execute(s, v, X, cycle)

	}

}	

func execute(op string, v int, X int, cycle int) (int, int) {

	if op == "noop" {
		draw(X, cycle)
		cycle += 1
	} else {
		for i := 0; i < 2; i++ {
			draw(X, cycle)
			cycle += 1
		}
		X += v
	}

	return X, cycle

}


func draw(X, cycle int) {
	if cycle > 0 && cycle % 40 == 0 {
		fmt.Println()
	}

	if cycle % 40 >= X && cycle % 40 <= X + 2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}


func get_op(line string) (string, int) {
	op := line[:4]
	v := 0
	
	if len(line) > 4 {
		v, _ = strconv.Atoi(line[5:])
	}
	
	return op, v
}