package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Called after getLines returns

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output

}

func ToInt(s []string) []int {
	var output []int
	for _, v := range(s) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, i)
	}
	return output
}
