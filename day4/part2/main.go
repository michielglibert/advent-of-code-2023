package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func readFile() []string {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return lines
}

func strArrToNumsArr(strs []string) []int {
	nums := make([]int, 0)
	for _, numString := range strs {
		num, err := strconv.Atoi(numString)
		if err != nil {
			fmt.Println("Error string converting to int")
			return nil
		}
		nums = append(nums, num)
	}
	return nums
}

type ScratchCard struct {
	gameNumber     int
	winningNumbers []int
	chosenNumbers  []int
}

// Gets a type representation of the scratchcard
func getScratchCard(line string) *ScratchCard {
	// Split between game number and actual card numbers
	lineSplitString := strings.Split(line, ":")
	gameString := lineSplitString[0]
	numbersString := lineSplitString[1]

	// Split between winning and chosen numbers
	numbersStringArr := strings.Split(numbersString, "|")
	winningNumbersString := numbersStringArr[0]
	chosenNumbersString := numbersStringArr[1]

	// Get all the numbers from the strings
	regex := regexp.MustCompile(`\d+`)
	gameNumberString := regex.FindString(gameString)
	winningNumsString := regex.FindAllString(winningNumbersString, -1)
	chosenNumsString := regex.FindAllString(chosenNumbersString, -1)

	// Convert []string to []int
	winningNums := strArrToNumsArr(winningNumsString)
	chosenNums := strArrToNumsArr(chosenNumsString)

	// Convert string to int
	gameNum, err := strconv.Atoi(gameNumberString)
	if err != nil {
		fmt.Println("Error converting gameNumberString to int")
		return nil
	}

	return &ScratchCard{
		gameNumber:     gameNum,
		winningNumbers: winningNums,
		chosenNumbers:  chosenNums,
	}
}

func checkWinningNumbersAmount(card *ScratchCard) int {
	amount := 0
	for _, num := range card.chosenNumbers {
		if slices.Contains(card.winningNumbers, num) {
			amount++
		}
	}
	return amount
}

func main() {
	lines := readFile()

	cards := make(map[int]int)
	sum := 0

	for _, line := range lines {
		sum++
		card := getScratchCard(line)
		amount := checkWinningNumbersAmount(card)

		// If no amount the copies nor the original matter
		if amount > 0 {
			copies, ok := cards[card.gameNumber]
			// Copies
			if ok {
				for i := card.gameNumber + 1; i <= card.gameNumber+amount; i++ {
					_, ok := cards[i]
					if ok {
						cards[i] += copies
					} else {
						cards[i] = copies
					}
					sum += copies
				}
			}
			// Original
			for i := card.gameNumber + 1; i <= card.gameNumber+amount; i++ {
				_, ok := cards[i]
				if ok {
					cards[i]++
				} else {
					cards[i] = 1
				}
				sum++
			}
		}
	}

	fmt.Println(sum)
}
