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
	file, err := os.Open("./input")
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
	winningNumbers []int
	chosenNumbers  []int
}

// Gets a type representation of the scratchcard
func getScratchCard(line string) *ScratchCard {
	// Split between game number and actual card numbers
	numbersString := strings.Split(line, ":")[1]
	numbersStringArr := strings.Split(numbersString, "|")

	// Split between winning and chosen numbers
	winningNumbersString := numbersStringArr[0]
	chosenNumbersString := numbersStringArr[1]

	// Get all the numbers from the strings
	regex := regexp.MustCompile(`\d+`)
	winningNumsString := regex.FindAllString(winningNumbersString, -1)
	chosenNumsString := regex.FindAllString(chosenNumbersString, -1)

	// Convert []string to []int
	winningNums := strArrToNumsArr(winningNumsString)
	chosenNums := strArrToNumsArr(chosenNumsString)

	return &ScratchCard{
		winningNumbers: winningNums,
		chosenNumbers:  chosenNums,
	}
}

func checkWinningNumbers(card *ScratchCard) int {
	points := 0
	for _, num := range card.chosenNumbers {
		if slices.Contains(card.winningNumbers, num) {
			if points == 0 {
				points++
			} else {
				points *= 2
			}
		}
	}
	return points
}

func main() {
	lines := readFile()

	sum := 0
	for _, line := range lines {
		card := getScratchCard(line)
		sum += checkWinningNumbers(card)
	}

	fmt.Println(sum)
}
