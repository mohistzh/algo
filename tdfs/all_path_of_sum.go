package main

import (
	"fmt"
	"strconv"
)

// MyTreeNode3 tree node
type MyTreeNode3 struct {
	value       int
	left, right *MyTreeNode3
}

/*
Given a binary tree where each node can only have a digit (0-9) value,
each root-to-leaf path will represent a number.
Find the total sum of all the numbers represented by all paths.
*/
func sumOfPathNumbers(root *MyTreeNode3) int {
	var currentPath []int
	sum1 := sumOfPathNumbersRecursiveBF(root, currentPath)
	sum2 := sumOfPathNumbersRecursiveMath(root, 0)
	fmt.Println(sum1, sum2)
	return sum1
}

func sumOfPathNumbersRecursiveBF(currentNode *MyTreeNode3, currentPath []int) int {
	if currentNode == nil {
		return 0
	}

	currentPath = append(currentPath, currentNode.value)
	if currentNode.left == nil && currentNode.right == nil {
		// concatenating elements in the int slice
		var str string
		for i := 0; i < len(currentPath); i++ {
			str += strconv.Itoa(currentPath[i])
		}
		sumOfPath, err := strconv.Atoi(str)
		if err != nil {
			panic("convert an integer slice to int occurs exception")
		}
		currentPath = currentPath[:len(currentPath)-1]
		return int(sumOfPath)
	}
	return sumOfPathNumbersRecursiveBF(currentNode.left, currentPath) + sumOfPathNumbersRecursiveBF(currentNode.right, currentPath)

}

func sumOfPathNumbersRecursiveMath(currentNode *MyTreeNode3, sumPath int) int {
	if currentNode == nil {
		return 0
	}
	sumPath = 10*sumPath + currentNode.value
	if currentNode.left == nil && currentNode.right == nil {
		return sumPath
	}
	return sumOfPathNumbersRecursiveMath(currentNode.left, sumPath) + sumOfPathNumbersRecursiveMath(currentNode.right, sumPath)
}

func main() {
	root := &MyTreeNode3{value: 1}
	root.left = &MyTreeNode3{value: 7}
	root.right = &MyTreeNode3{value: 9}
	root.right.left = &MyTreeNode3{value: 2}
	root.right.right = &MyTreeNode3{value: 9}
	fmt.Println(sumOfPathNumbers(root))

	root2 := &MyTreeNode3{value: 1}
	root2.left = &MyTreeNode3{value: 0}
	root2.right = &MyTreeNode3{value: 1}
	root2.left.left = &MyTreeNode3{value: 1}
	root2.right.left = &MyTreeNode3{value: 6}
	root2.right.right = &MyTreeNode3{value: 5}
	fmt.Println(sumOfPathNumbers(root2))
}
