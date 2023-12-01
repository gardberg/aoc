package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var stacks [][]string

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/5/input5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks = make([][]string, 9)
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}

	row := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)

		if row < 9 {
			append_line(line)
		}

		if row == 9 {
			for _, s := range stacks {
				reverse(s)
			}
		}

		if row > 10 {
			move(line)
		}

		row++
	}

	for _, s := range stacks {
		fmt.Println(s[len(s)-1])
	}



}

func append_line(line string) {
	for i := range stacks {
		c := string(line[i*4 + 1])

		if c != " " {
			stacks[i] = append(stacks[i], c)
		}
	}
}

func move(line string) {
	subs := strings.Split(line, " ")
	n, _ := strconv.Atoi(subs[1])
	from, _ := strconv.Atoi(subs[3])
	to, _ := strconv.Atoi(subs[5])

	var temp string

	for i := 0; i < n; i++ {
		stacks[from-1], temp = pop(stacks[from-1])
		stacks[to-1] = push(stacks[to-1], temp)
	}
}

func reverse(ss []string) {
    last := len(ss) - 1
    for i := 0; i < len(ss)/2; i++ {
        ss[i], ss[last-i] = ss[last-i], ss[i]
    }
}

func push(s []string, v string) []string {
    return append(s, v)
}

func pop(s []string) ([]string, string) {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    return  s[:l-1], s[l-1]
}