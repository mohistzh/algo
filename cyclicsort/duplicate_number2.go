package main

import "fmt"

/*
We are given an unsorted array containing ‘n+1’ numbers taken from the range 1 to ‘n’.
The array has only one duplicate but it can be repeated multiple times.
Find that duplicate number without using any extra space.

Can we solve the above problem in O(1) space and without modifying the input array?

We are using fast & slow pointers to find the duplicate number.
*/
func findDuplicateNumber2(input []int) int {
	slow, fast := input[0], input[input[0]]
	for slow != fast {
		slow = input[slow]
		fast = input[input[fast]]
	}
	current := input[input[slow]]
	cycleLength := 1
	for current != input[slow] {
		current = input[current]
		cycleLength++
	}
	return findStart(input, cycleLength)
}

func findStart(input []int, cycleLength int) int {
	pointer1, pointer2 := input[0], input[0]
	for cycleLength > 0 {
		pointer2 = input[pointer2]
		cycleLength--
	}
	for pointer1 != pointer2 {
		pointer1 = input[pointer1]
		pointer2 = input[pointer2]
	}
	return pointer1
}

func main() {
	fmt.Println(findDuplicateNumber2([]int{1, 4, 4, 3, 2}))
	fmt.Println(findDuplicateNumber2([]int{2, 1, 3, 3, 5, 4}))
	fmt.Println(findDuplicateNumber2([]int{2, 4, 1, 4, 4}))
}
