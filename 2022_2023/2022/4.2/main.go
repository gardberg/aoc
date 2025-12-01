package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

// how many pair overlap?
func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/4.2/input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		p1, p2 := get_arrays(line)

		if o := overlap(p1,p2); o {
			count += 1
		}

		fmt.Println(line)

	}

	fmt.Printf("%v\n", count)

}

func get_arrays(line string) ([]int, []int) {
	line_split := strings.Split(line, "-")

	mid := strings.Split(line_split[1], ",")
	p1 := make([]int, 2)
	p2 := make([]int, 2)

	j, _ := strconv.Atoi(line_split[0])
	j2, _ := strconv.Atoi(mid[0])
	l, _ := strconv.Atoi(mid[1])
	l2, _ := strconv.Atoi(line_split[2])

	p1[0] = j
	p1[1] = j2
	p2[0] = l
	p2[1] = l2

	return p1, p2
}

func overlap(p1 []int, p2 []int) bool {
	if p1[0] == p2[0] {return true}

	if p1[0] < p2[0] {
		if p1[1] >= p2[0] {return true}

	} else if p2[0] < p1[0] {
		if p2[1] >= p1[0] {return true}
	}

	return false
}
