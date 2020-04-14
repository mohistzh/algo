package main

import (
	"fmt"
)

/*
	Given a sorted array, create a new array containing squares of all the number of the input array in the sorted order.
*/
func makeSquare(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)
	leftPoint, rightPoint, highestIndex := 0, n-1, n-1
	for leftPoint <= rightPoint {
		leftSquare := arr[leftPoint] * arr[leftPoint]
		rightSquare := arr[rightPoint] * arr[rightPoint]

		if leftSquare > rightSquare {
			squares[highestIndex] = leftSquare
			leftPoint++
		} else {
			squares[highestIndex] = rightSquare
			rightPoint--
		}
		highestIndex--
	}
	return squares
}

func main() {
	fmt.Println(makeSquare([]int{-2, -1, 0, 2, 3}))
	fmt.Println(makeSquare([]int{-3, -1, 0, 1, 2}))
}
