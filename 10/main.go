package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

var strengths = make([]int, 0)

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
	//fmt.Printf("register: %v, cycle: %v", X, cycle)
	//fmt.Println()
	fmt.Println(strengths)
	sum := 0
	for _, v := range strengths {sum += v}
	fmt.Println(sum)

}	

func execute(op string, v int, X int, cycle int) (int, int) {

	if op == "noop" {
		cycle += 1
		update_strengths(X, cycle)
	} else {
		for i := 0; i < 2; i++ {
			cycle += 1
			update_strengths(X, cycle)
		}
		X += v
	}

	return X, cycle

}


func update_strengths(X int, cycle int) {
	if cycle > 0 && (cycle - 20) % 40 == 0 {
		strengths = append(strengths, (X+1) * cycle)
		fmt.Printf("Updating with %v * %v = %v", X, cycle, strengths[len(strengths)-1])
		fmt.Println()
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