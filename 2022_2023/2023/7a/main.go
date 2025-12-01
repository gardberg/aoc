package main

import (
	"fmt"
	"aoc/utils"
	"strings"
	"sort"
)

var value = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}	

func main() {
	lines := utils.GetLines("input.txt")

	bidMap := make(map[string]int)
	hands := make([]string, 0)

	for _, line := range(lines) {
		split := strings.Split(line, " ")
		hand := split[0]
		hands = append(hands, hand)
		bid := utils.StringToInt(split[1])
		bidMap[hand] = bid
	}

	sort.Slice(hands, func(i, j int) bool {
		return !compareHands(hands[i], hands[j])
	})

	fmt.Println(hands)

	total_winnings := 0
	for i, hand := range(hands) {
		total_winnings += (i+1) * bidMap[hand]
	}

	fmt.Println(total_winnings)


	
}

func compareHands(hand1 string, hand2 string) bool {
	hand1Type := getHandType(hand1)
	hand2Type := getHandType(hand2)

	if hand1Type == hand2Type {
		return breakTie(hand1, hand2)
	} else {
		return hand1Type > hand2Type
	}
}

func breakTie(hand1 string, hand2 string) bool {
	for i := 0; i < len(hand1); i++ {
		if value[string(hand1[i])] == value[string(hand2[i])] {
			continue
		} else {
			return value[string(hand1[i])] > value[string(hand2[i])]
		}
	}

	fmt.Println("ERROR: Tie!")
	return false
}


func getHandType(hand string) int {
	// 7 hand types:
	// five of a kind (6)
	// four of a kind (5)
	// full house (4)
	// three of a kind (3)
	// two pair (2)
	// one pair (1)
	// high card (0)

	count := getHandCount(hand)

	// Check for five of a kind
	if len(count) == 1 {
		return 6
	}

	// Check for four of a kind
	for _, v := range(count) {
		if v == 4 {
			return 5
		}
	}

	// Check for full house
	if len(count) == 2 {
		return 4
	}

	// Check for three of a kind
	for _, v := range(count) {
		if v == 3 {
			return 3
		}
	}

	// Check for two pair
	if len(count) == 3 {
		return 2
	}

	// Check for one pair
	for _, v := range(count) {
		if v == 2 {
			return 1
		}
	}

	// High card
	return 0

}

func getHandCount(hand string) map[string]int {
	// Count the number of each card in a hand
	// Return a map of the counts
	counts := make(map[string]int)
	for _, card := range(hand) {
		counts[string(card)]++
	}
	return counts
}