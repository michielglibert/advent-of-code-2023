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

func getGameSets(gameString string) []string {
	gameRegex := regexp.MustCompile(`.*:(.*)`)
	gameMatch := gameRegex.FindStringSubmatch(gameString)
	gameSetsString := gameMatch[1]
	gameSets := strings.Split(gameSetsString, ";")
	return gameSets
}

func getGameSetMap(gameStringSet string) map[string]int {
	colorCounts := make(map[string]int)

	for _, pick := range strings.Split(gameStringSet, ",") {
		setRegex := regexp.MustCompile(`(\d+) (\w+)`)
		setMatches := setRegex.FindStringSubmatch(pick)
		colorCount := setMatches[1]
		color := setMatches[2]

		colorCountNumber, err := strconv.Atoi(colorCount)
		if err != nil {
			fmt.Println("Color count could not be converted to number")
			return nil
		}

		colorCounts[color] = colorCountNumber
	}

	return colorCounts

}

func checkMaxMap(maxMap map[string]int, gameSetMap map[string]int) {
	for key, value := range gameSetMap {
		if _, ok := maxMap[key]; !ok {
			maxMap[key] = value
		} else if maxMap[key] < value {
			maxMap[key] = value
		}
	}
}

func main() {
	lines := fileio.ReadFile("../input")

	if lines == nil {
		fmt.Println("Lines is undefined")
	}

	sum := 0
	for _, line := range lines {
		maxMap := make(map[string]int)
		gameSets := getGameSets(line)
		for _, gameSet := range gameSets {
			gameSetMap := getGameSetMap(gameSet)
			checkMaxMap(maxMap, gameSetMap)
		}
		multi := 1
		for _, value := range maxMap {
			multi *= value
		}
		sum += multi
	}

	fmt.Println(sum)
}
