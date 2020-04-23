package main

import (
	"fmt"
)

/*
Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits, leads us to number ‘1’. All other (not-happy) numbers will never reach ‘1’.
Instead, they will be stuck in a cycle of numbers which does not include ‘1’.
*/
// findHappyNumber
func findHappyNumber(num int) bool {
	memo := make(map[int]int)
	return happyNumber(num, &memo)

}
func happyNumber(num int, memo *map[int]int) bool {
	if (*memo)[num] > 1 {
		return false
	}
	if num == 1 {
		return true
	}
	sum := calculateSquare(num)
	(*memo)[sum]++
	return happyNumber(sum, memo)
}

func calculateSquare(num int) int {
	currentSum := 0
	curNum := num
	for curNum > 0 {
		carry := curNum % 10
		currentSum += carry * carry
		curNum /= 10
	}
	return currentSum
}

func main() {
	fmt.Println(findHappyNumber(23))
	fmt.Println(findHappyNumber(12))
}
