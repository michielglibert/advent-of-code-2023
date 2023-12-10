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
	re := regexp.MustCompile(`([\dA-Z]+) = \(([\dA-Z]+), ([\dA-Z]+)\)`)

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

func getStartNodes(nodesMap map[string]Node) []string {
	var startNodes []string

	for key := range nodesMap {
		if key[2] == 'A' {
			startNodes = append(startNodes, key)
		}
	}

	return startNodes
}

func checkIfEndIsReached(node string) bool {
	isEndReached := true
	if node[2] != 'Z' {
		isEndReached = false
	}
	return isEndReached
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findLCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}

	return result
}

func getSteps(instructions string, nodesMap map[string]Node) int {
	currentNodes := getStartNodes(nodesMap)
	currentInstructionPos := 0
	stepsPerNode := make([]int, len(currentNodes))

	for i := range currentNodes {
		steps := 0
		for !checkIfEndIsReached(currentNodes[i]) {
			steps++

			currentInstruction := instructions[currentInstructionPos]

			if currentInstruction == 'L' {
				currentNodes[i] = nodesMap[currentNodes[i]].left
			} else {
				currentNodes[i] = nodesMap[currentNodes[i]].right
			}

			if currentInstructionPos == len(instructions)-1 {
				currentInstructionPos = 0
			} else {
				currentInstructionPos++
			}
		}
		stepsPerNode[i] = steps
	}

	return findLCM(stepsPerNode)
}

func main() {
	lines := utils.ReadFile("../input")

	instructions := lines[0]
	nodesMap := getNodesMap(lines[2:])

	steps := getSteps(instructions, nodesMap)

	fmt.Println(steps)

}
