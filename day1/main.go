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
	if rev == true {
		return reverseString(num)
	} else {
		return num
	}
}

func checkForTextDigit(code string, rev bool) byte {
	var index int

	index = strings.Index(code, getDigitString("one", rev))
	if index == 0 {
		return '1'
	}

	index = strings.Index(code, getDigitString("two", rev))
	if index == 0 {
		return '2'
	}

	index = strings.Index(code, getDigitString("three", rev))
	if index == 0 {
		return '3'
	}

	index = strings.Index(code, getDigitString("four", rev))
	if index == 0 {
		return '4'
	}

	index = strings.Index(code, getDigitString("five", rev))
	if index == 0 {
		return '5'
	}

	index = strings.Index(code, getDigitString("six", rev))
	if index == 0 {
		return '6'
	}

	index = strings.Index(code, getDigitString("seven", rev))
	if index == 0 {
		return '7'
	}

	index = strings.Index(code, getDigitString("eight", rev))
	if index == 0 {
		return '8'
	}

	index = strings.Index(code, getDigitString("nine", rev))
	if index == 0 {
		return '9'
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

		for i := 0; i < len(value); i++ {
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
