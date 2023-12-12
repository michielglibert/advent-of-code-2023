package main

import (
	utils "advent-of-code"
	"regexp"
	"strings"
)

type Record struct {
	springs          []rune
	duplicatedFormat []int
}

func getRecords(lines []string) []Record {
	reg := regexp.MustCompile(`([\?\.#]+)\s([,\d]+)`)
	records := make([]Record, len(lines))

	for i, line := range lines {
		matches := reg.FindAllString(line, -1)
		springsStr := strings.Split(matches[1], "")
		duplicatedFormatStr := strings.Split(matches[2], ",")

		springs := utils.StrArrToRunes(springsStr)
		duplicatedFormat := utils.StrArrToInt(duplicatedFormatStr)

		records[i] = Record{springs: springs, duplicatedFormat: duplicatedFormat}
	}

	return records
}

func main() {
	input := utils.ReadFile("../testinput")
}
