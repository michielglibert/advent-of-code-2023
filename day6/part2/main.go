package main

import (
	utils "advent-of-code"
	"fmt"
	"regexp"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func raceMiniboat(holdDuration int, allowedTime int) int {
	remainingTime := allowedTime - holdDuration
	return holdDuration * remainingTime
}

func getRace(lines []string) Race {
	numRegex := regexp.MustCompile(`\d+`)
	timeMatches := numRegex.FindAllString(lines[0], -1)
	distanceMatches := numRegex.FindAllString(lines[1], -1)

	timeStr := strings.Join(timeMatches, "")
	distanceStr := strings.Join(distanceMatches, "")

	time := utils.StrToInt(timeStr)
	distance := utils.StrToInt(distanceStr)

	race := Race{
		time:     time,
		distance: distance,
	}

	return race
}

func main() {
	lines := utils.ReadFile("../input")
	race := getRace(lines)

	possibleWays := 1

	beatenTimes := make([]int, 0)
	for i := 1; i < race.time-1; i++ {
		distance := raceMiniboat(i, race.time)
		if distance > race.distance {
			beatenTimes = append(beatenTimes, i)
		}
	}
	possibleWays *= len(beatenTimes)

	fmt.Println(possibleWays)
}
