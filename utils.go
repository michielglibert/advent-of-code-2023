package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
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

func StrArrToRunes(strings []string) (runes []rune) {
	for _, s := range strings {
		runes = append(runes, rune(s[0]))
	}
	return
}

func StrArrToInt(strings []string) (ints []int) {
	for _, s := range strings {
		ints = append(ints, StrToInt(s))
	}
	return
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error converting seed string to number")
	}
	return num
}

func GetMulti(nums []int) int {
	multi := 1
	for _, num := range nums {
		multi *= num
	}
	return multi
}

func ReverseArray(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
}
