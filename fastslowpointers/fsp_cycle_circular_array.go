package main

import (
	"fmt"
)

/*
	We are given an array containing positive and negative numbers. Suppose the array contains a number 'M' at a particular index.
	Now, if 'M' is positive we will move forward 'M' indices and if 'M' is negative move backwards 'M' indices.
	You should assume that the array is circular which means two things:
	1. If, while moving forward, we reach the end of the array, we will jump to the first element to continue the movement.
	2. If, while moving backward, we reach the beginning of the array, we will jump to the last element to continue the movement.
*/

func isCircularArrayLoop(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		isForward := arr[i] >= 0 // check is move forward or backward
		slow, fast := i, i
		for {
			slow = findNextIndex(arr, isForward, slow) // move one step for slow pointer
			fast = findNextIndex(arr, isForward, fast) // move one step for fast pointer
			if fast != -1 {
				fast = findNextIndex(arr, isForward, fast) // move another step for fast pointer
			}
			if slow == -1 || fast == -1 || slow == fast { // if slow or fast pointer is becomes '-1' this means we can't find cycle for this number
				break
			}
		}
		if slow != -1 && slow == fast {
			return true
		}
	}
	return false
}

func findNextIndex(arr []int, isForward bool, currentIndex int) int {
	nextIndex := -1
	direction := arr[currentIndex] >= 0
	if isForward != direction {
		return nextIndex // change in direction
	}
	nextIndex = (currentIndex + arr[currentIndex]) % len(arr)
	if nextIndex == currentIndex { // one element cycle found
		nextIndex = -1
	}
	return nextIndex
}

func main() {
	fmt.Println(isCircularArrayLoop([]int{1, 2, -1, 2, 2}))
	fmt.Println(isCircularArrayLoop([]int{2, 2, -1, 2}))
	fmt.Println(isCircularArrayLoop([]int{2, 1, -1, -2}))
}
