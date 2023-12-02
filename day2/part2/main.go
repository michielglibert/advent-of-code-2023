package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var LIMITS = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
	lines := readFile()

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
