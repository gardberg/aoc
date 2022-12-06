package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var bset map[byte]int

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/6/input6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bset = make(map[byte]int)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	//fmt.Println(line)

	fmt.Println(find_start(line))

}

func find_start(line string) int {
	//var buffer [4]byte
	buffer := make([]byte, 0)

	for i := 0; i < len(line); i++ {

		buffer = queue(buffer, line[i])
		//fmt.Println(bset)

		if is_unique(buffer) {
			return i+1
		}
		//fmt.Println(buffer)
		//return 0
	}
	return 0
}

func queue(arr []byte, b byte) []byte {
	if len(arr) > 3 {
		bset[arr[0]] -= 1

		if bset[arr[0]]==0{
			delete(bset, arr[0])
		}

		arr = arr[1:]
	}

	bset[b] += 1

	arr = append(arr, b)
	return arr
}

func is_unique(arr []byte) bool {
	return len(bset) == 4
}