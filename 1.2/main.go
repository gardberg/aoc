package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"sort"
	//"strings"
)

// sum all lists of integers separated by whitespaces
// return the *three largest sums*
func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/1.2/input.txt")
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

	sort.Ints(sums)

	top3_sum := 0
	for _, n := range sums[len(sums)-3:] {
		top3_sum += n
	}

	fmt.Println(top3_sum)

}