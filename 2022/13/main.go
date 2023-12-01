package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

// test case passes, but the real input doesn't
// check so that the recursion we are doing matches the test case
func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/13/input.txt")

	sum := 0

	for i, row := range strings.Split(strings.TrimSpace(string(file)), "\n\n") {
		//fmt.Println(row)
		var left_arr, right_arr []interface{}
		strings := strings.Split(row, "\n")
		left, right := strings[0], strings[1]

		fmt.Println()
		fmt.Printf("Pair %v\n", i+1)
		json.Unmarshal([]byte(left), &left_arr)
		json.Unmarshal([]byte(right), &right_arr)

		o := compare_arrays(left_arr, right_arr)
		fmt.Printf("left: %s, right: %s, in order: %v\n", left, right, o > 0)

		if o > 0 {
			sum += i+1
		}
 	}

	fmt.Println(sum)

}

func compare_arrays(left, right []interface{}) int {
	for i := range left {
		
		if i >= len(right) {
			return -1
		}

		switch type_left := left[i].(type) {
		// if the left element is a list

		case []interface{}:
			// check if the right element is a list
			type_right, ok := right[i].([]interface{})
			if ok {
				if v := compare_arrays(type_left, type_right); v != 0 {
					return v
				}
			} else {
				// then right is a value, put it into a list and compare
				tmp := []interface{}{right[i]}
				if v := compare_arrays(type_left, tmp); v != 0 {
					return v
				}
			}

		default: // left is a value
			// check if right is a list
			type_right, ok := right[i].([]interface{})
			if ok {
				if v := compare_arrays([]interface{}{type_left}, type_right); v != 0 {
					return v
				}
			} else {
				// then right is a value, compare them
				if left[i].(float64) < right[i].(float64) {
					return 1
				} else if left[i].(float64) > right[i].(float64) {
					return -1
				}
			}
		}
	}
	// if right still has items in it
	if len(right) > len(left) {
		return 1
	}

	// both lists equally long, and all elements are equal
	return 0
}