package main

import (
	utils "advent-of-code"
	"fmt"
	"strings"
)

type Record struct {
	springs string
	groups  []int
}

func getRecords(lines []string) []Record {
	records := make([]Record, len(lines))

	for i, line := range lines {
		matches := strings.Split(line, " ")
		springs := matches[0]

		groupsStr := strings.Split(matches[1], ",")
		groups := utils.StrArrToInt(groupsStr)

		records[i] = Record{springs: springs, groups: groups}
	}

	return records
}

func noDot(str string) bool {
	hasNoDots := true
	for _, val := range str {
		if val == '.' {
			hasNoDots = false
		}
	}
	return hasNoDots
}

func analyze(spring string, groups []int) int {
	// If spring is empty, groups should be empty too
	if len(spring) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	// If groups is empty, spring should not contain `#`
	if len(groups) == 0 {
		if strings.Contains(spring, "#") {
			return 0
		}
		return 1
	}

	// If character is ., skip it
	if spring[0] == '.' {
		return analyze(spring[1:], groups)
		// If character is question mark, try out . and # and make a sum of it (multiple arrangements are possible)
	} else if spring[0] == '?' {
		return analyze(utils.ReplaceChar(spring, 0, '.'), groups) + analyze(utils.ReplaceChar(spring, 0, '#'), groups)
	} else if spring[0] == '#' {
		if len(spring) < groups[0] || !noDot(spring[:groups[0]]) {
			return 0
		} else if len(spring) == groups[0] {
			if len(groups) == 1 {
				return 1
			}
		} else if spring[groups[0]] == '#' {
			return 0
		} else {
			return analyze(spring[groups[0]+1:], groups[1:])
		}
	}

	return 0
}

func main() {
	input := utils.ReadFile("../input")
	records := getRecords(input)

	sum := 0
	for _, record := range records {
		sum += analyze(record.springs, record.groups)
	}
	fmt.Println(sum)
}
