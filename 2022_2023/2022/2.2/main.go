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

var loser_to = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var winner_to = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var draw_to = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/2.2/test.txt")
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

	var chosen string

	if b == "X" {
		// lose
		chosen = loser_to[a]
		score += 0
	} else if b == "Y" {
		// draw
		chosen = draw_to[a]
		score += 3
	} else {
		// win
		chosen = winner_to[a]
		score += 6
	}

	score += points[chosen]

	return score



}