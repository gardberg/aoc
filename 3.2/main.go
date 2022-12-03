package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	//"strconv"
)

var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// find the item (letter) common in three different lines
func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/3.2/input.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	item_map := create_item_map(items)

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()

		line2 := scanner.Text()
		scanner.Scan()

		line3 := scanner.Text()

		c := get_common_item(line1, line2, line3)
		
		sum += item_map[c]

	}

	fmt.Println(sum)
}


// Return a string of all characters which exist in both strings
func get_duplicate_items(line1 string, line2 string) map[rune]bool {

	dup_chars := make(map[rune]bool)

	line1_count := make(map[rune]int, len(line1))

	for _, c := range line1 {
		line1_count[c] = 1
	}

	for _,c := range line2 {
		// if the character from line2 is in line1 map
		if _, ok := line1_count[c]; ok {
			dup_chars[c] = true
			//break
		}
	}

	return dup_chars

}

func get_common_item(line1 string, line2 string, line3 string) rune {
	// Get a set (map) containing all duplicate items
	c12 := get_duplicate_items(line1, line2)
	c23 := get_duplicate_items(line2, line3)
	c13 := get_duplicate_items(line1, line3)

	for c := range c12 {
		if _, ok1 := c23[c]; ok1 {
			if _, ok2 := c13[c]; ok2 {
				//fmt.Printf("%c", c)
				return c
			}
		}
	}

	return 0
}

func create_item_map(items string) map[rune]int {
	item_map := make(map[rune]int)
	for i, c := range items {
		item_map[c] = i+1
	}

	return item_map
}