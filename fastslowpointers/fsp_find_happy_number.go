package main

import (
	"fmt"
)

/*
Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits,
leads us to number ‘1’. All other (not-happy) numbers will never reach ‘1’. Instead, they will be stuck in a cycle of numbers which does not include ‘1’.
*/
func findHappyNumber(num int) bool {
	slow, fast := num, num
	for {
		slow = squareNumber(slow)               // move one step
		fast = squareNumber(squareNumber(fast)) // move two steps
		if slow == fast {                       // cycle found
			break
		}
	}
	return slow == 1 // check if the cycle is stuck on the number 1
}

func squareNumber(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(findHappyNumber(23))
	fmt.Println(findHappyNumber(12))
}
