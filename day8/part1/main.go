package main

import (
	utils "advent-of-code"
	"fmt"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func getNodesMap(lines []string) map[string]Node {
	re := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

	nodesMap := make(map[string]Node)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		nodeName := matches[1]
		left := matches[2]
		right := matches[3]

		nodesMap[nodeName] = Node{left: left, right: right}
	}

	return nodesMap
}

func getSteps(instructions string, nodesMap map[string]Node) int {
	currentNode := "AAA"
	currentInstructionPos := 0
	steps := 0

	for currentNode != "ZZZ" {
		steps++

		currentInstruction := instructions[currentInstructionPos]

		if currentInstruction == 'L' {
			currentNode = nodesMap[currentNode].left
		} else {
			currentNode = nodesMap[currentNode].right
		}

		if currentInstructionPos == len(instructions)-1 {
			currentInstructionPos = 0
		} else {
			currentInstructionPos++
		}
	}

	return steps
}

func main() {
	lines := utils.ReadFile("../input")

	instructions := lines[0]
	nodesMap := getNodesMap(lines[2:])

	steps := getSteps(instructions, nodesMap)

	fmt.Println(steps)

}
