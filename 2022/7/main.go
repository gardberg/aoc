package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/lukgar/Desktop/repos/aoc/7/input7.txt") // <-- change me!
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dir_stack := make([]string, 0)
	dir_sizes := make(map[string]int)
	seen_files := make(map[string]bool)

	dir_sizes["/"] = 0

	for scanner.Scan() {
		line := scanner.Text()

		if string(line[0]) == "$" {
			// fmt.Println(line_start) 

			// when we enter into a directory, we should add it to our dir_stack
			// so that all files it contains can contribute to its size
			if string(line[2:4]) == "cd" {

				if line[5:] == ".." {
					// pop from the dir stack
					//fmt.Println(dir_stack)
					dir_stack = pop(dir_stack)
					//fmt.Println(dir_stack)
				} else {
					//fmt.Println(line)
					var string_to_add string
					if len(dir_stack) > 0 {
						string_to_add = dir_stack[len(dir_stack)-1] + line[5:]
					} else {
						string_to_add = line[5:]
					}
					
					dir_stack = append(dir_stack, string_to_add)
				}
			}

		} else if string(line[:3]) == "dir" {
			// We might need to check if the directory has already been
			// added to the map, so that we dont reset it to zero
			dir_path := dir_stack[0] + string(line[4:])
			if _, ok := dir_sizes[dir_path]; !ok {
				dir_sizes[string(line[4:])] = 0
			} else {
				fmt.Println("Already added:")
				fmt.Println(string(line[4:]))
				fmt.Println()
			}
			
		} else {
			// add the file size to all directories in dir_stack
			size_name := strings.Split(line, " ")

			// if we havent seen the file, count it
			if _, ok := seen_files[line]; !ok {

				size_temp, _ := strconv.Atoi(size_name[0])

				for _, d := range dir_stack {
					dir_sizes[d] += size_temp
				}

				seen_files[line] = true
			}

		}

		fmt.Println(dir_stack)
		//fmt.Println(line, dir_stack[len(dir_stack)-1])
	}
	//fmt.Println(dir_sizes)

	sum := 0
	for _, v := range dir_sizes {
		if v < 100000 {
			//fmt.Println(k, v)
			sum += v
		}
	}

	fmt.Println(sum)


}

func pop(s []string) []string {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    return s[:l-1]
}
