package main

import (
	"fmt"
)

// MyTreeNode1 tree node
type MyTreeNode1 struct {
	value       int
	left, right *MyTreeNode1
}

/*
Given a binary tree and a number ‘S’, find if the tree has a path from root-to-leaf such that the sum of all the node values of that path equals ‘S’.

*/
func findBinaryTreePathWithSum(root *MyTreeNode1, sum int) bool {
	if root == nil {
		return false
	}
	if root.value == sum && root.left == nil && root.right == nil {
		return true
	}
	return findBinaryTreePathWithSum(root.left, sum-root.value) || findBinaryTreePathWithSum(root.right, sum-root.value)
}

func main() {
	root := &MyTreeNode1{value: 12}
	root.left = &MyTreeNode1{value: 7}
	root.right = &MyTreeNode1{value: 1}
	root.left.left = &MyTreeNode1{value: 9}
	root.right.left = &MyTreeNode1{value: 10}
	root.right.right = &MyTreeNode1{value: 5}

	fmt.Println(findBinaryTreePathWithSum(root, 23))
	fmt.Println(findBinaryTreePathWithSum(root, 16))
}
