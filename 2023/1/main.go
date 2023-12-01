package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	s := 0

	for scanner.Scan() {
		line := scanner.Text()
		number := ""
		for _, c := range(line) {
			// If the character is an int, append to number
			if c >= '0' && c <= '9' {
				number += string(c)
			}
		}
		
		// Create a new string of the first and last character in number
		to_add := string(number[0]) + string(number[len(number)-1])
		to_add_int, _ := strconv.Atoi(to_add)
		s += to_add_int

		
	}
	fmt.Println(s)


}