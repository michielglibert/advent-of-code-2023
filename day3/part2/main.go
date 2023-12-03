/*
The solution here was to continue on part 1 but keep track of the gears position.
Numbers that have the same gears position related to each other.
We make sure to check on the length of 2 and multiply these together.
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

type Position struct {
	x, y int
}

func checkGearNumber(posY int, posX int, schematic [][]string) *Position {
	xLeft := posX - 1
	xRight := posX + 1

	yUp := posY - 1
	yDown := posY + 1

	coordsMap := map[string]Position{
		"topLeft":     {x: xLeft, y: yUp},
		"top":         {x: posX, y: yUp},
		"topRight":    {x: xRight, y: yUp},
		"left":        {x: xLeft, y: posY},
		"right":       {x: xRight, y: posY},
		"bottomLeft":  {x: xLeft, y: yDown},
		"bottom":      {x: posX, y: yDown},
		"bottomRight": {x: xRight, y: yDown},
	}

	for _, value := range coordsMap {
		coordsVal := getValueForSchematic(value.y, value.x, schematic)
		if coordsVal == "*" {
			return &Position{x: value.x, y: value.y}
		}
	}

	return nil
}

func checkIsNextNumber(posY int, posX int, schematic [][]string) bool {
	nextValue := getValueForSchematic(posY, posX+1, schematic)
	return isNumber(nextValue)
}

func main() {
	lines := fileio.ReadFile("../testinput")
	schematic := to2DArray(lines)

	gearMap := make(map[Position][]int)

	for i, line := range schematic {
		var gearPosition *Position
		currentNum := ""

		for j := range line {
			posContent := schematic[i][j]

			if isNumber(posContent) {
				currentNum += schematic[i][j]

				currPos := checkGearNumber(i, j, schematic)
				if currPos != nil {
					gearPosition = currPos
				}

				if !checkIsNextNumber(i, j, schematic) {
					if gearPosition != nil {
						totalNumber, err := strconv.Atoi(currentNum)
						if err != nil {
							fmt.Println("Issue converting currentNum to number")
							return
						}

						if _, ok := gearMap[*gearPosition]; !ok {
							gearMap[*gearPosition] = make([]int, 0)
						}

						gearMap[*gearPosition] = append(gearMap[*gearPosition], totalNumber)
					}
					gearPosition = nil
					currentNum = ""
				}
			}
		}
	}

	sum := 0

	for _, value := range gearMap {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	fmt.Println(sum)
}
