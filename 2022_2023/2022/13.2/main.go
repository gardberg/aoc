package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"sort"
)

// test case passes, but the real input doesn't
// check so that the recursion we are doing matches the test case

type Packet [][]interface{}

func (s Packet) Len() int {
	return len(s)
}

func (s Packet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Packet) Less(i, j int) bool {
	return compare_packets(s[i], s[j])
}
	
func main() {
	file, _ := os.ReadFile("/Users/lukgar/Desktop/repos/aoc/13/input.txt")

	packets := make([][]interface{}, 0)

	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		if row != "" {
			var par []interface{}

			json.Unmarshal([]byte(row), &par)
			//fmt.Printf("Adding packet: %v\n", par)
			packets = append(packets, par)
		}
 	}

	var div1, div2 []interface{}

	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)

	packets = append(packets, div1)
	packets = append(packets, div2)

	fmt.Printf("Packets len: %v\n", len(packets))

	sort.Sort(Packet(packets))

	j, k := 0, 0
	for i := range packets {
		if compare_arrays(div1, packets[i]) == 0 { 
			j = i 
		} else if compare_arrays(div2, packets[i]) == 0 { k = i }
	}
	// Use sort.Search to find the index of div1
	//i := sort.Search(len(packets), func(i int) bool { return compare_arrays(packets[i], div1) == 0 })
	//j := sort.Search(len(packets), func(i int) bool { return compare_arrays(packets[i], div2) == 0 })

	fmt.Println((j+1) * (k+1))
	//fmt.Printf("i: %v, j: %v", i, j)
	

}

func compare_packets(left, right []interface{}) bool {
	return compare_arrays(left, right) > 0
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