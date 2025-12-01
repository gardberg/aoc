package main

import (
	"fmt"
	"aoc/utils"
	"regexp"
	"math"
)

func main() {
	lines := utils.GetLines("input.txt")


	times := parseLine(lines[0]) 
	distances := parseLine(lines[1])

	fmt.Println(times)
	fmt.Println(distances)

	var prod int = 1
	for i, _ := range(times) {
		prod *= calcWinning(times[i], distances[i])
	}

	fmt.Println(prod)

}

func parseLine(line string) []int {
	r := regexp.MustCompile(`\d+`)
	return utils.ToInt(r.FindAllString(line, -1))
}

func calcWinning(time, distance int) int {
	// calc both solutions to x * (time - x) = distance

	s := math.Sqrt(float64(time * time / 4 - distance))
	n1 := float64(time) / 2 + s
	n2 := float64(time) / 2 - s
	fmt.Println(n1, n2)
	return int(math.Ceil(n1) - math.Floor(n2) - 1)


}