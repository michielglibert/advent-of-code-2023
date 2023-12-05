package main

import (
	fileio "advent-of-code"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var LIMITS = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameNumber(gameString string) int {
	gameRegex := regexp.MustCompile(`Game (\d+):`)
	gameMatch := gameRegex.FindStringSubmatch(gameString)
	gameNumber, err := strconv.Atoi(gameMatch[1])
	if err != nil {
		fmt.Println("Could not convert gameNumber to int")
		return 0
	}
	return gameNumber
}

func getGameSets(gameString string) []string {
	gameRegex := regexp.MustCompile(`.*:(.*)`)
	gameMatch := gameRegex.FindStringSubmatch(gameString)
	gameSetsString := gameMatch[1]
	gameSets := strings.Split(gameSetsString, ";")
	return gameSets
}

func isSetValid(gameStringSet string) bool {
	colorCounts := make(map[string]int)

	for _, pick := range strings.Split(gameStringSet, ",") {
		setRegex := regexp.MustCompile(`(\d+) (\w+)`)
		setMatches := setRegex.FindStringSubmatch(pick)
		colorCount := setMatches[1]
		color := setMatches[2]

		colorCountNumber, err := strconv.Atoi(colorCount)
		if err != nil {
			fmt.Println("Color count could not be converted to number")
			return false
		}

		colorCounts[color] = colorCountNumber
	}

	for key, value := range LIMITS {
		if colorCounts[key] > value {
			return false
		}
	}
	return true
}

func main() {
	lines := fileio.ReadFile("../input")

	if lines == nil {
		fmt.Println("Lines is undefined")
	}

	sum := 0
	for _, line := range lines {
		var isGamePossible bool = true
		gameNumber := getGameNumber(line)
		gameSets := getGameSets(line)
		for _, gameSet := range gameSets {
			isGamePossible = isSetValid(gameSet)
			if !isGamePossible {
				break
			}
		}
		if isGamePossible {
			sum += gameNumber
		}
	}

	fmt.Println(sum)
}
