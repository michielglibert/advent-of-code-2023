package main

import (
	utils "advent-of-code"
	"fmt"
)

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

// Contains what direction to go for each character
var dirMap = map[rune]map[Direction]Direction{
	'|': {
		NORTH: SOUTH,
		SOUTH: NORTH,
	},
	'-': {
		EAST: WEST,
		WEST: EAST,
	},
	'L': {
		NORTH: EAST,
		EAST:  NORTH,
	},
	'J': {
		NORTH: WEST,
		WEST:  NORTH,
	},
	'7': {
		SOUTH: WEST,
		WEST:  SOUTH,
	},
	'F': {
		EAST:  SOUTH,
		SOUTH: EAST,
	},
}

var revMap = map[Direction]Direction{
	NORTH: SOUTH,
	SOUTH: NORTH,
	WEST:  EAST,
	EAST:  WEST,
}

type Network [][]rune

// Get a 2D array of the network
func getNetwork(lines []string) Network {
	var network Network

	for _, line := range lines {
		lineArray := []rune(line)
		network = append(network, lineArray)
	}

	return network
}

// Loop over network for finding "S"
func getAnimalPos(network Network) (int, int) {
	for i := range network {
		for j := range network[i] {
			if network[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func getNetworkIndex(i, j int, dir Direction, network Network) (int, int) {
	switch dir {
	case NORTH:
		newI := max(i-1, 0)
		return newI, j
	case SOUTH:
		newI := min(i+1, len(network)-1)
		return newI, j
	case EAST:
		newJ := min(j+1, len(network[0])-1)
		return i, newJ
	case WEST:
		newJ := max(j-1, 0)
		return i, newJ
	default:
		fmt.Println("Unknown direction")
		return -1, -1
	}
}

func getStartPos(i, j int, network Network) (int, int, Direction) {
	var startDir Direction
	var newI, newJ int
	directions := []Direction{NORTH, SOUTH, EAST, WEST}

	for _, dir := range directions {
		posI, posJ := getNetworkIndex(i, j, dir, network)
		char := network[posI][posJ]
		_, ok := dirMap[char][revMap[dir]]
		if ok {
			newI = posI
			newJ = posJ
			startDir = revMap[dir]
		}
	}

	return newI, newJ, startDir
}

// Traverse network from starting pos to find S again
func getSteps(i, j int, network Network) int {
	steps := 1
	posI, posJ, dir := getStartPos(i, j, network)

	currentChar := network[posI][posJ]

	for currentChar != 'S' {
		steps++

		nextDir := dirMap[currentChar][dir]

		posI, posJ = getNetworkIndex(posI, posJ, nextDir, network)

		currentChar = network[posI][posJ]

		dir = revMap[nextDir]
	}

	return steps
}

func main() {
	lines := utils.ReadFile("../input")
	network := getNetwork(lines)
	i, j := getAnimalPos(network)
	steps := getSteps(i, j, network)
	fmt.Println(steps / 2)
}
