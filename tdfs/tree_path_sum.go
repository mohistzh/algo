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
func findBinaryTreePathWithSum(root *MyTreeNode1, sum int) []int {
	var result []int
	return result
}

func main() {
	root := &MyTreeNode1{value: 1}
	root.left = &MyTreeNode1{value: 2}
	root.right = &MyTreeNode1{value: 3}
	root.left.left = &MyTreeNode1{value: 4}
	root.left.right = &MyTreeNode1{value: 5}
	root.right.left = &MyTreeNode1{value: 6}
	root.right.right = &MyTreeNode1{value: 7}

	fmt.Println(findBinaryTreePathWithSum(root, 10))
}
