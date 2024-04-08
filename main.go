package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read JSON file
	file, err := ioutil.ReadFile("hard.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON into 2D array of integers
	var graph [][]int
	err = json.Unmarshal(file, &graph)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Find maximum path sum
	maxSum := findMaxPathSum(graph)
	fmt.Println("Maximum path sum:", maxSum)
}

func findMaxPathSum(graph [][]int) int {
	rows := len(graph)
	if rows == 0 {
		return 0
	}

	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, len(graph[i]))
	}

	// Initialize the bottom row of dp
	for j := range graph[rows-1] {
		dp[rows-1][j] = graph[rows-1][j]
	}

	// Traverse the graph bottom-up to calculate maximum path sum
	for i := rows - 2; i >= 0; i-- {
		for j := range graph[i] {
			dp[i][j] = graph[i][j] + max(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
