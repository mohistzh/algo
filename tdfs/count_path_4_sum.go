package main

import (
	"fmt"
)

// MyTreeNode5 tree node
type MyTreeNode5 struct {
	value       int
	left, right *MyTreeNode5
}

/*
Given a binary tree and a number ‘S’, find all paths in the tree such that the sum of all the node values of each path equals ‘S’.
Please note that the paths can start or end at any node but all paths must follow direction from parent to child (top to bottom).
*/
func countPathsForSum(root *MyTreeNode5, sum int) int {
	var currentPath []int // path memo, we are going to use it to calculate sum
	return countPathsForSumRecursive(root, sum, currentPath)
}
func countPathsForSumRecursive(currentNode *MyTreeNode5, sum int, currentPath []int) int {
	if currentNode == nil {
		return 0
	}
	count, sumPath := 0, 0
	currentPath = append(currentPath, currentNode.value)
	for i := len(currentPath) - 1; i >= 0; i-- {
		sumPath += currentPath[i]
		if sumPath == sum {
			count++
			break
		}
	}
	count += countPathsForSumRecursive(currentNode.left, sum, currentPath)
	count += countPathsForSumRecursive(currentNode.right, sum, currentPath)
	currentPath = currentPath[:len(currentPath)-1]
	return count

}

func main() {
	root := &MyTreeNode5{value: 1}
	root.left = &MyTreeNode5{value: 7}
	root.right = &MyTreeNode5{value: 9}
	root.left.left = &MyTreeNode5{value: 6}
	root.left.right = &MyTreeNode5{value: 5}
	root.right.left = &MyTreeNode5{value: 2}
	root.right.right = &MyTreeNode5{value: 3}

	fmt.Println(countPathsForSum(root, 12))

	// root = &MyTreeNode5{value: 12}
	// root.left = &MyTreeNode5{value: 7}
	// root.right = &MyTreeNode5{value: 1}
	// root.left.left = &MyTreeNode5{value: 4}
	// root.right.left = &MyTreeNode5{value: 10}
	// root.right.right = &MyTreeNode5{value: 5}
	// fmt.Println(countPathsForSum(root, 11))

}
