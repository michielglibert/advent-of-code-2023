package main

import (
	utils "advent-of-code"
	"fmt"
	"regexp"
)

type Race struct {
	time     int
	distance int
}

func raceMiniboat(holdDuration int, allowedTime int) int {
	remainingTime := allowedTime - holdDuration
	return holdDuration * remainingTime
}

func getRaces(lines []string) []Race {
	numRegex := regexp.MustCompile(`\d+`)
	timeMatches := numRegex.FindAllString(lines[0], -1)
	distanceMatches := numRegex.FindAllString(lines[1], -1)

	races := make([]Race, len(timeMatches))

	for i, timeStr := range timeMatches {
		time := utils.StrToInt(timeStr)
		distance := utils.StrToInt(distanceMatches[i])

		race := Race{
			time:     time,
			distance: distance,
		}
		races[i] = race
	}

	return races
}

func main() {
	lines := utils.ReadFile("../input")
	races := getRaces(lines)

	possibleWays := 1

	for _, race := range races {
		beatenTimes := make([]int, 0)
		for i := 1; i < race.time-1; i++ {
			distance := raceMiniboat(i, race.time)
			if distance > race.distance {
				beatenTimes = append(beatenTimes, i)
			}
		}
		possibleWays *= len(beatenTimes)
	}

	fmt.Println(possibleWays)
}
