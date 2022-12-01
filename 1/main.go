package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	//"strings"
)

// sum all lists of integers separated by whitespaces
// return the largest sum
func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/1/test.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

	sums := make([]int, 1)
	var i int = 0

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if line == "" {
			i += 1
			//fmt.Println("Space at line", i)
			sums = append(sums, 0)
		}

		current_int, _ := strconv.Atoi(line)
		sums[i] += current_int
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	largest := -1
	for _, n := range(sums) {
		if n > largest {
			largest = n
		}
	}

fmt.Println(largest)

}