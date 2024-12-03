package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parseInput processes the input string and splits it into two slices of integers.
// Each line is expected to contain two space-separated integers.
func parseInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	var leftList, rightList []int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 2 {
			// Convert both numbers in the line to integers
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])

			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		}
	}

	return leftList, rightList
}

// calculateDistance computes the total distance between corresponding elements
// of two sorted integer slices. Distance is the absolute difference between values.
func calculateDistance(leftList, rightList []int) int {
	sort.Ints(leftList)  // Sort left list in ascending order
	sort.Ints(rightList) // Sort right list in ascending order

	totalDistance := 0
	for i := range leftList {
		totalDistance += abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

// abs returns the absolute value of the given integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// calculateSimilarityScore determines a "similarity score" by counting occurrences
// of each value in the right list and summing their products with the corresponding
// values in the left list.
func calculateSimilarityScore(leftList, rightList []int) int {
	rightCounts := make(map[int]int) // Map to store counts of each number in rightList
	for _, num := range rightList {
		rightCounts[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCounts[num] // Add weighted contribution
	}

	return similarityScore
}

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	input := string(inputBytes)

	leftList, rightList := parseInput(input)

	distance := calculateDistance(leftList, rightList)
	fmt.Printf("[PART 1] The total distance between the lists is: %d\n", distance)

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("[PART 2] The similarity score is: %d\n", similarityScore)
}
