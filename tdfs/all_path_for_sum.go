package main

import (
	"fmt"
	"math"
)

// MyTreeNode2 tree node
type MyTreeNode2 struct {
	value       int
	left, right *MyTreeNode2
}

// AllPath use to store path data
type AllPath struct {
	store [][]int
}

/*
Given a binary tree and a number ‘S’, find all paths from root-to-leaf such that the sum of all the node values of each path equals ‘S’.

*/
func findAllPathForSum(root *MyTreeNode2, sum int) [][]int {
	var currentPath []int
	var allPath AllPath
	findAllPathRecursive(root, sum, currentPath, &allPath)
	return allPath.store
}

/*
Find path by recursive solution
*/
func findAllPathRecursive(currentNode *MyTreeNode2, sum int, currentPath []int, allPath *AllPath) {
	if currentNode == nil {
		return
	}
	currentPath = append(currentPath, currentNode.value)
	if currentNode.value == sum && currentNode.left == nil && currentNode.right == nil {
		allPath.store = append(allPath.store, currentPath)
		// remove currentNode in currentPath
		currentPath = currentPath[:len(currentPath)-1]
	} else {
		findAllPathRecursive(currentNode.left, sum-currentNode.value, currentPath, allPath)
		findAllPathRecursive(currentNode.right, sum-currentNode.value, currentPath, allPath)
	}

}

/*
Given a binary tree, find the root-to-leaf path with the maximum sum.
*/

func findMaximumSumPath(root *MyTreeNode2) (int, [][]int) {
	var allPath AllPath
	var currentPath []int
	// finding maximum sum first
	maximumSum := findMaximumSumRecursive(root)
	// finding path by sum
	findMaximumSumPathRecursive(root, maximumSum, currentPath, &allPath)
	return maximumSum, allPath.store
}
func findMaximumSumPathRecursive(currentNode *MyTreeNode2, sum int, currentPath []int, allPath *AllPath) {
	if currentNode == nil {
		return
	}
	currentPath = append(currentPath, currentNode.value)
	if currentNode.value == sum && currentNode.left == nil && currentNode.right == nil {
		allPath.store = append(allPath.store, currentPath)
		currentPath = currentPath[:len(currentPath)-1]
	} else {
		findMaximumSumPathRecursive(currentNode.left, sum-currentNode.value, currentPath, allPath)
		findMaximumSumPathRecursive(currentNode.right, sum-currentNode.value, currentPath, allPath)
	}

}
func findMaximumSumRecursive(currentNode *MyTreeNode2) int {
	if currentNode == nil {
		return 0
	}
	maxSumLeft := findMaximumSumRecursive(currentNode.left)
	maxSumRight := findMaximumSumRecursive(currentNode.right)
	return int(math.Max(float64(maxSumLeft), float64(maxSumRight))) + currentNode.value
}

func main() {
	root := &MyTreeNode2{value: 1}
	root.left = &MyTreeNode2{value: 7}
	root.right = &MyTreeNode2{value: 9}
	root.left.left = &MyTreeNode2{value: 4}
	root.left.right = &MyTreeNode2{value: 5}
	root.right.left = &MyTreeNode2{value: 2}
	root.right.right = &MyTreeNode2{value: 7}
	fmt.Println(findAllPathForSum(root, 12))
	root = &MyTreeNode2{value: 12}
	root.left = &MyTreeNode2{value: 7}
	root.right = &MyTreeNode2{value: 1}
	root.left.left = &MyTreeNode2{value: 4}
	root.right.left = &MyTreeNode2{value: 10}
	root.right.right = &MyTreeNode2{value: 5}
	fmt.Println(findAllPathForSum(root, 23))
	fmt.Println(findMaximumSumPath(root))

}
