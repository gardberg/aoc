package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Short variable declaration ':=' cannot be used outside of 
// a function scope!
var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var convert = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var winner_to = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/2/input.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += get_score(string(line[0]), string(line[2]))
	}

	fmt.Println(sum)

}

func get_score(a string, b string) int {
	score := 0
	score += points[b]

	if b == convert[a] {
		// draw
		score += 3
	} else if b == winner_to[a] {
		// win
		score += 6
	} else {
		score += 0
	}

	return score



}