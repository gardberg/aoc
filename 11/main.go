package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

type monkey struct {
	items []int
	op func(int) int
	test func(int) int
	inspects int
}

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/11/input11.txt")

	monkeys := make([]monkey, 0)

	// split file on empty lines and loop over each block

	for i, block := range strings.Split(strings.TrimSpace(string(file)), "\n\n") {
		//fmt.Println(row)
		//fmt.Println(i)
		monkeys = append(monkeys, monkey{})
		//fmt.Println(block)

		var div_by int
		var true_to int
		var false_to int
		var op string

		for _, row := range strings.Split(block, "\n") {
			s := strings.TrimSpace(row)

			if strings.HasPrefix(s, "Starting") {
				for _, j := range strings.Split(s[16:], ", ") {
					v, _ := strconv.Atoi(j)
					monkeys[i].items = append(monkeys[i].items, v)
				}
			}

			if strings.HasPrefix(s, "Operation") {
				op = s[17:]
			}

			if strings.HasPrefix(s, "Test") {
				div_by, _ = strconv.Atoi(s[19:])
				
			}

			if strings.HasPrefix(s, "If true") {
				true_to, _ = strconv.Atoi(s[25:])
			}

			if strings.HasPrefix(s, "If false") {
				false_to, _ = strconv.Atoi(s[26:])
				//fmt.Println(false_to)
			}
		}
		// fmt.Printf("Monkey %v, div by %v, %v, %v\n", i, div_by, true_to, false_to)
		monkeys[i].test = func(n int) int { if n % div_by == 0 { return true_to } else { return false_to }}
		monkeys[i].op = get_op(op)

	}
	fmt.Println(monkeys)

	for j := 0; j < 20; j++ {

		for i := range monkeys {
			fmt.Println(monkeys[i].items)
			for _, item := range monkeys[i].items {
				item_to_throw := monkeys[i].op(item)
				item_to_throw = int(item_to_throw / 3)
				catcher := monkeys[i].test(item_to_throw)
	
				//fmt.Printf("Items of catcher before throw: %v\n", monkeys[catcher].items)
				monkeys[catcher].items = append(monkeys[catcher].items, item_to_throw)
				//fmt.Println(monkeys[catcher].items)
				monkeys[i].items = pop(monkeys[i].items)

				monkeys[i].inspects += 1
	
			}
			fmt.Printf("After turn %v\n", i)
			fmt.Println(monkeys)
			fmt.Println()
		}
	
		fmt.Println(monkeys)

	}
	ins := make([]int, 0)
	for _, m := range monkeys {
		ins = append(ins, m.inspects)
	}

	sort.Ints(ins)

	fmt.Println(ins[len(ins)-1] * ins[len(ins)-2])

}

func get_op(op string) func(int) int {
	if string(op[4]) == "*" {
		v := op[6:]
		if v == "old" {
			return func(a int) int {return a * a}
		} else {
			vi, _ := strconv.Atoi(v)
			return func(a int) int {return a * vi}
		}
	} else if string(op[4]) == "+" {
		v := op[6:]
		if v == "old" {
			return func(a int) int {return a + a}
		} else {
			vi, _ := strconv.Atoi(v)
			return func(a int) int {return a + vi}
		}
	} 

	return nil
}

func pop(s []int) []int {

    l := len(s)
    return s[:l-1]
}