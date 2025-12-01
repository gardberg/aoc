package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

var numbers = map[string]int{
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func main() {
	lines := utils.GetLines("input.txt")
	s := 0
	for _, line := range(lines) {
		// fmt.Println(line)

		number := getNumber(line)
		s += number
	}

	fmt.Println(s)
	
}

func getNumber(line string) int {
	first := getFirstNumber(line)
	//fmt.Println(first)
	last := getLastNumber(line)
	//fmt.Println(last)

	number_str := first + last
	number_int, _ := strconv.Atoi(number_str)

	return number_int
}

func getFirstNumber(line string) string {
	// loop over characters
	for i, char := range(line) {
		// check if character is a number using strconv
		if _, err := strconv.Atoi(string(char)); err == nil { 
			return string(char)
		}
		// check if substring from this char exists in "numbers" dict
		for key, value := range(numbers) {
			if i + len(key) <= len(line) && line[i:i+len(key)] == key {
				return strconv.Itoa(value)
			}
		}

	}

	return "NOT FOUND"
}

func getLastNumber(line string) string {
	// loop over chars in line backwards
	for i := len(line)-1; i >= 0; i-- {
		char := line[i]
		if _, err := strconv.Atoi(string(char)); err == nil { 
			return string(char)
		}

		for key, value := range(numbers) {
			if i >= len(key) - 1 && line[i-len(key)+1:i+1] == key {
				return strconv.Itoa(value)
			}
		}

	}

	return "NOT FOUND"
}
