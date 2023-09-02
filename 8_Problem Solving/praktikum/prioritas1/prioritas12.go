package main

import "fmt"

func generatePascalTriangle(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)
	triangle[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		prevRow := triangle[i-1]
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = prevRow[j-1] + prevRow[j]
		}

		triangle[i] = row
	}

	return triangle
}

func main() {
	numRows := 5
	result := generatePascalTriangle(numRows)
	for _, row := range result {
		fmt.Println(row)
	}
}
