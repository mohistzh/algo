package main

import "fmt"

/**
	Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

	A straight forward solution using O(mn) space is probably a bad idea.
	A simple improvement uses O(m + n) space, but still not the best solution.
	Could you devise a constant space solution?
**/
func setMatrixZeroes(matrix [][]int) {
	firstRowZero := false
	firstColumnZero := false
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][0] == 0 {
				firstColumnZero = true
			}
			if i == 0 && matrix[i][j] == 0 {
				firstRowZero = true
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if firstColumnZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

}

func printMatrix(matrix [][]int) {
	n := len(matrix)
	w := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(matrix[i][j], "  ")
		}
		fmt.Println()
	}

}

func main() {
	sample1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	sample2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setMatrixZeroes(sample1)
	setMatrixZeroes(sample2)
	printMatrix(sample1)
	fmt.Println()
	printMatrix(sample2)
}
