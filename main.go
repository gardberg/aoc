package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/2022/8.2/input8.txt")

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		fmt.Println(row)
	}

}
