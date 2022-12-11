package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

type monkey struct {
	items []int64
	op func(int64) int64
	test func(int64) (int, bool)
	inspects int
	div_by int
}

var common int64

func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/11/input11.txt")

	monkeys := make([]monkey, 0)

	// split file on empty lines and loop over each block
	common = 1

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
					monkeys[i].items = append(monkeys[i].items, int64(v))
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
		common *= int64(div_by)
		monkeys[i].div_by = div_by
		monkeys[i].test = func(n int64) (int, bool) { 
			if n % int64(div_by) == 0 { 
				return true_to, true
				} else { 
					return false_to, false
				}
			}
		
		monkeys[i].op = get_op(op)

	}
	//fmt.Println(monkeys)

	for j := 1; j < 10001; j++ {

		for i := range monkeys {
			//fmt.Println(monkeys[i].items)
			for _, item := range monkeys[i].items {
				item_to_throw := monkeys[i].op(item)

				catcher, _ := monkeys[i].test(item_to_throw)

				item_to_throw = item_to_throw % common
	
				//fmt.Printf("Items of catcher before throw: %v\n", monkeys[catcher].items)
				monkeys[catcher].items = append(monkeys[catcher].items, item_to_throw)
				//fmt.Println(monkeys[catcher].items)
				monkeys[i].items = pop(monkeys[i].items)

				monkeys[i].inspects += 1
	
			}
			// fmt.Printf("After turn %v\n", i)
			// fmt.Println(monkeys)
			// fmt.Println()
		}
	
		if j % 1000 == 0 || j == 1 || j == 20 {
			fmt.Printf("Round %v\n", j)
			for _, m := range monkeys {
				
				fmt.Println(m.inspects)
				
			}
			fmt.Println()
		}

	}
	ins := make([]int, 0)
	for _, m := range monkeys {
		ins = append(ins, m.inspects)
	}

	sort.Ints(ins)
	fmt.Println(ins)

	fmt.Println(ins[len(ins)-1] * ins[len(ins)-2])

}

func get_op(op string) func(int64) int64 {
	if string(op[4]) == "*" {
		v := op[6:]
		if v == "old" {
			return func(a int64) int64 {return a * a}
		} else {
			vi, _ := strconv.Atoi(v)
			return func(a int64) int64 {return a * int64(vi)}
		}
	} else if string(op[4]) == "+" {
		v := op[6:]
		if v == "old" {
			return func(a int64) int64 {return a + a}
		} else {
			vi, _ := strconv.Atoi(v)
			return func(a int64) int64 {return a + int64(vi)}
		}
	} 

	return nil
}

func pop(s []int64) []int64 {

    l := len(s)
    return s[:l-1]
}