/*
The solution here was to make a 2D array of the schematic and loop over each item and check
the adjacent symbols. Meanwhile we build up a memory of the number we are looping over.
*/

package main

import (
	fileio "advent-of-code"
	"fmt"
	"strconv"
	"strings"
)

func to2DArray(lines []string) [][]string {
	result := make([][]string, len(lines))

	for i, str := range lines {
		chars := strings.Split(str, "")
		result[i] = chars
	}

	return result
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getValueForSchematic(posY int, posX int, schematic [][]string) string {
	if posX >= 0 && posY >= 0 && posY < len(schematic) && posX < len(schematic[posY]) {
		return schematic[posY][posX]
	}
	return "."
}

func checkIsPartNumber(posY int, posX int, schematic [][]string) bool {
	isPartNumber := false

	xLeft := posX - 1
	xRight := posX + 1

	yUp := posY - 1
	yDown := posY + 1

	coordsMap := map[string]string{
		"topLeft":     getValueForSchematic(yUp, xLeft, schematic),
		"top":         getValueForSchematic(yUp, posX, schematic),
		"topRight":    getValueForSchematic(yUp, xRight, schematic),
		"left":        getValueForSchematic(posY, xLeft, schematic),
		"right":       getValueForSchematic(posY, xRight, schematic),
		"bottomLeft":  getValueForSchematic(yDown, xLeft, schematic),
		"bottom":      getValueForSchematic(yDown, posX, schematic),
		"bottomRight": getValueForSchematic(yDown, xRight, schematic),
	}

	for _, value := range coordsMap {
		if !isNumber(value) && value != "." {
			isPartNumber = true
		}
	}

	return isPartNumber
}

func checkIsNextNumber(posY int, posX int, schematic [][]string) bool {
	nextValue := getValueForSchematic(posY, posX+1, schematic)
	return isNumber(nextValue)
}

func main() {
	lines := fileio.ReadFile("../input")
	schematic := to2DArray(lines)

	sum := 0

	for i, line := range schematic {
		isPartNumber := false
		currentNum := ""

		for j := range line {
			posContent := schematic[i][j]

			if isNumber(posContent) {
				currentNum += schematic[i][j]
				if !isPartNumber {
					isPartNumber = checkIsPartNumber(i, j, schematic)
				}

				if !checkIsNextNumber(i, j, schematic) {
					if isPartNumber {
						totalNumber, err := strconv.Atoi(currentNum)
						if err != nil {
							fmt.Println("Issue converting currentNum to number")
							return
						}
						sum += totalNumber
					}
					isPartNumber = false
					currentNum = ""
				}
			}
		}

	}

	fmt.Println(sum)
}
