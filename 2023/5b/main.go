package main

import (
	"fmt"
	"aoc/utils"
	"regexp"
	"strings"
	"strconv"
	// "math"
	"sort"
)

func main() {
	lines := utils.GetLines("input.txt")

	re_nbr := regexp.MustCompile(`\d+`)
	seeds := utils.ToInt(re_nbr.FindAllString(lines[0], -1))

	seed_ranges := getSeedRanges(seeds)
	fmt.Println("Seed ranges:", seed_ranges)

	// starts at 1!
	maps := parseInput(lines[1:])
	fmt.Println("Maps:", maps) // (dest_start, src_start, range)

	resulting_intervals := make([][]int, 0)

	var current_intervals [][]int
	for _, seed_interval := range seed_ranges {
		// create a 2d array with only seed_interval as the array
		current_intervals = [][]int{seed_interval}
		new_intervals := make([][]int, 0)

		// loop over all maps
		for i := 1; i <= len(maps); i++ {
			for _, interval := range current_intervals {
				// do all submappings of the map to each of our intervals

				// remap_chan := remap(interval[0], interval[1], maps[i])
				// for r := range remap_chan {
				// 	new_intervals = append(new_intervals, r[:])
				// 	fmt.Println("New interval:", new_intervals)
				// }

				mapped := remap_new(interval[0], interval[1], maps[i])
				for _, r := range mapped {
					new_intervals = append(new_intervals, r[:])
					fmt.Println("New interval:", new_intervals)
				}

			}

			current_intervals = new_intervals
			new_intervals = make([][]int, 0)
		}

		resulting_intervals = append(resulting_intervals, current_intervals...)

		fmt.Println("Resulting intervals:", resulting_intervals)
	}

	// sort resulting intervals by the first value in each array
	sort.Slice(resulting_intervals, func(i, j int) bool {
		return resulting_intervals[i][0] < resulting_intervals[j][0]
	})
	fmt.Println("Min value:", resulting_intervals[0][0])

}

// wrong
func remap(lo, hi int, m [][]int) <-chan [2]int {
	resultChan := make(chan [2]int)

	go func() {
		defer close(resultChan)

		var ans [][3]int

		for _, interval := range m {
			dst, src, R := interval[0], interval[1], interval[2]
			end := src + R - 1
			D := dst - src

			if !(end < lo || src > hi) {
				ans = append(ans, [3]int{utils.MaxInt(src, lo), utils.MinInt(end, hi), D})
			}
		}

		for i, interval := range ans {
			l, r, D := interval[0], interval[1], interval[2]
			resultChan <- [2]int{l + D, r + D}

			if i < len(ans)-1 && ans[i+1][0] > r+1 {
				resultChan <- [2]int{r + 1, ans[i+1][0] - 1}
			}
		}

		if len(ans) == 0 {
			resultChan <- [2]int{lo, hi}
			return
		}

		if ans[0][0] != lo {
			resultChan <- [2]int{lo, ans[0][0] - 1}
		}
		if ans[len(ans)-1][1] != hi {
			resultChan <- [2]int{ans[len(ans)-1][1] + 1, hi}
		}
	}()

	return resultChan
}

func getSeedRanges(seeds []int) [][]int {
	// [start, stop)
	res := make([][]int, 0)
	for i := 0; i < len(seeds); i += 2 {
		res = append(res, []int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	return res
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

	// for each key, sort all arrays
	for k, v := range maps {
		sort.Slice(v, func(i, j int) bool {
			return v[i][1] < v[j][1]
		})
		maps[k] = v
	}

    return maps
}

func remap_new(lo, hi int, m [][]int) [][]int {
	var ans [][]int

	for _, interval := range m {
		dst := interval[0]
		src := interval[1]
		R := interval[2]
		end := src + R - 1
		D := dst - src

		if !(end < lo || src > hi) {
			ans = append(ans, []int{utils.MaxInt(src, lo), utils.MinInt(end, hi), D})
		}
	}

	var result [][]int

	for i, interval := range ans {
		l := interval[0]
		r := interval[1]
		D := interval[2]

		result = append(result, []int{l + D, r + D})

		if i < len(ans)-1 && ans[i+1][0] > r+1 {
			result = append(result, []int{r + 1, ans[i+1][0] - 1})
		}
	}

	// End and start ranges
	if len(ans) == 0 {
		return [][]int{{lo, hi}}
	}

	if ans[0][0] != lo {
		result = append([][]int{{lo, ans[0][0] - 1}}, result...)
	}
	if ans[len(ans)-1][1] != hi {
		result = append(result, []int{ans[len(ans)-1][1] + 1, hi})
	}

	return result
}
