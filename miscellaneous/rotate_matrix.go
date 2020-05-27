package main

import (
	"fmt"
)

/*
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

direction: 1. clockwise direction.   2. anti-clockwise direction

formula: matrix[i][j] = matrix[n-1-j][i]
*/
func rotateMatrix(matrix [][]int, direction int) {
	N := len(matrix)
	if direction == 1 {
		rotateMatrixClockWise(matrix, N)
	} else if direction == 2 {
		rotateMatrixAntiClockWise(matrix, N)
	}
}

func rotateMatrixClockWise(matrix [][]int, N int) {
	for x := 0; x < N/2; x++ {
		for y := x; y < N-x-1; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[N-1-y][x]
			matrix[N-1-y][x] = matrix[N-1-x][N-1-y]
			matrix[N-1-x][N-1-y] = matrix[y][N-1-x]
			matrix[y][N-1-x] = temp
		}
	}
}

func rotateMatrixAntiClockWise(matrix [][]int, N int) {
	for x := 0; x < N/2; x++ {
		for y := x; y < N-x-1; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[y][N-1-x]
			matrix[y][N-1-x] = matrix[N-1-x][N-1-x]
			matrix[N-1-x][N-1-y] = matrix[N-1-y][x]
			matrix[N-1-y][x] = temp
		}
	}
}
func prettyPrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

/*
	1. There is N/2 squares or cycles in a matrix of side N. Process a square one at a time. Run a loop to traverse
		the matrix a cycle at a time. i.e loop from 0 to N/2-1, loop counter is i
	2. Consider elements in group of 4 in current square, rotate the 4 elements at a time. So the number of such groups in a cycle
		is N-2*i

*/
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	clockWise := matrix
	antiClockWise := matrix
	rotateMatrix(clockWise, 1)
	prettyPrintMatrix(clockWise)
	fmt.Println()
	rotateMatrix(antiClockWise, 2)
	prettyPrintMatrix(antiClockWise)
}
