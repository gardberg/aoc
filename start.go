package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/3/test.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	
}