package main

import (
	utils "advent-of-code"
	"container/list"
	"fmt"
	"os"
)

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

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

func networkToString(network Network) string {
	var result string

	for _, row := range network {
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

func saveNetworkToFile(network Network, filename string) error {
	// Convert the network to a human-readable string
	outputString := networkToString(network)

	// Write the string to the file using os.WriteFile
	err := os.WriteFile(filename, []byte(outputString), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Network saved to %s\n", filename)
	return nil
}

func mapNetworkTo3By3(network Network) Network {
	newNetwork := make(Network, len(network)*3)

	for i := 0; i < len(newNetwork); i++ {
		newNetwork[i] = make([]rune, len(network[0])*3)
	}

	for i, networkLine := range network {
		iPos := i * 3
		for j, node := range networkLine {
			jPos := j * 3

			switch node {
			case '|':
				newNetwork[iPos][jPos+1] = '|'
				newNetwork[iPos+1][jPos+1] = '|'
				newNetwork[iPos+2][jPos+1] = '|'
			case '-':
				newNetwork[iPos+1][jPos] = '-'
				newNetwork[iPos+1][jPos+1] = '-'
				newNetwork[iPos+1][jPos+2] = '-'
			case 'L':
				newNetwork[iPos][jPos+1] = '|'
				newNetwork[iPos+1][jPos+1] = 'L'
				newNetwork[iPos+1][jPos+2] = '-'
			case 'J':
				newNetwork[iPos][jPos+1] = '|'
				newNetwork[iPos+1][jPos+1] = 'J'
				newNetwork[iPos+1][jPos] = '-'
			case '7':
				newNetwork[iPos+2][jPos+1] = '|'
				newNetwork[iPos+1][jPos+1] = '7'
				newNetwork[iPos+1][jPos] = '-'
			case 'F':
				newNetwork[iPos+2][jPos+1] = '|'
				newNetwork[iPos+1][jPos+1] = 'F'
				newNetwork[iPos+1][jPos+2] = '-'
			}
		}
	}

	return newNetwork
}

type Tile struct {
	i int
	j int
}

// BFS gives back empty array if edge was found
func doBfs(i, j int, network Network) map[Tile]bool {
	directions := []Direction{NORTH, SOUTH, WEST, EAST}
	startNode := Tile{i: i, j: j}
	visited := make(map[Tile]bool)
	queue := list.New()

	if network[startNode.i][startNode.j] == 0 {
		queue.PushBack(startNode)
		visited[startNode] = true
	}

	for queue.Len() > 0 {
		currentNode := queue.Front()
		queue.Remove(currentNode)

		currentNodeI := currentNode.Value.(Tile).i
		currentNodeJ := currentNode.Value.(Tile).j

		if currentNodeI == 0 || currentNodeJ == 0 || currentNodeI == len(network)-1 || currentNodeJ == len(network[0])-1 {
			return map[Tile]bool{}
		}

		for _, direction := range directions {
			posI, posJ := getNetworkIndex(currentNodeI, currentNodeJ, direction, network)
			neighbor := Tile{i: posI, j: posJ}
			if !visited[neighbor] && network[neighbor.i][neighbor.j] == 0 {
				queue.PushBack(neighbor)
				visited[neighbor] = true
			}
		}
	}

	return visited
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

func getTiles(network Network) int {
	enclosedTiles := make(map[Tile]bool)

	for i := range network {
		for j := range network[i] {
			if !enclosedTiles[Tile{i, j}] {
				resultMap := doBfs(i, j, network)
				for key := range resultMap {
					enclosedTiles[key] = true
				}
			}
		}
	}

	actualEnclosedTiles := make(map[Tile]bool)

	// Since this was mapped 3x3 we should only keep the middle tiles per 3x3
	for key := range enclosedTiles {
		if key.i%3 == 1 && key.j%3 == 1 {
			actualEnclosedTiles[key] = true
		}
	}

	return len(actualEnclosedTiles)
}

// Was not able to solve this one
func main() {
	lines := utils.ReadFile("../testinput3")
	network := getNetwork(lines)
	network3x := mapNetworkTo3By3(network)
	saveNetworkToFile(network3x, "testoutput")
	tiles := getTiles(network3x)
	fmt.Println(tiles)
}
