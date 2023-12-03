package main

import (
	"fmt"
	"aoc/utils"
)

func main() {
	lines := utils.GetLines("test.txt")

	for _, line := range(lines) {
		fmt.Println(line)
	}
}