package main

import (
	utils "advent-of-code"
	"fmt"
	"math"
	"os"
)

func galaxyToString(galaxy Galaxy) string {
	var result string

	for _, row := range galaxy {
		for _, cell := range row {
			if cell == 0 {
				result += "."
			} else {
				result += string(cell)
			}
		}
		result += "\n"
	}

	return result
}

func saveGalaxyToFile(galaxy Galaxy, filename string) error {
	// Convert the network to a human-readable string
	outputString := galaxyToString(galaxy)

	// Write the string to the file using os.WriteFile
	err := os.WriteFile(filename, []byte(outputString), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Galaxy saved to %s\n", filename)
	return nil
}

type Galaxy [][]rune

// Get a 2D array of the galaxy
func getGalaxy(lines []string) Galaxy {
	var galaxy Galaxy

	for _, line := range lines {
		lineArray := []rune(line)
		galaxy = append(galaxy, lineArray)
	}

	return galaxy
}

func insertRow(galaxy Galaxy, position, numRows int, newElement rune) Galaxy {
	// Ensure position is within bounds
	if position < 0 || position > len(galaxy) {
		fmt.Println("Invalid position for inserting row")
		return galaxy
	}

	// Insert multiple rows at the specified position with the new element
	for i := 0; i < numRows; i++ {
		newRow := make([]rune, len(galaxy[0]))
		for j := range newRow {
			newRow[j] = newElement
		}
		galaxy = append(galaxy, nil)
		copy(galaxy[position+1:], galaxy[position:])
		galaxy[position] = newRow
		position++
	}

	return galaxy
}

func insertColumn(galaxy Galaxy, position, numCols int, newElement rune) Galaxy {
	// Ensure position is within bounds
	if position < 0 || position > len(galaxy[0]) {
		fmt.Println("Invalid position for inserting column")
		return galaxy
	}

	// Insert multiple columns at the specified position with the new element
	for i := range galaxy {
		// Ensure each row has enough space for the new columns
		galaxy[i] = append(galaxy[i], make([]rune, numCols)...)
		copy(galaxy[i][position+numCols:], galaxy[i][position:])
		for j := 0; j < numCols; j++ {
			galaxy[i][position+j] = newElement
		}
	}

	return galaxy
}

// Returns rows first (i) and cols next (j)
func getEmptyRowsAndCols(galaxy Galaxy) ([]int, []int) {
	filledRows := make([]bool, len(galaxy))
	filledCols := make([]bool, len(galaxy[0]))

	for i := range galaxy {
		for j, val := range galaxy[i] {
			if val != '.' {
				filledRows[i] = true
				filledCols[j] = true
			}
		}
	}

	var emptyRows []int
	for i, val := range filledRows {
		if !val {
			emptyRows = append(emptyRows, i)
		}
	}

	var emptyCols []int
	for i, val := range filledCols {
		if !val {
			emptyCols = append(emptyCols, i)
		}
	}

	return emptyRows, emptyCols
}

func expandGalaxy(emptyRows, emptyCols []int, galaxy Galaxy, amount int) Galaxy {
	// Create a new slice with the same length as galaxy
	newGalaxy := make(Galaxy, len(galaxy))
	for i := range newGalaxy {
		newGalaxy[i] = make([]rune, len(galaxy[i]))
		copy(newGalaxy[i], galaxy[i])
	}

	extraRows := 0
	for i, val := range emptyRows {
		newGalaxy = insertRow(newGalaxy, val, amount, '.')
		extraRows += amount

		if len(emptyRows) > i+1 {
			emptyRows[i+1] += extraRows
		}
	}

	extraCols := 0
	for i, val := range emptyCols {
		newGalaxy = insertColumn(newGalaxy, val, amount, '.')
		extraCols += amount

		if len(emptyCols) > i+1 {
			emptyCols[i+1] += extraCols
		}
	}

	return newGalaxy
}

type Coord struct {
	x int
	y int
}

func getGalaxyCoords(galaxy Galaxy) []Coord {
	var galaxyCoords []Coord
	for y, line := range galaxy {
		for x, val := range line {
			if val == '#' {
				galaxyCoords = append(galaxyCoords, Coord{x: x, y: y})
			}
		}
	}
	return galaxyCoords
}

type GalaxyPair struct {
	one Coord
	two Coord
}

func getAllShortestPathsSum(coords []Coord) int {
	resultMap := make(map[GalaxyPair]int)

	for _, coordOne := range coords {
		for _, coordTwo := range coords {
			if coordOne != coordTwo {
				pair := GalaxyPair{one: coordOne, two: coordTwo}
				otherPair := GalaxyPair{one: coordTwo, two: coordOne}
				_, ok := resultMap[pair]
				_, ok2 := resultMap[otherPair]
				if !ok && !ok2 {
					distance := int(math.Abs(float64(coordOne.x-coordTwo.x))) + int(math.Abs(float64(coordOne.y-coordTwo.y)))
					resultMap[pair] = distance
				}
			}
		}
	}

	sum := 0
	for _, val := range resultMap {
		sum += val
	}
	return sum
}

func doLinearInterpol(y0 int, y1 int, x0 int, x1 int, x int) int {
	return (y0 + (x-x0)*((y1-y0)/(x1-x0)))
}

func main() {
	lines := utils.ReadFile("../input")

	galaxy1 := getGalaxy(lines)
	emptyRows, emptyCols := getEmptyRowsAndCols(galaxy1)
	galaxy1 = expandGalaxy(emptyRows, emptyCols, galaxy1, 1)
	galaxyCoords1 := getGalaxyCoords(galaxy1)
	sum1 := getAllShortestPathsSum(galaxyCoords1)
	fmt.Println(sum1)

	galaxy2 := getGalaxy(lines)
	emptyRows2, emptyCols2 := getEmptyRowsAndCols(galaxy2)
	galaxy2 = expandGalaxy(emptyRows2, emptyCols2, galaxy2, 2)
	galaxyCoords2 := getGalaxyCoords(galaxy2)
	sum2 := getAllShortestPathsSum(galaxyCoords2)
	fmt.Println(sum2)

	// Linear interpol
	VALUE := 1000000
	fmt.Println(doLinearInterpol(sum1, sum2, 1, 2, VALUE-1))
}
