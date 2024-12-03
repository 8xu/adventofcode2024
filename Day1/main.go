package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Parse the input into two separate slices
func parseInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	var leftList, rightList []int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 2 {
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])

			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		}
	}

	return leftList, rightList
}

// Function to calculate the total distance
func calculateDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := range leftList {
		totalDistance += abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Function to calculate the similarity score
func calculateSimilarityScore(leftList, rightList []int) int {
	// Count occurrences in the right list
	rightCounts := make(map[int]int)
	for _, num := range rightList {
		rightCounts[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCounts[num]
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
