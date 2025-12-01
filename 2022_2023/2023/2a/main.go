package main

import (
	"fmt"
	"aoc/utils"
	"strings"
	"strconv"
)

var limits = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	lines := utils.GetLines("input.txt")
	// Example line:
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	s := 0

	for _, game := range(lines) {
		id, rounds := splitLine(game)
		id_s, _ := strconv.Atoi(id)
		fmt.Println(id, rounds)

		valid := true
		for _, round := range(rounds) {
			round_map := parseRound(round)

			for color, count := range(round_map) {
				if count > limits[color] {
					fmt.Println("Invalid game", id_s)
					valid = false
					break	
				}
			}	

			if !valid {
				break
			}


		}

		if valid {
			s += id_s
		}

	}
	fmt.Println(s)
}

// return game id and string array of games
func splitLine(line string) (string, []string) {
	s := strings.Split(line, ":")	
	id := strings.Split(s[0], " ")[1]

	rounds := strings.Split(s[1][1:], "; ")
	return id, rounds 
}

// Return dict of (string, ing)
func parseRound(line string) map[string]int {
	// "3 blue, 4 red"
	count := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}
	die := strings.Split(line, ", ")
	
	for _, d := range(die) {
		s := strings.Split(d, " ")
		n, _ := strconv.Atoi(s[0])
		color := s[1]
		count[color] += n	

	}
	return count
}
