package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func validateDistances(distances map[[2]Point]int) (int, bool) {
	sum := 0
	valid := true

	for key, expectedDistance := range distances {
		pointA := key[0]
		pointB := key[1]

		calculatedDistance := manhattanDistance(pointA, pointB)

		if calculatedDistance != expectedDistance {
			valid = false
			fmt.Printf("Invalid distance for %v: Expected %d, Got %d\n", key, expectedDistance, calculatedDistance)
		}

		sum += calculatedDistance
	}

	return sum, valid
}

func main() {
	distances := map[[2]Point]int{
		{{0, 2}, {0, 11}}:   9,
		{{0, 2}, {1, 6}}:    5,
		{{0, 2}, {5, 11}}:   14,
		{{0, 2}, {8, 5}}:    11,
		{{0, 2}, {10, 10}}:  18,
		{{0, 2}, {12, 7}}:   17,
		{{0, 11}, {5, 11}}:  5,
		{{1, 6}, {0, 11}}:   6,
		{{1, 6}, {5, 11}}:   9,
		{{1, 6}, {10, 10}}:  13,
		{{1, 6}, {12, 7}}:   12,
		{{4, 0}, {0, 2}}:    6,
		{{4, 0}, {0, 11}}:   15,
		{{4, 0}, {1, 6}}:    9,
		{{4, 0}, {5, 11}}:   12,
		{{4, 0}, {8, 5}}:    9,
		{{4, 0}, {10, 1}}:   7,
		{{4, 0}, {10, 10}}:  16,
		{{4, 0}, {12, 7}}:   15,
		{{8, 5}, {0, 11}}:   14,
		{{8, 5}, {1, 6}}:    8,
		{{8, 5}, {5, 11}}:   9,
		{{8, 5}, {10, 10}}:  7,
		{{8, 5}, {12, 7}}:   6,
		{{10, 1}, {0, 2}}:   11,
		{{10, 1}, {0, 11}}:  20,
		{{10, 1}, {1, 6}}:   14,
		{{10, 1}, {5, 11}}:  15,
		{{10, 1}, {8, 5}}:   6,
		{{10, 1}, {10, 10}}: 9,
		{{10, 1}, {12, 7}}:  8,
		{{10, 10}, {0, 11}}: 11,
		{{10, 10}, {5, 11}}: 6,
		{{12, 7}, {0, 11}}:  16,
		{{12, 7}, {5, 11}}:  11,
		{{12, 7}, {10, 10}}: 5,
	}

	sum, valid := validateDistances(distances)

	if valid {
		fmt.Printf("All Manhattan distances are correct. Sum of distances: %d\n", sum)
	} else {
		fmt.Println("Some distances are incorrect.")
	}
}
