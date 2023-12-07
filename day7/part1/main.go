package main

import (
	utils "advent-of-code"
	"fmt"
	"sort"
	"strings"
)

var cardStrength = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var handStrength = map[string]int{
	"five":    7,
	"four":    6,
	"fh":      5,
	"three":   4,
	"twopair": 3,
	"pair":    2,
	"high":    1,
}

type Hand struct {
	handScore  int
	cardsScore []int
	bid        int
}

func getSortedValues(cardMap map[rune]int) []int {
	values := make([]int, 0)

	for _, value := range cardMap {
		values = append(values, value)
	}

	sort.Ints(values)
	utils.ReverseArray(values)

	return values
}

func getCardsScore(hand string) []int {
	cardsScore := make([]int, len(hand))

	for i, card := range hand {
		cardsScore[i] = cardStrength[card]
	}

	return cardsScore
}

func getHandScore(hand string) int {
	cardMap := make(map[rune]int)

	for _, card := range hand {
		_, ok := cardMap[card]
		if ok {
			cardMap[card]++
		} else {
			cardMap[card] = 1
		}
	}

	sortedValues := getSortedValues(cardMap)

	result := "high"
	if sortedValues[0] == 5 {
		result = "five"
	} else if sortedValues[0] == 4 {
		result = "four"
	} else if sortedValues[0] == 3 && sortedValues[1] == 2 {
		result = "fh"
	} else if sortedValues[0] == 3 {
		result = "three"
	} else if sortedValues[0] == 2 && sortedValues[1] == 2 {
		result = "twopair"
	} else if sortedValues[0] == 2 {
		result = "pair"
	}

	return handStrength[result]
}

func getHands(lines []string) []Hand {
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		lineArr := strings.Split(line, " ")
		handStr := lineArr[0]
		bidStr := lineArr[1]

		hand := Hand{
			handScore:  getHandScore(handStr),
			cardsScore: getCardsScore(handStr),
			bid:        utils.StrToInt(bidStr),
		}

		hands[i] = hand
	}

	return hands
}

func sortHandsByHandScore(hands []Hand) []Hand {
	sortedHands := make([]Hand, len(hands))
	copy(sortedHands, hands)
	sort.Slice(sortedHands, func(i, j int) bool {
		return sortedHands[i].handScore > sortedHands[j].handScore
	})
	return sortedHands
}

func sortHandsByCardsScore(hands []Hand) []Hand {
	handsByHandScore := make(map[int][]Hand)

	for _, hand := range hands {
		key := hand.handScore
		_, ok := handsByHandScore[key]
		if ok {
			handsByHandScore[key] = append(handsByHandScore[key], hand)
		} else {
			handsByHandScore[hand.handScore] = make([]Hand, 0)
			handsByHandScore[key] = append(handsByHandScore[key], hand)
		}
	}

	for key, hands := range handsByHandScore {
		sortedHands := make([]Hand, len(hands))
		copy(sortedHands, hands)
		sort.Slice(sortedHands, func(i, j int) bool {
			res := false
			for pos := range sortedHands[i].cardsScore {
				if sortedHands[i].cardsScore[pos] != sortedHands[j].cardsScore[pos] {
					res = sortedHands[i].cardsScore[pos] > sortedHands[j].cardsScore[pos]
					break
				}
			}
			return res
		})
		handsByHandScore[key] = sortedHands
	}

	sortedHandsByCardArray := make([]Hand, 0)

	for _, values := range handsByHandScore {
		sortedHandsByCardArray = append(sortedHandsByCardArray, values...)
	}

	return sortedHandsByCardArray
}

func getTotalWinnings(hands []Hand) int {
	sum := 0
	for i := 0; i < len(hands); i++ {
		rank := len(hands) - i
		totalWin := hands[i].bid * rank
		sum += totalWin
	}
	return sum
}

// First we wil creater an array of hands
// We will have to compare the card scores for each hand with the same handScore
// Once we have a sorted array we will be able to give the ranks

func main() {
	lines := utils.ReadFile("../input")
	hands := getHands(lines)
	sortedByHandScore := sortHandsByHandScore(hands)
	sortedByCardsScore := sortHandsByCardsScore(sortedByHandScore)
	totalWinnings := getTotalWinnings(sortedByCardsScore)
	fmt.Println(totalWinnings)
}
