package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	//"strconv"
)

var items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// For each line we want to find the letter which exists
// in both the first and second part of the string
func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/3/input.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	item_map := create_item_map(items)

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		
		//fmt.Println(line)
		dup_char := get_duplicate_item(line)

		//fmt.Printf("Char %c has value %v\n", dup_char, item_map[dup_char])
		sum += item_map[dup_char]
		//fmt.Println(strconv.Itoa(dup_char))
	}

	fmt.Println(sum)
}

func get_duplicate_item(line string) rune {
	line1 := line[:len(line)/2]
	line2 := line[len(line)/2:]

	line1_count := make(map[rune]int, len(line1))

	for _, c := range line1 {
		line1_count[c] = 1
	}

	var dup_char rune

	for _,c := range line2 {
		// if the character from line2 is in line1 map
		if _, ok := line1_count[c]; ok {
			dup_char = c
			break
		}
	}

	return dup_char

}

func create_item_map(items string) map[rune]int {
	item_map := make(map[rune]int)
	for i, c := range items {
		item_map[c] = i+1
	}

	return item_map
}