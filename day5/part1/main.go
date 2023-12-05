/*
The solution here was to continue on part 1 but keep track of the gears position.
Numbers that have the same gears position related to each other.
We make sure to check on the length of 2 and multiply these together.
*/

package main

import (
	fileio "advent-of-code"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Transformation struct {
	dstStart    int
	srcStart    int
	rangeLength int
}

type Almanac struct {
	seeds           []int
	transformations [][]Transformation
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error converting seed string to number")
	}
	return num
}

func getMin(arr []int) int {
	min := arr[0]
	for _, curr := range arr[1:] {
		if curr < min {
			min = curr
		}
	}
	return min
}

func buidAlmanac(lines []string) Almanac {
	numberRegex := regexp.MustCompile(`\d+`)
	sectionTitleRegex := regexp.MustCompile(`[a-zA-Z-]+\smap:`)

	// First line contains seeds
	seedsString := numberRegex.FindAllString(lines[0], -1)
	seeds := make([]int, len(seedsString))
	for i, seedString := range seedsString {
		seedNum := strToInt(seedString)
		seeds[i] = seedNum
	}

	almanac := Almanac{seeds: seeds}

	// Other lines contain transformations
	var currTranformations []Transformation
	for _, line := range lines[1:] {
		isNewSection := sectionTitleRegex.MatchString(line)

		if isNewSection && len(currTranformations) > 0 {
			almanac.transformations = append(almanac.transformations, currTranformations)
			currTranformations = make([]Transformation, 0)
		} else {
			numMatches := numberRegex.FindAllString(line, -1)
			if len(numMatches) > 0 {
				// dst is at index 0, src at index 1 and range at index 2
				dst := strToInt(numMatches[0])
				src := strToInt(numMatches[1])
				rng := strToInt(numMatches[2])

				transformation := Transformation{dstStart: dst, srcStart: src, rangeLength: rng}
				currTranformations = append(currTranformations, transformation)
			}
		}
	}
	if len(currTranformations) > 0 {
		almanac.transformations = append(almanac.transformations, currTranformations)
	}

	return almanac
}

func main() {
	lines := fileio.ReadFile("../input")

	almanac := buidAlmanac(lines)
	locations := make([]int, len(almanac.seeds))

	// Looping over each seed
	for i, seed := range almanac.seeds {
		currLoc := seed
		// Looping over reach section
		for _, transformationCollection := range almanac.transformations {
			// Looping over each transformation per section
			for _, transformation := range transformationCollection {
				if currLoc >= transformation.srcStart && currLoc < (transformation.srcStart+transformation.rangeLength) {
					offset := transformation.dstStart - transformation.srcStart
					currLoc += offset
					break
				}
			}
		}
		locations[i] = currLoc
	}

	closestLoc := getMin(locations)

	fmt.Println(closestLoc)

}
