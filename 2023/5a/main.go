package main

import (
	"fmt"
	"aoc/utils"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	lines := utils.GetLines("input.txt")

	re_nbr := regexp.MustCompile(`\d+`)
	seeds := utils.ToInt(re_nbr.FindAllString(lines[0], -1))
	//fmt.Println(seeds)

	// starts at 1!
	maps := parseInput(lines[1:])
	//fmt.Println(maps)
	final_seeds := make([]int, 0)
	for _, seed := range seeds {
		// make a copy of seed

		// golang maps not ordered!
		for i := 1; i <= len(maps); i++ {
			seed = doMapping(seed, maps[i])
		}

		final_seeds = append(final_seeds, seed)

	}

	fmt.Println(utils.MinOfSlice(final_seeds))

}

func doMapping(seed int, mapping [][]int) int {
	// mapping: [[dest_range, source_range, range], ...]
	for _, submap := range mapping {
		dr := submap[0]
		sr := submap[1]
		r := submap[2]

		if seed >= sr && seed < sr + r {
			//fmt.Println("Mapping", seed, "to", dr + (seed - sr))
			return dr + (seed - sr)
		}
	}

	//fmt.Println("Mapping", seed, "to", seed)
	return seed
}

func parseInput(lines []string) map[int][][]int {
    maps := make(map[int][][]int)
    var currentMapKey int

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if strings.HasSuffix(line, ":") {
            // Start of a new map
            currentMapKey++
            maps[currentMapKey] = make([][]int, 0)
        } else if line != "" {
            // Parse the numbers in the line
            numbers := strings.Fields(line)
            row := make([]int, len(numbers))
            for i, numStr := range numbers {
                num, _ := strconv.Atoi(numStr)
                row[i] = num
            }
            // Append the row to the current map
            maps[currentMapKey] = append(maps[currentMapKey], row)
        }
    }

    return maps
}
