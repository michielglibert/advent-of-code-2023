package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func reverseString(input string) string {
	runes := []rune(input)

	// Reverse the order of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func getDigitString(num string, rev bool) string {
	if rev {
		return reverseString(num)
	} else {
		return num
	}
}

func checkForTextDigit(code string, rev bool) byte {
	var index int

	digitsMap := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for key, value := range digitsMap {
		index = strings.Index(code, getDigitString(key, rev))
		if index == 0 {
			return value
		}
	}

	return 0
}

func main() {
	lines := readFile()

	if lines == nil {
		fmt.Println("Lines is undefined")
	}

	sum := 0

	for _, value := range lines {
		var firstDigit byte
		var lastDigit byte

		i := 0
		for (firstDigit == 0 || lastDigit == 0) && i < len(value) {
			if firstDigit == 0 {
				if isDigit(value[i]) {
					firstDigit = value[i]
				} else {
					firstDigit = checkForTextDigit(value[i:], false)
				}
			}

			if lastDigit == 0 {
				lastDigitIndex := len(value) - 1 - i
				if isDigit(value[lastDigitIndex]) {
					lastDigit = value[lastDigitIndex]
				} else {
					lastDigit = checkForTextDigit(reverseString(value)[i:], true)
				}
			}
			i++
		}

		addedString := string(firstDigit) + string(lastDigit)

		intValue, err := strconv.Atoi(addedString)

		if err != nil {
			fmt.Println("Issue converting string to number")
		} else {
			sum += intValue
		}

	}

	fmt.Println(sum)
}
