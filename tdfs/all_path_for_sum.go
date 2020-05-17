package main

import (
	"fmt"
)

// MyTreeNode2 tree node
type MyTreeNode2 struct {
	value       int
	left, right *MyTreeNode2
}

/*
Given a binary tree and a number ‘S’, find all paths from root-to-leaf such that the sum of all the node values of each path equals ‘S’.

*/
func findAllPathForSum(root *MyTreeNode2, sum int) {
	var currentPath []int
	findAllPathRecursive(root, sum, currentPath)
}

/*
Find path by recursive solution
*/
func findAllPathRecursive(currentNode *MyTreeNode2, sum int, currentPath []int) {
	if currentNode == nil {
		return
	}
	currentPath = append(currentPath, currentNode.value)
	if currentNode.value == sum && currentNode.left == nil && currentNode.right == nil {
		fmt.Println(currentPath)
		// remove currentNode in currentPath
		currentPath = currentPath[:len(currentPath)-1]
	} else {
		findAllPathRecursive(currentNode.left, sum-currentNode.value, currentPath)
		findAllPathRecursive(currentNode.right, sum-currentNode.value, currentPath)
	}

}

func main() {
	root := &MyTreeNode2{value: 1}
	root.left = &MyTreeNode2{value: 7}
	root.right = &MyTreeNode2{value: 9}
	root.left.left = &MyTreeNode2{value: 4}
	root.left.right = &MyTreeNode2{value: 5}
	root.right.left = &MyTreeNode2{value: 2}
	root.right.right = &MyTreeNode2{value: 7}
	findAllPathForSum(root, 12)

}
