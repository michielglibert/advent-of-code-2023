package main

import (
	utils "advent-of-code"
	"fmt"
	"regexp"
)

func every(elements []int, condition func(int) bool) bool {
	for _, element := range elements {
		if !condition(element) {
			return false
		}
	}
	return true
}

func getDifferences(sequence []int) [][]int {
	sequences := make([][]int, 1)
	sequences[0] = sequence
	currentSequence := sequences[len(sequences)-1]

	for !every(currentSequence, func(elem int) bool {
		return elem == 0
	}) {
		newSequence := make([]int, len(currentSequence)-1)
		for i := 0; i < len(currentSequence)-1; i++ {
			newSequence[i] = currentSequence[i+1] - currentSequence[i]
		}
		sequences = append(sequences, newSequence)
		currentSequence = sequences[len(sequences)-1]
	}

	return sequences
}

func extrapolate(sequences [][]int) int {
	newVal := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		newVal = sequences[i][0] - newVal
	}
	return newVal
}

func getSequence(line string) []int {
	reg := regexp.MustCompile(`(-?\d)+`)
	matches := reg.FindAllString(line, -1)

	sequence := make([]int, len(matches))

	for i, match := range matches {
		sequence[i] = utils.StrToInt(match)
	}

	return sequence
}

func main() {
	lines := utils.ReadFile("../input")

	getSequence(lines[0])

	sum := 0
	for _, line := range lines {
		sequence := getSequence(line)
		differences := getDifferences(sequence)
		extrapolatedValue := extrapolate(differences)
		fmt.Println(extrapolatedValue)
		sum += extrapolatedValue
	}

	fmt.Println(sum)
}
